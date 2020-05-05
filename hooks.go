package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HookMessage struct {
	Text        string           `json:"text"`
	Attachments []HookAttachment `json:"attachments"`
}

type HookAttachment struct {
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
	ImageURL  string `json:"image_url"`
	Color     string `json:"color"`
}

func (c *Client) Hooks(msg *HookMessage, token string) {
	opt, _ := json.Marshal(msg)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/hooks/%s", c.baseURL, token),
		bytes.NewBuffer(opt))
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		log.Fatal("Request error")
		//return nil, err
	}

}
