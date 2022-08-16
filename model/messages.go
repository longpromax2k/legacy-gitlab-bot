package model

var (
	PushEventMsg          = "*%s* pushed to branch *%s* in [%s/%s](%s) \n*commit*: `%s`."
	IssueEventMsg         = "Issue [#%d %s](%s) opened by [%s](gitlab.com/%s). \n> *%s*\n> %s."
	MergeRequestEventsMsg = "*%s* created in branch *%s* in [%s/%s](%s) \n* a merge request*."
	PipelineEventsMsg     = "*%s* pipeline in branch *%s* in [%s/%s](%s) \n* %s*."
	DeployEventsMsg       = "*%s* deploy production in branch *%s* in [%s/%s](%s) \n* %s*."
	ReleaseEventsMsg      = "*%s* release project in branch *%s* in [%s/%s](%s) \n*message*: `%s`."
	WikipageEventsMsg     = "*%s* created a Wikipage with title *%s*."
	TagEventsMsg          = "*%s* created a tag in branch *%s*."
	CmtCommitMsg          = "*%s* commented on the commit [%s](%s).\n> %s"
	CmtIssueMsg           = "*%s* commented on the issue [#%d %s](%s).\n> %s"
	CmtMergeMsg           = "*%s* commented on the merge request [%s](%s).\n> %s"
	CmtSnippetMsg         = "*%s* commented on the code snippet [#%s](%s).\n> %s"
	JobsEvent             = "%s in %s was %s."
	FeatFlagMsg           = "%s is %s."
	StatusDraftMrMsg      = "Merge Request is marked as Draft  %s"
	StatusReadyMrMsg      = "Merge Request is marked as Ready  %s "
)

var (
	ChatExistMsg  = "Your chat is already added! Use this Webhook URL to setup the notification:\n\n`https://%s/%s/%s.%s`"
	ChatInsertMsg = "Hi there! To setup notifications *for this chat* with your GitLab repository, open Settings/Webhooks and add this URL:\n\n`https://%s/%s/%s.%s`"
	ChatNotCmdMsg = "Invalid Command. To interact with me:\n\n`/start` to get a webhook link.\n`/drop` to drop webhook link."
	ChatDropMsg   = "Your notification is dropped, you'll no longer receive any message until you start a new one."
)
