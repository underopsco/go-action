package action

import (
	"os"
	"strconv"
)

// https://docs.github.com/en/actions/learn-github-actions/environment-variables#default-environment-variables
type GitHubContext struct {
	// A token to authenticate on behalf of the GitHub App installed on your repository
	Token string

	// Indicates if it's running in a CI environment
	CI bool

	// The name of the action currently running, or the id of a step
	Action string

	// The path where an action is located. This property is only supported in composite actions.
	// You can use this path to access files located in the same repository as the action
	ActionPath string

	// For a step executing an action, this is the owner and repository name of the action
	ActionRepository string

	// Always set to true when GitHub Actions is running the workflow.
	// You can use this variable to differentiate when tests are being run locally or by GitHub Actions
	Actions bool

	// The name of the person or app that initiated the workflow
	Actor string

	// The API URL
	APIURL string

	// The name of the base ref or target branch of the pull request in a workflow run. This is only
	// set when the event that triggers a workflow run is either pull_request or pull_request_target
	BaseRef string

	// The path on the runner to the file that sets environment variables from workflow commands.
	// This file is unique to the current step and changes for each step in a job
	// https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#setting-an-environment-variable
	Env string

	// The name of the event that triggered the workflow
	EventName string

	// The path to the file on the runner that contains the full event webhook payload
	EventPath string

	// The GraphQL API URL
	GraphQLURL string

	// The head ref or source branch of the pull request in a workflow run.
	// This property is only set when the event that triggers a workflow run is either pull_request
	// or pull_request_target
	HeadRef string

	// The job_id of the current job
	Job string

	// The path on the runner to the file that sets system PATH variables from workflow commands.
	// This file is unique to the current step and changes for each step in a job
	Path string

	// The branch or tag ref that triggered the workflow run. For branches this is the format refs/heads/<branch_name>,
	// for tags it is refs/tags/<tag_name>, and for pull requests it is refs/pull/<pr_number>/merge.
	// This variable is only set if a branch or tag is available for the event type
	Ref string

	// The branch or tag name that triggered the workflow run
	RefName string

	// Is true if branch protections are configured for the ref that triggered the workflow run
	RefProtected string

	// The type of ref that triggered the workflow run. Valid values are branch or tag.
	RefType string

	// The owner and repository name
	Repository string

	// The repository owner's name
	RepositoryOwner string

	// The number of days that workflow run logs and artifacts are kept
	RetentionDays int

	// A unique number for each attempt of a particular workflow run in a repository.
	// This number begins at 1 for the workflow run's first attempt, and increments with each re-run
	RunAttempt int

	// A unique number for each workflow run within a repository. This number does not change if you re-run
	// the workflow run
	RunID int

	// A unique number for each run of a particular workflow in a repository. This number begins at 1
	// for the workflow's first run, and increments with each new run. This number does not change if you
	// re-run the workflow run
	RunNumber int

	// The URL of the GitHub server
	ServerURL string

	// The commit SHA that triggered the workflow
	SHA string

	// The name of the workflow. If the workflow file doesn't specify a name, the value of this variable is
	// the full path of the workflow file in the repository
	Workflow string

	// The default working directory on the runner for steps, and the default location of your repository
	// when using the checkout action
	Workspace string

	// The architecture of the runner executing the job. Possible values are X86, X64, ARM, or ARM64.
	RunnerArch string

	// The name of the runner executing the job
	RunnerName string

	// The operating system of the runner executing the job. Possible values are Linux, Windows, or macOS
	RunnerOS string

	// The path to a temporary directory on the runner. This directory is emptied at the beginning and end
	// of each job. Note that files will not be removed if the runner's user account does not have permission to delete them
	RunnerTemp string

	// The path to the directory containing preinstalled tools for GitHub-hosted runners
	RunnerToolCache string

	RunnerDebug bool
}

var Context = &GitHubContext{
	Token:            os.Getenv("GITHUB_TOKEN"),
	CI:               boolEnv("CI"),
	Action:           os.Getenv("GITHUB_ACTION"),
	ActionPath:       os.Getenv("GITHUB_ACTION_PATH"),
	ActionRepository: os.Getenv("GITHUB_ACTION_REPOSITORY"),
	Actions:          boolEnv("GITHUB_ACTIONS"),
	Actor:            os.Getenv("GITHUB_ACTOR"),
	APIURL:           os.Getenv("GITHUB_API_URL"),
	BaseRef:          os.Getenv("GITHUB_BASE_REF"),
	Env:              os.Getenv("GITHUB_ENV"),
	EventName:        os.Getenv("GITHUB_EVENT_NAME"),
	EventPath:        os.Getenv("GITHUB_EVENT_PATH"),
	GraphQLURL:       os.Getenv("GITHUB_GRAPHQL_URL"),
	HeadRef:          os.Getenv("GITHUB_HEAD_REF"),
	Job:              os.Getenv("GITHUB_JOB"),
	Path:             os.Getenv("GITHUB_PATH"),
	Ref:              os.Getenv("GITHUB_REF"),
	RefName:          os.Getenv("GITHUB_REF_NAME"),
	RefProtected:     os.Getenv("GITHUB_REF_PROTECTED"),
	RefType:          os.Getenv("GITHUB_REF_TYPE"),
	Repository:       os.Getenv("GITHUB_REPOSITORY"),
	RepositoryOwner:  os.Getenv("GITHUB_REPOSITORY_OWNER"),
	RetentionDays:    intEnv("GITHUB_RETENTION_DAYS"),
	RunAttempt:       intEnv("GITHUB_RUN_ATTEMPT"),
	RunID:            intEnv("GITHUB_RUN_ID"),
	RunNumber:        intEnv("GITHUB_RUN_NUMBER"),
	ServerURL:        os.Getenv("GITHUB_SERVER_URL"),
	SHA:              os.Getenv("GITHUB_SHA"),
	Workflow:         os.Getenv("GITHUB_WORKFLOW"),
	Workspace:        os.Getenv("GITHUB_WORKSPACE"),
	RunnerArch:       os.Getenv("RUNNER_ARCH"),
	RunnerName:       os.Getenv("RUNNER_NAME"),
	RunnerOS:         os.Getenv("RUNNER_OS"),
	RunnerTemp:       os.Getenv("RUNNER_TEMP"),
	RunnerToolCache:  os.Getenv("RUNNER_TOOL_CACHE"),
	RunnerDebug:      boolEnv("RUNNER_DEBUG"),
}

func boolEnv(key string) bool {
	switch os.Getenv(key) {
	case "true", "1":
		return true

	case "false", "0":
		return false
	}

	return false
}

func intEnv(key string) int {
	if envVal := os.Getenv(key); envVal != "" {
		val, err := strconv.Atoi(envVal)
		if err != nil {
			panic(err)
		}

		return val
	}

	return 0
}
