package vk

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	Token   string
	GroupID string
}

func New(token, groupID string) *Client {
	return &Client{
		Token:   token,
		GroupID: groupID,
	}
}
func (c *Client) SendMessage(userID int, message string) error {
	url := "https://api.vk.com/method/messages.send"
	payload := map[string]any{
		"user_id": userID,
		"random_id": 0,
		"message": message,
		"access_token": c.Token,
		"v" : "5.199",
	}
	body, _ := json.Marshal(payload)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	return err
}
