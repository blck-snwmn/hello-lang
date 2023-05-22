package main

import (
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var languages = []language.Tag{language.Japanese, language.English}
var matcher = language.NewMatcher(languages)

func Test_ParseTag(t *testing.T) {
	type args struct {
		al string
	}
	tests := []struct {
		name string
		args args
		want language.Tag
	}{
		{
			name: "ja",
			args: args{al: "ja"},
			want: language.Japanese,
		},
		{
			name: "ja-JP",
			args: args{al: "ja-JP"},
			want: language.Japanese,
		},
		{
			name: "en",
			args: args{al: "en"},
			want: language.English,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, i := language.MatchStrings(matcher, tt.args.al)

			if got := languages[i]; got != tt.want {
				t.Errorf("xxx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Message(t *testing.T) {
	err := message.SetString(language.Japanese, "hello", "こんにちわ")
	if err != nil {
		t.Fatal(err)
	}
	err = message.SetString(language.English, "hello", "Hello")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		al string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ja",
			args: args{al: "ja"},
			want: "こんにちわ",
		},
		{
			name: "ja-JP",
			args: args{al: "ja-JP"},
			want: "こんにちわ",
		},
		{
			name: "en",
			args: args{al: "en"},
			want: "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag, _ := language.MatchStrings(matcher, tt.args.al)
			p := message.NewPrinter(tag)
			if got := p.Sprintf("hello"); got != tt.want {
				t.Errorf("xxx() = %v, want %v", got, tt.want)
			}
		})
	}
}
