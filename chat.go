package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Alias       string       `json:"alias"`
	Avatar      string       `json:"avatar"`
	Channel     string       `json:"channel"`
	Emoji       string       `json:"emoji"`
	RoomID      string       `json:"roomId"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	AudioURL          string        `json:"audio_url"`
	AuthorIcon        string        `json:"author_icon"`
	AuthorLink        string        `json:"author_link"`
	AuthorName        string        `json:"author_name"`
	Collapsed         bool          `json:"collapsed"`
	Color             string        `json:"color"`
	Fields            []AttachField `json:"fields"`
	ImageURL          string        `json:"image_url"`
	MessageLink       string        `json:"message_link"`
	Text              string        `json:"text"`
	ThumbURL          string        `json:"thumb_url"`
	Title             string        `json:"title"`
	TitleLink         string        `json:"title_link"`
	TitleLinkDownload bool          `json:"title_link_download"`
	Ts                time.Time     `json:"ts"`
	VideoURL          string        `json:"video_url"`
}

type AttachField struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type RespPostMessage struct {
	Ts      int64           `json:"ts"`
	Channel string          `json:"channel"`
	Message RespMessageData `json:"message"`
	Success bool            `json:"success"`
}

type RespMessageData struct {
	Alias     string    `json:"alias"`
	Msg       string    `json:"msg"`
	ParseUrls bool      `json:"parseUrls"`
	Groupable bool      `json:"groupable"`
	Ts        time.Time `json:"ts"`
	U         U         `json:"u"`
	Rid       string    `json:"rid"`
	UpdatedAt time.Time `json:"_updatedAt"`
	ID        string    `json:"_id"`
}

type U struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}

func (c *Client) PostMessage(msg *Message) (*RespPostMessage, error) {

	opt, _ := json.Marshal(msg)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/chat.postMessage", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := RespPostMessage{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
