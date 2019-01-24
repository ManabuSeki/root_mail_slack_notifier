package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"strings"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

// Config is configuration struct
type Config struct {
	WebhookURL string `json:"WebhookURL"`
	Username   string `json:"Username"`
	Channel    string `json:"Channel"`
	IconEmoji  string `json:"IconEmoji"`
	Color      string `json:"Color"`
	HostName   string `json:"HostName"`
}

// flag
var (
	s string
)

func main() {
	flag.StringVar(&s, "config", "/etc/postfix/slack_notice.json", "config filepath")
	flag.Parse()

	// JSON Input
	jsonString, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	if err := json.Unmarshal(jsonString, &config); err != nil {
		log.Fatal(err)
	}
	// WebhookURL
	if config.WebhookURL == "" {
		log.Fatal("WebhookURL is not setting")
	}

	// Username
	if config.Username == "" {
		config.Username = "RootMail"
	}

	// Channel
	if config.Channel == "" {
		config.Channel = "#general"
	}

	// IconEmoji
	if config.IconEmoji == "" {
		config.IconEmoji = ":email:"
	}

	// HostName
	if config.HostName == "" {
		config.HostName, _ = os.Hostname()
	}
	mailInput, _ := ioutil.ReadAll(os.Stdin)
	m, err := mail.ReadMessage(strings.NewReader(string(mailInput)))

	if err != nil {
		log.Fatal(err)
	}
	header := m.Header
	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}

	txt := "*ホスト:*\n" +
		config.HostName + "\n" +
		"*件名:*\n" +
		"```" +
		header.Get("Subject") +
		"```\n\n" +
		"*本文:*\n" +
		"```" +
		string(body) +
		"```"
	attachment := slack.Attachment{Text: &txt, Color: &config.Color}
	payload := slack.Payload{
		Text:        "*Received an email addressed to root.*",
		Username:    config.Username,
		Channel:     config.Channel,
		IconEmoji:   config.IconEmoji,
		Attachments: []slack.Attachment{attachment},
	}

	errs := slack.Send(config.WebhookURL, "", payload)

	if len(errs) > 0 {
		fmt.Printf("error: %s\n", errs)
	}
}
