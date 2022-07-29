package model

var (
	PushEventMsg          = "*%s* pushed to branch *%s* in [%s/%s](%s) \n*commit*: `%s`"
	MergeRequestEventsMsg = "*%s* created in branch *%s* in [%s/%s](%s) \n* a merge request*"
	PipelineEventsMsg     = "*%s* pipeline in branch *%s* in [%s/%s](%s) \n* %s*"
	DeployEventsMesg      = "*%s* deploy production in branch *%s* in [%s/%s](%s) \n* %s*"
)
