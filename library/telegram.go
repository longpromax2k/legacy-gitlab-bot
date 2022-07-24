package lib

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/tatsuxyz/GitLabHook/model"
)

func PostMessage(pay model.Gitlab) {
	api := "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_TOKEN") + "/sendMessage"
	message := "<b>" + pay.UserUsername + "</b> pushed to branch <a href='" + pay.Commits[0].URL + "'>" + pay.Ref + "</a> of <a href='" + pay.Repository.URL + "'>" + pay.Project.Namespace + "</a> (<a href='" + pay.Commits[0].URL + "'>Compare changes</a>) &#013;	&gt; <a href='" + pay.Commits[0].URL + "'>" + strconv.Itoa(pay.ProjectID) + "</a>: " + pay.Commits[0].Message
	text := url.QueryEscape(message)

	url := api + "?chat_id=" + os.Getenv("CHAT_ID") + "&text=" + text + "&parse_mode=HTML"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
}
