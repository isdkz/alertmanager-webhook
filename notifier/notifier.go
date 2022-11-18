package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/isdkz/alertmanager-webhook/model"
	"github.com/isdkz/alertmanager-webhook/transformer"
)

// Send send markdown message to dingtalk
func Send(alert model.Alert, RobotUrl string, msgtype string, tplfile string) (err error) {
	message,  err := transformer.TransformToMessage(alert, msgtype, tplfile)

	if err != nil {
		return
	}
	data, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	if len(RobotUrl) == 0 {
		panic("Robot url not set")
	}
	req, err := http.NewRequest(
		"POST",
		RobotUrl,
		bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
