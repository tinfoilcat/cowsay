package slackbackend

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ashwanthkumar/slack-go-webhook"
)

var webhookUrl string
var sendSlackEnabled bool

func InitBackend() {
	var err error

	webhookUrl, err = LoadSlackWebhookUrl()
	checkSlackMessageEnable()

	if err != nil {
		log.Println("Error loading slack webhook\n" + err.Error())
	}
	fmt.Printf("Slack webhook url is: %v", webhookUrl)
}

func checkSlackMessageEnable() {
	_, ok := os.LookupEnv("COWSAY_SENDSLACK")
	
	if !ok {
		sendSlackEnabled = false
	} else {
		sendSlackEnabled = true
	}
}

func LoadSlackWebhookUrl() (string, error) {
	// akv_buddy -akvurl "https://terraform-ansible-keys.vault.azure.net" -akvsecret "slack-webhook-automation-pipeline"
	Cmd := exec.Command("/usr/bin/akv_buddy", "-akvurl", "https://terraform-ansible-keys.vault.azure.net", "-akvsecret", "slack-webhook-automation-pipeline")

	stdout, err := Cmd.Output()

	if err != nil {
		return "", err
	}
	return string(stdout), err

}

func SendSlackMessage(message string) {
	if sendSlackEnabled {
		payload := slack.Payload{
			Text:      message,
			Username:  "Terraform IaC Tester",
			Channel:   "#automation_pipeline",
			IconEmoji: ":gopherbatman:",
		}
		err := slack.Send(webhookUrl, "", payload)
		if len(err) > 0 {
			fmt.Printf("error: %s\n", err)
		}
	}
}
