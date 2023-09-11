package action

import (
	"encoding/json"
	"io/ioutil"

	"github.com/google/go-github/v55/github"
)

type Event interface{}

func GetEvent() (Event, error) {
	var evt Event

	data, err := ioutil.ReadFile(Context.EventPath)
	if err != nil {
		return nil, err
	}

	switch Context.EventName {
	case "branch_protection_rule":
		return parseEvent(data, &github.BranchProtectionRuleEvent{})

	case "check_run":
		return parseEvent(data, &github.CheckRunEvent{})

	case "check_suite":
		return parseEvent(data, &github.CheckSuiteEvent{})

	case "create":
		return parseEvent(data, &github.CreateEvent{})

	case "delete":
		return parseEvent(data, &github.DeleteEvent{})

	case "deployment":
		return parseEvent(data, &github.DeploymentEvent{})

	case "deployment_status":
		return parseEvent(data, &github.DeploymentStatus{})

	case "discussion":
		return parseEvent(data, &github.Discussion{})

	case "fork":
		return parseEvent(data, &github.ForkEvent{})

	case "gollum":
		return parseEvent(data, &github.GollumEvent{})

	case "issue_comment":
		return parseEvent(data, &github.IssueCommentEvent{})

	case "issues":
		return parseEvent(data, &github.IssuesEvent{})

	case "label":
		return parseEvent(data, &github.LabelEvent{})

	case "milestone":
		return parseEvent(data, &github.MilestoneEvent{})

	case "page_build":
		return parseEvent(data, &github.PageBuildEvent{})

	case "project":
		return parseEvent(data, &github.ProjectEvent{})

	case "project_card":
		return parseEvent(data, &github.ProjectCardEvent{})

	case "project_column":
		return parseEvent(data, &github.ProjectColumnEvent{})

	case "public":
		return parseEvent(data, &github.PublicEvent{})

	case "pull_request":
		return parseEvent(data, &github.PullRequestEvent{})

	// https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#pull_request_comment-use-issue_comment
	case "pull_request_comment":
		return parseEvent(data, &github.IssueCommentEvent{})

	case "pull_request_review":
		return parseEvent(data, &github.PullRequestReviewEvent{})

	case "pull_request_review_comment":
		return parseEvent(data, &github.PullRequestReviewCommentEvent{})

	case "pull_request_target":
		return parseEvent(data, &github.PullRequestTargetEvent{})

	case "push":
		return parseEvent(data, &github.PushEvent{})

	case "release":
		return parseEvent(data, &github.ReleaseEvent{})

	case "repository_dispatch":
		return parseEvent(data, &github.RepositoryDispatchEvent{})

	case "status":
		return parseEvent(data, &github.StatusEvent{})

	case "watch":
		return parseEvent(data, &github.WatchEvent{})

	case "workflow_dispatch":
		return parseEvent(data, &github.WorkflowDispatchEvent{})

	case "workflow_run":
		return parseEvent(data, &github.WorkflowRunEvent{})

	// events not implemented
	case "discussion_comment":
	case "registry_package":
	case "schedule":
	case "workflow_call":
	default:
		return parseEvent(data, &map[string]interface{}{})
	}

	return evt, nil
}

func parseEvent(data []byte, evt Event) (Event, error) {
	if err := json.Unmarshal(data, evt); err != nil {
		return nil, err
	}

	return evt, nil
}
