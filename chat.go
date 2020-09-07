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
	Ts        int64           `json:"ts"`
	Channel   string          `json:"channel"`
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

type SingleMessageId struct {
	MessageId string `json:"messageId"`
}

type GetMessageResponse struct {
	Message message `json:"message"`
	Success bool    `json:"success"`
}

type message struct {
	ID  string    `json:"_id"`
	Rid string    `json:"rid"`
	Msg string    `json:"msg"`
	Ts  time.Time `json:"ts"`
	U   struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
	} `json:"u"`
}

type DeleteMessageRequest struct {
	RoomID string `json:"roomId"`
	MsgID  string `json:"msgId"`
	AsUser bool   `json:"asUser,omitempty"`
}

type DeleteMessageResponse struct {
	ID      string `json:"_id"`
	Ts      int64  `json:"ts"`
	Success bool   `json:"success"`
}

type GetPinnedMsgRequest struct {
	RoomId string
	Count  int
	Offset int
}

type GetPinnedMsgResponse struct {
	Messages []pinnedMessage `json:"messages"`
	Count    int             `json:"count"`
	Offset   int             `json:"offset"`
	Total    int             `json:"total"`
	Success  bool            `json:"success"`
}

type pinnedMessage struct {
	ID  string    `json:"_id"`
	Rid string    `json:"rid"`
	Msg string    `json:"msg"`
	Ts  time.Time `json:"ts"`
	U   struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
		Name     string `json:"name"`
	} `json:"u"`
	Mentions  []interface{} `json:"mentions"`
	Channels  []interface{} `json:"channels"`
	UpdatedAt time.Time     `json:"_updatedAt"`
	Pinned    bool          `json:"pinned"`
	PinnedAt  time.Time     `json:"pinnedAt"`
	PinnedBy  struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
	} `json:"pinnedBy"`
}

type PinMessageResponse struct {
	Message struct {
		T   string    `json:"t"`
		Rid string    `json:"rid"`
		Ts  time.Time `json:"ts"`
		Msg string    `json:"msg"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Groupable   bool `json:"groupable"`
		Attachments []struct {
			Text       string    `json:"text"`
			AuthorName string    `json:"author_name"`
			AuthorIcon string    `json:"author_icon"`
			Ts         time.Time `json:"ts"`
		} `json:"attachments"`
		UpdatedAt time.Time `json:"_updatedAt"`
		ID        string    `json:"_id"`
	} `json:"message"`
	Success bool `json:"success"`
}

// Posts a new chat message.
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

// Retrieves a single chat message by the provided id.
// Callee must have permission to access the room where the message resides.
func (c *Client) GetMessage(param *SingleMessageId) (*GetMessageResponse, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/chat.getMessage", c.baseURL, c.apiVersion),
		nil)

	if param.MessageId == "" {
		return nil, fmt.Errorf("False parameters")
	}

	url := req.URL.Query()
	if param.MessageId != "" {
		url.Add("msgId", param.MessageId)
	}
	req.URL.RawQuery = url.Encode()

	if err != nil {
		return nil, err
	}

	res := GetMessageResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Chat Message Delete
func (c *Client) DeleteMessage(param *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/chat.delete", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := DeleteMessageResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Callee must have permission to access the room where the message resides.
func (c *Client) GetPinnedMessages(param *GetPinnedMsgRequest) (*GetPinnedMsgResponse, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/chat.getPinnedMessages", c.baseURL, c.apiVersion),
		nil)

	if param.RoomId == "" {
		return nil, fmt.Errorf("False parameters")
	}

	url := req.URL.Query()
	if param.RoomId != "" {
		url.Add("roomId", param.RoomId)
	}
	if param.Offset != 0 {
		url.Add("offset", string(param.Offset))
	}
	if param.Count != 0 {
		url.Add("count", string(param.Count))
	}
	req.URL.RawQuery = url.Encode()

	if err != nil {
		return nil, err
	}

	res := GetPinnedMsgResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Pins a chat message to the message's channel.
func (c *Client) PinMessage(param *SingleMessageId) (*PinMessageResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/chat.pinMessage", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := PinMessageResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Unpins a chat message to the message's channel.
func (c *Client) UnpinMessage(param *SingleMessageId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/chat.unPinMessage", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := SimpleSuccessResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
