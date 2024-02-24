package main

import (
	"fmt"
)

type GithubAction struct {
	Ref string `json:"ref"`
	Before string `json:"before"`
	After string `json:"after"`
	Repository GithubActionRepository `json:"repository"`
	Pusher GithubActionPusher `json:"pusher"`
	Sender GithubActionSender `json:"sender"`
	Created bool `json:"created"`
	Deleted bool `json:"deleted"`
	Forced bool `json:"forced"`
	BaseRef string `json:"base_ref"`
	Compare string `json:"compare"`
	Commits []GithubActionCommit `json:"commits"`
	HeadCommit GithubActionHeadCommit `json:"head_commit"`
}

type GithubActionRepository struct {
	Id uint64 `json:"id"`
	NodeId string `json:"node_id"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Private bool `json:"private"`
	Owner GithubActionRepositoryOwner `json:"owner"`
	HtmlUrl string `json:"html_url"`
	Description string `json:"description"`
	Fork bool `json:"fork"`
	Url string `json:"url"`
	ForksUrl string `json:"forks_url"`
	KeysUrl string `json:"keys_url"`
	CollaboratorsUrl string `json:"collaborators_url"`
	TeamsUrl string `json:"teams_url"`
	HooksUrl string `json:"hooks_url"`
	IssuEventsUrl string `json:"issue_events_url"`
	EventsUrl string `json:"events_url"`
	AssigneesUrl string `json:"assignees_url"`
	BranchesUrl string `json:"branches_url"`
	TagsUrl string `json:"tags_url"`
	BlobsUrl string `json:"blobs_url"`
	GitTagsUrl string `json:"git_tags_url"`
	GitRefsUrl string `json:"git_refs_url"`
	TreesUrl string `json:"trees_url"`
	StatusesUrl string `json:"statuses_url"`
	LanguagesUrl string `json:"languages_url"`
	StargazersUrl string `json:"stargazers_url"`
	ContributorsUrl string `json:"contributors_url"`
	SubscribersUrl string `json:"subscribers_url"`
	SubscriptionUrl string `json:"subscription_url"`
	CommitsUrl string `json:"commits_url"`
	GitCommitsUrl string `json:"git_commits_url"`
	CommentsUrl string `json:"comments_url"`
	IssueCommentUrl string `json:"issue_comment_url"`
	ContentsUrl string `json:"contents_url"`
	CompareUrl string `json:"compare_url"`
	MergesUrl string `json:"merges_url"`
	ArchiveUrl string `json:"archive_url"`
	DownloadsUrl string `json:"downloads_url"`
	IssuesUrl string `json:"issues_url"`
	PullsUrl string `json:"pulls_url"`
	MilestonesUrl string `json:"milestones_url"`
	NotificationsUrl string `json:"notifications_url"`
	LabelsUrl string `json:"labels_url"`
	ReleasesUrl string `json:"releases_url"`
	DeploymentsUrl string `json:"deployments_url"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	PushedAt uint64 `json:"pushed_at"`
	GitUrl string `json:"git_url"`
	SshUrl string `json:"ssh_url"`
	CloneUrl string `json:"clone_url"`
	SvnUrl string `json:"svn_url"`
	Homepage string `json:"homepage"`
	Size uint64 `json:"size"`
	StargazersCount uint64 `json:"stargazers_count"`
	WatchersCount uint64 `json:"watchers_count"`
	Language string `json:"language"`
	HasIssues bool`json:"has_issues"`
	HasProjects bool`json:"has_projects"`
	HasDownloads bool `json:"has_downloads"`
	HasWiki bool `json:"has_wiki"`
	HasPages bool `json:"has_pages"`
	HasDiscussions bool `json:"has_discussions"`
	ForksCount uint64 `json:"forks_count"`
	MirrorUrl string `json:"mirror_url"`
	Archived bool `json:"archived"`
	Disabled bool `json:"disabled"`
	OpenIssuesCount uint64 `json:"open_issues_count"`
	License string `json:"license"`
	AllowForking bool `json:"allow_forking"`
	IsTemplate bool `json:"is_template"`
	WebCommitSignoffRequired bool `json:"web_commit_signoff_required"`
	Visibility string `json:"visibility"`
	Forks uint64 `json:"forks"`
	OpenIssues uint64 `json:"open_issues"`
	Watchers uint64 `json:"watchers"`
	DefaultBranch string `json:"default_branch"`
	Stargazers uint64 `json:"stargazers"`
	MasterBranch string `json:"master_branch"`
	Organization string `json:"organization"`
	Topics []interface{} `json:"topics"`
}

type GithubActionRepositoryOwner struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Login string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	GravatarId string `json:"gravatar_id"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	FollowersUrl string `json:"followers_url"`
	GistsUrl string `json:"gists_url"`
	StarredUrl string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	OrganizationsUrl string `json:"organizations_url"`
	ReposUrl string `json:"repos_url"`
	EventsUrl string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
}

type GithubActionPusher struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type GithubActionSender struct {
	Login string `json:"login"`
	Id uint64 `json:"id"`
	NodeId string `json:"node_id"`
	AvatarUrl string `json:"avatar_url"`
	GravatarId string `json:"gravatar_id"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	FollowersUrl string `json:"followers_url"`
	FollowingUrl string `json:"following_url"`
	GistsUrl string `json:"gists_url"`
	StarredUrl string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	OrganizationUrl string `json:"organizations_url"`
	ReposeUrl string `json:"repos_url"`
	EventsUrl string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
}

type GithubActionCommit struct {
	Id string `json:"id"`
	TreeId string `json:"tree_id"`
	Distinct bool `json:"distinct"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	Url string `json:"url"`
	Author GithubActionCommitAuthor `json:"author"`
	Committer GithubActionCommitCommitter `json:"committer"`
	Added []string `json:"added"`
	Removed []string `json:"removed"`
	Modified []string `json:"modified"`
}

type GithubActionCommitAuthor struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
}

type GithubActionCommitCommitter struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
}

type GithubActionHeadCommit struct {
	Id string `json:"id"`
	TreeId string `json:"tree_id"`
	Distinct bool `json:"distinct"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	Url string `json:"url"`
	Author GithubActionHeadCommitAuthor `json:"author"`
	Committer GithubActionHeadCommitCommiter `json:"committer"`
	Added []string `json:"added"`
	Removed []string `json:"removed"`
	Modified []string `json:"modified"`
}

type GithubActionHeadCommitAuthor struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
}

type GithubActionHeadCommitCommiter struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
}

func main() {
	fmt.Println("Welcome github action notifier")
}