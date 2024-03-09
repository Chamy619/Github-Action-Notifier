package main

import "encoding/json"

const (
	ADDED = "added"
	MODIFIED = "modified"
	REMOVED = "removed"
)

type GithubActionService struct {
	*Config
}

func NewGithubActionService(config *Config) *GithubActionService {
	return &GithubActionService{
		config,
	}
}

func (g *GithubActionService) Pushed(request GithubActionHttpRequestBody) error {
	var githubAction GithubAction
	err := json.Unmarshal([]byte(request.Payload), &githubAction)
	if err != nil {
		return err
	}

	for _, notifier := range g.Config.Notifies {
		for _, file := range notifier.Files {
			if file.Added && githubAction.IsContained(file.Name, ADDED) {
				Send(file.Messenger.Type, file.Messenger.Url, file.Message)
			}
			if file.Modified && githubAction.IsContained(file.Name, MODIFIED) {
				Send(file.Messenger.Type, file.Messenger.Url, file.Message)
			}
			if file.Removed && githubAction.IsContained(file.Name, REMOVED) {
				Send(file.Messenger.Type, file.Messenger.Url, file.Message)
			}
		}
	}
	return nil
}