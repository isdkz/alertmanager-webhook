package transformer

import (
	"bytes"
	"text/template"
	"time"

	"github.com/isdkz/alertmanager-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to Message
func TransformToMessage(alert model.Alert, msgtype string, tplfile string) (message any, err error) {
	tmpl, err := template.ParseFiles(tplfile)
	if err != nil {
		return
	}
	buffer := new(bytes.Buffer)

	alert.StartsAt = alert.StartsAt.Add(time.Hour * 8)
	if alert.Status != "firing" {
		alert.EndsAt = alert.EndsAt.Add(time.Hour * 8)
	}

	err = tmpl.Execute(buffer, alert)
	if err != nil {
		return
	}
	switch msgtype {
	case "text":
		message = &model.QywxText{
			MsgType: "text",
			Text: &model.Content{
				Content: buffer.String(),
			},
		}
	case "md":
		message = &model.QywxMarkdown{
			MsgType: "markdown",
			Markdown: &model.Content{
				Content: buffer.String(),
			},
		}
	}

	return
}
