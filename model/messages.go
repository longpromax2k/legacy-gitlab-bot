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
)
