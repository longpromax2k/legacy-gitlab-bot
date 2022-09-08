package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	mdl "gitbot/models"
	webhook "gitbot/models/webhook"
	cmt "gitbot/models/webhook/comment"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var previousTestID int
var previousBuildID int
var previousDeployID int
var typeJob string
var statusJob string
var statusPipeline string
var previousPipeline int
var currentMesID int
var count int

func SendTelegramMessage(pay mdl.ObjectKind, body []byte, cId string) {
	cid, _ := strconv.Atoi(cId)
	var chatId = int64(cid)
	var dt, url, text string
	var err error
	bot, err := LoadBot()

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
		statusPipeline = p.ObjectAttributes.Status
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
		typeJob = p.BuildStage
		statusJob = p.BuildStatus
		url, text = p.Repository.Homepage, ""
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
	msg1 := tgbot.NewEditMessageText(chatId, currentMesID, dt)
	msg1.ParseMode = "markdown"
	msg1.ReplyMarkup = &tgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbot.InlineKeyboardButton{
			{
				tgbot.InlineKeyboardButton{
					Text: text,
					URL:  &url,
				},
			},
		},
	}
	//var p webhook.JobsEvent

	log.Println(pay.ObjectKind, "1")
	if pay.ObjectKind == "build" {
		if typeJob == "test" && statusJob != "created" {
			msg1 = tgbot.NewEditMessageText(chatId, currentMesID, dt)
			bot.Send(msg1)
			//previousTestID=m1.MessageID
		}
		if typeJob == "build" && statusJob != "created" {
			msg1 = tgbot.NewEditMessageText(chatId, currentMesID, dt)
			bot.Send(msg1)
			//previousBuildID=m1.MessageID
		}
		if typeJob == "deploy" && statusJob != "created" {
			msg1 = tgbot.NewEditMessageText(chatId, currentMesID, dt)
			bot.Send(msg1)
			//previousDeployID=m1.MessageID
		}
		if statusJob == "created" {

			count++
			log.Println("hello", count)
			if count%3 == 1 {
				m1, _ := bot.Send(msg)
				currentMesID = m1.MessageID
			}
		}
		// if typeJob == "test" && statusJob == "created" {
		// 	m1, _ := bot.Send(msg)
		// 	previousTestID = m1.MessageID
		// }
		// if typeJob == "build" && statusJob == "created" {
		// 	m1, _ := bot.Send(msg)
		// 	previousBuildID = m1.MessageID
		// }
		// if typeJob == "deploy" && statusJob == "created" {
		// 	m1, _ := bot.Send(msg)
		// 	previousDeployID = m1.MessageID
		// }
	} else if pay.ObjectKind == "pipeline" {
		log.Println(statusPipeline)
		if statusPipeline == "pending" {
			m1, _ := bot.Send(msg)
			previousPipeline = m1.MessageID
		} else {
			msg1 = tgbot.NewEditMessageText(chatId, previousPipeline, dt)
			bot.Send(msg1)
		}

	} else {
		bot.Send(msg)
		//currentMesID = m1.MessageID
	}
	//bot.Send(msg)
	//return
}

//func EditTelegramMessgae()
