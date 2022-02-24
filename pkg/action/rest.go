package action

import (
	"context"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

var REST = createRESTClient()

func createRESTClient() *github.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: Context.Token},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	return github.NewClient(httpClient)
}
