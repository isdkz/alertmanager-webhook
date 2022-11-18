package transformer

import (
	"bytes"
	"github.com/isdkz/alertmanager-webhook/model"
	"text/template"
	"time"
)

// TransformToMarkdown transform alertmanager notification to Message
func TransformToMessage(alert model.Alert, msgtype string, tplfile string) (message interface{}, err error) {
	tmpl, err := template.ParseFiles(tplfile)
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)

	alert.StartsAt = alert.StartsAt.Add(time.Hour * 8)
	if alert.Status != "firing" {
		alert.EndsAt = alert.EndsAt.Add(time.Hour * 8)
	}

	err = tmpl.Execute(buffer, alert)
	if err != nil {
		panic(err)
	}
	switch msgtype {
	case "text":
		message = &model.QywxText{
			MsgType: "text",
			Text: &model.Content{
				Content: buffer.String(),
			},
		}
	case "markdown":
		message = &model.QywxMarkdown{
			MsgType: "markdown",
			Markdown: &model.Content{
				Content: buffer.String(),
			},
		}
	}

	return
}
