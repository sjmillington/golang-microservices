package config

import "os"

//can set these environment variables in docker.
//ENV GO_ENVIRONMENT=production
//ENV API_GITHUB_ACCESS_TOKEN=sdsdsajidsaiodjsaiodj
//ENV LOG_LEVEL=....

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel             = "info"
	goEnvironment        = "GO_ENVIRONMENT"
	production           = "production"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}

func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
