package main

import (
	"testing"
)

func TestSend(t *testing.T) {
	// key 값을 정상적인 값으로 바꾸면 하단 if err != nil 활성화
	err := Send("slack", "https://hooks.slack.com/services/key1/key2/key3", "Hello")
	want := "sending message status code is not ok"

	if err.Error() != want {
		t.Errorf("%v is not equal %v", err.Error(), want)
	}

	// if err != nil {
	// 	t.Error("send message fail")
	// }
}