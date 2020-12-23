package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"alertmanger-feishu-webhook/config"
	"alertmanger-feishu-webhook/utils"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func checkPrefix(str string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

func MsgText(alert *Alert) string {
	graphURL := utils.PromGraphURL(alert.GeneratorURL, config.Args.PrometheusExternalDomain)
	if alert.Status == "firing" {
		alert.Status = "发生中"
	}
	if alert.Status == "resolved" {
		alert.Status = "已解决"
	}
	if alert.EndsAt.IsZero() {
		return fmt.Sprintf("告警状态: %v\n告警信息: %v\n开始时间: %v\n查看详情: %v",
			alert.Status, alert.Annotations["description"], alert.StartsAt, graphURL)
	} else {
		return fmt.Sprintf("告警状态: %v\n告警信息: %v\n开始时间: %v\n结束时间: %v\n查看详情: %v",
			alert.Status, alert.Annotations["description"], alert.StartsAt, alert.EndsAt, graphURL)
	}
}

func SendMsgToWebhook(notification *Notification) {
	var msg = &Message{}
	msg.Title = "prometheus 告警信息"
	var alerts = notification.Alerts
	for _, alert := range alerts {
		webhook := config.Args.DefaultAlertWebhook
		switch {
		case checkPrefix(alert.Labels["alertname"], "IstioEnvoy", "Nginx"):
			if alert.Labels["source_workload_namespace"] == "glzh-pre" {
				webhook = config.Args.PreAlertWebhook
			} else if config.Args.DomainAlertWebhook != "" {
				webhook = config.Args.DomainAlertWebhook
			}
		case checkPrefix(alert.Labels["alertname"], "App"):
			if alert.Labels["namespace"] == "glzh-pre" {
				webhook = config.Args.PreAlertWebhook
			} else {
				switch {
				case alert.Labels["label_group"] == "frontend":
					if config.Args.FrontendAppAlertWebhook != "" {
						webhook = config.Args.FrontendAppAlertWebhook
					}
				case alert.Labels["label_group"] == "backend":
					if config.Args.BackendAppAlertWebhook != "" {
						webhook = config.Args.BackendAppAlertWebhook
					}
				default:
					webhook = config.Args.DefaultAlertWebhook
				}
			}
		default:
			webhook = config.Args.DefaultAlertWebhook
		}
		msg.Text = MsgText(&alert)
		sendMsg, _ := json.Marshal(msg)
		req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(sendMsg))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			log.Printf("post feishu webhook error, error is %v", err.Error())
		}
		log.Printf("post feishu %v alert status: %v\n", alert.Annotations["description"], resp.Status)
	}
}
