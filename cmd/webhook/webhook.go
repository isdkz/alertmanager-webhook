package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	model "github.com/isdkz/alertmanager-webhook/model"
	"github.com/isdkz/alertmanager-webhook/notifier"
	"net/http"
)

var (
	h            bool
	defaultRobot string
	sport        string
	msgtype      string
	tplfile      string
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "u", "", "global robot webhook url, you can overwrite by alert rule with annotations Robot")
	flag.StringVar(&sport, "p", "8080", "port on which the webhook server runs")
	flag.StringVar(&msgtype, "t", "text", "type of the push message, text and md are supprted")
	flag.StringVar(&tplfile, "f", "send.tpl", "the template file name")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if robotUrl, ok := notification.CommonAnnotations["Robot"]; ok {
			defaultRobot = robotUrl
		}
		for _, alert := range notification.Alerts {
			err = notifier.Send(alert, defaultRobot, msgtype, tplfile)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			}

			c.JSON(http.StatusOK, gin.H{"message": "send successful!"})

		}
	})
	router.Run(":" + sport)
}
