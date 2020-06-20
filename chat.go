package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Alias       string       `json:"alias,omitempty"`
	Avatar      string       `json:"avatar,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Emoji       string       `json:"emoji,omitempty"`
	RoomID      string       `json:"roomId,omitempty"`
	Text        string       `json:"text"`
	Attachments []attachment `json:"attachments"`
}

type attachment struct {
	AudioURL          string        `json:"audio_url"`
	AuthorIcon        string        `json:"author_icon"`
	AuthorLink        string        `json:"author_link"`
	AuthorName        string        `json:"author_name"`
	Collapsed         bool          `json:"collapsed"`
	Color             string        `json:"color"`
	Fields            []attachField `json:"fields"`
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

type attachField struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type RespPostMessage struct {
	Ts      int64  `json:"ts"`
	Channel string `json:"channel"`
	// TODO когда нет авторизации (заголовков), то поле Message имеет тип string и поэтому возникает ошибка, надо разобраться
	Message   respMessageData `json:"message"`
	Success   bool            `json:"success"`
	Error     string          `json:"error,omitempty"`
	ErrorType string          `json:"errorType,omitempty"`
}

type respMessageData struct {
	Alias     string    `json:"alias"`
	Msg       string    `json:"msg"`
	ParseUrls bool      `json:"parseUrls"`
	Groupable bool      `json:"groupable"`
	Ts        time.Time `json:"ts"`
	U         uChat     `json:"u"`
	Rid       string    `json:"rid"`
	UpdatedAt time.Time `json:"_updatedAt"`
	ID        string    `json:"_id"`
}

type uChat struct {
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
	fmt.Printf("\nHeaders:\n")
	// Loop over header names
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("\nBody:\n%+v\n", req.GetBody)
	res := RespPostMessage{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
