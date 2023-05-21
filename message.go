package main

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	Greet = &i18n.Message{
		ID:          "Greet",
		Description: "greet",
		Other:       "こんにちわ {{.Name}}.",
	}
	AskName = &i18n.Message{
		ID:          "AskName",
		Description: "what your name",
		Other:       "あなたの名前はなんですか？",
	}
)
