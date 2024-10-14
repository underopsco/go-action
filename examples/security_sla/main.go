package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v55/github"
	"github.com/underopsco/go-action/pkg/action"
)

func main() {
	if err := action.Execute(&Action{}); err != nil {
		action.SetFailed(err, map[string]string{})
	}
}

type Action struct {
	Critical int `action:"critical-threshold"`
	High     int `action:"high-threshold"`
	Medium   int `action:"medium-threshold"`
	Low      int `action:"low-threshold"`
}

func (a *Action) Run() error {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	ctx := context.Background()

	criticalThreshold := time.Duration(a.Critical*24) * time.Hour
	highThreshold := time.Duration(a.High*24) * time.Hour
	mediumThreshold := time.Duration(a.Medium*24) * time.Hour
	lowThreshold := time.Duration(a.Low*24) * time.Hour

	event, err := action.GetEvent()
	if err != nil {
		return err
	}

	prEvent, ok := event.(*github.PullRequestEvent)
	if !ok {
		return fmt.Errorf("unexpected event type: %T", event)
	}

	slog.Debug("Configuration",
		slog.String("critical", criticalThreshold.String()),
		slog.String("high", highThreshold.String()),
		slog.String("medium", mediumThreshold.String()),
		slog.String("low", lowThreshold.String()))

	alerts, err := fetchRepoAlerts(
		ctx,
		action.Context.RepositoryOwner,
		action.Context.RepositoryName,
	)
	if err != nil {
		return err
	}

	if len(alerts) == 0 {
		_, _, err := action.REST.PullRequests.CreateReview(
			ctx,
			action.Context.RepositoryOwner,
			action.Context.RepositoryName,
			prEvent.GetNumber(),
			&github.PullRequestReviewRequest{
				Event: github.String("APPROVE"),
			},
		)
		return err
	}

	breached := filterBreachedAlerts(
		alerts,
		criticalThreshold,
		highThreshold,
		mediumThreshold,
		lowThreshold,
	)

	slog.Info("Alerts found",
		slog.Int("total", len(alerts)),
		slog.Int("breached", len(breached)))

	if len(breached) == 0 {
		_, _, err := action.REST.PullRequests.CreateReview(
			ctx,
			action.Context.RepositoryOwner,
			action.Context.RepositoryName,
			prEvent.GetNumber(),
			&github.PullRequestReviewRequest{
				Event: github.String("APPROVE"),
			},
		)
		return err
	}

	_, _, err = action.REST.PullRequests.CreateReview(
		ctx,
		action.Context.RepositoryOwner,
		action.Context.RepositoryName,
		prEvent.GetNumber(),
		&github.PullRequestReviewRequest{
			Event: github.String("REQUEST_CHANGES"),
			Body:  github.String("Security SLA alerts"),
		},
	)
	return err
}

type Alert struct {
	Kind      string
	Severity  string
	Link      string
	CreatedAt time.Time
}

func fetchRepoAlerts(ctx context.Context, owner, name string) ([]*Alert, error) {
	var alerts []*Alert

	codeScanningAlerts, _, err := action.REST.CodeScanning.ListAlertsForRepo(
		ctx, owner, name,
		&github.AlertListOptions{
			State: "open",
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		},
	)
	if err != nil && !isDisabledError(err) {
		return nil, err
	}

	for _, a := range codeScanningAlerts {
		alerts = append(alerts, &Alert{
			Kind:      "CodeScanning",
			Severity:  *a.RuleSeverity,
			Link:      *a.HTMLURL,
			CreatedAt: a.GetCreatedAt().Time,
		})
	}

	dependabotAlerts, _, err := action.REST.Dependabot.ListRepoAlerts(
		ctx, owner, name,
		&github.ListAlertsOptions{
			State: github.String("open"),
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		},
	)
	if err != nil && !isDisabledError(err) {
		return nil, err
	}

	for _, a := range dependabotAlerts {
		alerts = append(alerts, &Alert{
			Kind:      "Dependabot",
			Severity:  a.SecurityAdvisory.GetSeverity(),
			Link:      a.GetHTMLURL(),
			CreatedAt: a.GetCreatedAt().Time,
		})
	}

	secretScanningAlerts, _, err := action.REST.SecretScanning.ListAlertsForRepo(
		ctx, owner, name,
		&github.SecretScanningAlertListOptions{
			State: "open",
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		},
	)
	if err != nil && !isDisabledError(err) {
		return nil, err
	}

	for _, a := range secretScanningAlerts {
		alerts = append(alerts, &Alert{
			Kind:      "SecretScanning",
			Severity:  "critical",
			Link:      a.GetHTMLURL(),
			CreatedAt: a.GetCreatedAt().Time,
		})
	}

	return alerts, nil
}

func filterBreachedAlerts(alerts []*Alert, criticalThreshold, highThreshold, mediumThreshold, lowThreshold time.Duration) []*Alert {
	var breached []*Alert

	for _, a := range alerts {
		s := time.Since(a.CreatedAt)

		if s > criticalThreshold || s > highThreshold || s > mediumThreshold || s > lowThreshold {
			breached = append(breached, a)
		}
	}

	return breached
}

func isDisabledError(err error) bool {
	resp, ok := err.(*github.ErrorResponse)
	if ok && strings.Contains(resp.Message, "disabled") {
		slog.Warn("Security feature disabled", slog.String("error", resp.Message))
		return true
	}
	return false
}
