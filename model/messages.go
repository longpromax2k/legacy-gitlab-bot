package model

var (
	PushEventMsg  = "*%s* pushed to branch *%s* in [%s/%s](%s) \n*commit*: `%s`"
	IssueEventMsg = "Issue [#%d %s](%s) opened by [%s](gitlab.com/%s). \n> *%s*\n> %s"
)
