package main

import (
	"encoding/json"
	"github-telegram-notify/types"
	"github-telegram-notify/utils"
	"os"
)

func main() {
	tg_token := os.Getenv("INPUT_BOT_TOKEN")
	chatID := os.Getenv("INPUT_CHAT_ID")
        threadID := os.Getenv("INPUT_THREAD_ID")
	gitEventRaw := os.Getenv("INPUT_GIT_EVENT")
	print(gitEventRaw)
	var gitEvent *types.Metadata
	err := json.Unmarshal([]byte(gitEventRaw), &gitEvent)
	if err != nil {
		panic(err)
	}
	text, markupText, markupUrl, err := utils.CreateContents(gitEvent)
	if err != nil {
		panic(err)
	}
	error := utils.SendMessage(tg_token, chatID, threadID, text, markupText, markupUrl)
	if error.Description != "" {
		panic(error.String())
	}

}
