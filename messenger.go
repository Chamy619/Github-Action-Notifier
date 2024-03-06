package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type SlackBody struct {
	Text string `json:"text"`
}

func Send(messengerType string, url string, message string) error {
	switch messengerType {
	case "slack":
		messageData := SlackBody{message}

		body, err := json.Marshal(messageData)
		if err != nil {
			return errors.New("marshaling fail")
		}

		response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			return errors.New("sending message fail")
		}
		if response.StatusCode != 200 {
			return errors.New("sending message status code is not ok")
		}
		return nil
	default:
		return errors.New("unsupported messenger type")
	}
}