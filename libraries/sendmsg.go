package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	h "gitlabhook/helpers"
	mdl "gitlabhook/model"
	webhook "gitlabhook/model/webhook"
	cmt "gitlabhook/model/webhook/comment"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendTelegramMessage(pay mdl.ObjectKind, body []byte, cId string) {
	cid, _ := strconv.Atoi(cId)
	var chatId = int64(cid)
	var dt, url, text string
	var err error

	switch pay.ObjectKind {
	case "push":
		var p webhook.PushEventPayload
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.PushEventMsg, p.UserUsername, p.Ref, p.UserUsername, p.Project.Name, p.Project.Homepage, p.Commits[0].Message)
		url, text = p.Commits[0].URL, "Open Commit"
	case "issue":
		var p webhook.Issues
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.IssueEventMsg, p.ObjectAttributes.Iid, p.ObjectAttributes.Title, p.ObjectAttributes.URL, p.User.Name, p.User.Username, p.ObjectAttributes.Title, p.ObjectAttributes.Description)
		url, text = p.ObjectAttributes.URL, "Open Issue"
	case "note":
		var nType cmt.NoteableType
		err = json.Unmarshal(body, &nType)
		switch nType.ObjectAttributes.NoteableType {
		case "Commit":
			var p cmt.Commit
			err = json.Unmarshal(body, &p)
			dt = fmt.Sprintf(mdl.CmtCommitMsg, p.Commit.Author.Name, p.Commit.ID, p.Commit.URL, p.ObjectAttributes.Note)
			url, text = p.Commit.URL, "Open Commit"
		case "Issue":
			var p cmt.Issues
			err = json.Unmarshal(body, &p)
			dt = fmt.Sprintf(mdl.CmtIssueMsg, p.User.Name, p.Issue.Iid, p.Issue.Title, p.ObjectAttributes.URL, p.ObjectAttributes.Note)
			url, text = p.ObjectAttributes.URL, "Open Issue"
		case "MergeRequest":
			var p cmt.MergeRequest
			err = json.Unmarshal(body, &p)
			statusMR := strings.Split(p.MergeRequest.Title, ":")
			draftMR := statusMR[0]
			if draftMR != "Draft" {
				dt = fmt.Sprintf(mdl.CmtMergeMsg, p.User.Name, p.MergeRequest.Title, p.ObjectAttributes.URL, p.ObjectAttributes.Note)
				url, text = p.ObjectAttributes.URL, "Open Merge Request"
			}
		case "Snippet":
			var p cmt.CodeSnippet
			err = json.Unmarshal(body, &p)
			dt = fmt.Sprintf(mdl.CmtSnippetMsg, p.User.Name, p.ObjectAttributes.ID, p.ObjectAttributes.URL, p.ObjectAttributes.Note)
			url, text = p.ObjectAttributes.URL, "Open Code Snippet"
		default:
			log.Fatalf("Invalid NoteableType.\n")
		}
	case "merge_request":
		var p webhook.MergeRequestEventsLoad
		err = json.Unmarshal(body, &p)
		statusMR := strings.Split(p.ObjectAttributes.Title, ":")
		draftMR := statusMR[0]

		if p.Changes.Title.Current != p.Changes.Title.Previous {
			if draftMR != "Draft" {
				dt = fmt.Sprintf(mdl.StatusReadyMrMsg, p.ObjectAttributes.Title)
				url, text = p.ObjectAttributes.URL, "Open Request"
			} else {
				dt = fmt.Sprintf(mdl.StatusDraftMrMsg, p.ObjectAttributes.Title)
				url, text = p.ObjectAttributes.URL, "Open Request"
			}
		} else {
			if draftMR != "Draft" {
				dt = fmt.Sprintf(mdl.MergeRequestEventsMsg, p.User.Username, p.ObjectAttributes.SourceBranch, p.User.Username, p.Project.Name, p.Project.Homepage)
				url, text = p.ObjectAttributes.URL, "Open Request"
			}
		}
	case "pipeline":
		var p webhook.PipelineEventsLoad
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.PipelineEventsMsg, p.User.Username, p.ObjectAttributes.Ref, p.User.Username, p.Project.Name, p.Project.DefaultBranch, p.ObjectAttributes.Status)
		url, text = p.Project.WebURL, "Open Request"
	case "deployment":
		var p webhook.DeploymentEventsLoad
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.DeployEventsMsg, p.User.Username, p.Project.DefaultBranch, p.User.Username, p.Project.Name, p.Project.Homepage, p.Status)
		url, text = p.Project.WebURL, "Open Deloyment"
	case "release":
		var p webhook.ReleaseEventsLoad
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.ReleaseEventsMsg, p.Commit.Author.Name, p.Project.DefaultBranch, p.Commit.Author.Name, p.Project.Name, p.Project.Homepage, p.Commit.Message)
		url, text = p.Project.WebURL, "Open Release Event"
	case "wiki_page":
		var p webhook.WikipageEventsLoad
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.WikipageEventsMsg, p.User.Username, p.ObjectAttributes.Title)
		url, text = p.ObjectAttributes.URL, "Open WikiPage"
	case "tag_push":
		var p webhook.TagEventsLoad
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.TagEventsMsg, p.UserName, p.Ref)
		url, text = p.Project.WebURL, "Open WikiTag"
	case "build":
		var p webhook.JobsEvent
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.JobsEvent, p.BuildName, p.Ref, p.BuildStatus)
		url, text = p.Repository.Homepage, "Open Repository"
	case "feature_flag":
		var p webhook.FeatureFlag
		err = json.Unmarshal(body, &p)
		var s string
		if p.ObjectAttributes.Active {
			s = "active"
		} else {
			s = "unactive"
		}
		dt = fmt.Sprintf(mdl.FeatFlagMsg, p.ObjectAttributes.Name, s)
		url, text = p.Project.Homepage, "Open Project"
	default:
		log.Fatalf("Invalid Event\n")
	}

	if err != nil {
		log.Fatalf("Json unmarshal error: , %v\n", err)
	}

	msg := tgbot.NewMessage(chatId, dt)
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = tgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbot.InlineKeyboardButton{
			{
				tgbot.InlineKeyboardButton{
					Text: text,
					URL:  &url,
				},
			},
		},
	}
	h.Bot.Send(msg)
}
