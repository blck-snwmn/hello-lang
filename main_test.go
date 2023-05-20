package main

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"golang.org/x/text/message"
)

var matcher = language.NewMatcher([]language.Tag{language.Japanese, language.English})

func Test_ParseTag(t *testing.T) {
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
			want: "日本語",
		},
		{
			name: "ja-JP",
			args: args{al: "ja-JP"},
			want: "日本語",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tags, _, err := language.ParseAcceptLanguage(tt.args.al)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("%v\n", tags)
			tag, _, _ := matcher.Match(tags...)
			fmt.Printf("%v\n", tag)

			if got := display.Self.Name(tag); got != tt.want {
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
			want: "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tags, _, err := language.ParseAcceptLanguage(tt.args.al)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("%v\n", tags)
			tag, _, _ := matcher.Match(tags...)
			fmt.Printf("%v\n", tag)

			p := message.NewPrinter(tag)
			if got := p.Sprint("hello"); got != tt.want {
				t.Errorf("xxx() = %v, want %v", got, tt.want)
			}
		})
	}
}
