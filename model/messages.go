package model

var (
	PushEventMsg          = "*%s* pushed to branch *%s* in [%s/%s](%s) \n*commit*: `%s`"
	MergeRequestEventsMsg = "*%s* created in branch *%s* in [%s/%s](%s) \n* a merge request*"
	PipelineEventsMsg     = "*%s* pipeline in branch *%s* in [%s/%s](%s) \n* %s*"
	DeployEventsMsg       = "*%s* deploy production in branch *%s* in [%s/%s](%s) \n* %s*"
	ReleaseEventsMsg      = "*%s* release project in branch *%s* in [%s/%s](%s) \n*message*: `%s`"
	WikipageEventsMsg     = "*%s* created a Wikipage with title *%s*"
	TagEventsMsg          = "*%s* created a tag in branch *%s*"
)
