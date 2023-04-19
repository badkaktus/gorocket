package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Alias       string       `json:"alias,omitempty"`
	Avatar      string       `json:"avatar,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Emoji       string       `json:"emoji,omitempty"`
	RoomID      string       `json:"roomId,omitempty"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	AudioURL          string        `json:"audio_url,omitempty"`
	AuthorIcon        string        `json:"author_icon,omitempty"`
	AuthorLink        string        `json:"author_link,omitempty"`
	AuthorName        string        `json:"author_name,omitempty"`
	Collapsed         bool          `json:"collapsed,omitempty"`
	Color             string        `json:"color,omitempty"`
	Fields            []AttachField `json:"fields,omitempty"`
	ImageURL          string        `json:"image_url,omitempty"`
	MessageLink       string        `json:"message_link,omitempty"`
	Text              string        `json:"text,omitempty"`
	ThumbURL          string        `json:"thumb_url,omitempty"`
	Title             string        `json:"title,omitempty"`
	TitleLink         string        `json:"title_link,omitempty"`
	TitleLinkDownload bool          `json:"title_link_download,omitempty"`
	Ts                time.Time     `json:"ts,omitempty"`
	VideoURL          string        `json:"video_url,omitempty"`
}

type AttachField struct {
	Short bool   `json:"short,omitempty"`
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
}

type RespPostMessage struct {
	ErrStatus
	Ts      int64           `json:"ts"`
	Channel string          `json:"channel"`
	Message RespMessageData `json:"message"`
}

type RespMessageData struct {
	Alias     string    `json:"alias,omitempty"`
	Msg       string    `json:"msg,omitempty"`
	ParseUrls bool      `json:"parseUrls,omitempty"`
	Groupable bool      `json:"groupable,omitempty"`
	Ts        time.Time `json:"ts,omitempty"`
	U         UChat     `json:"u,omitempty"`
	Rid       string    `json:"rid,omitempty"`
	UpdatedAt time.Time `json:"_updatedAt,omitempty"`
	ID        string    `json:"_id,omitempty"`
}

type UChat struct {
	ID       string `json:"_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type SingleMessageId struct {
	MessageId string `json:"messageId"`
}

type GetMessageResponse struct {
	ErrStatus
	Message MessageResp `json:"message"`
}

type MessageResp struct {
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
	ErrStatus
	ID string `json:"_id"`
	Ts int64  `json:"ts"`
}

type GetPinnedMsgRequest struct {
	RoomId string
	Count  int
	Offset int
}

type GetPinnedMsgResponse struct {
	ErrStatus
	Messages []PinnedMessage `json:"messages"`
	Count    int             `json:"count"`
	Offset   int             `json:"offset"`
	Total    int             `json:"total"`
}

type PinnedMessage struct {
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
	ErrStatus
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
}

// PostMessage posts a new chat message.
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

// GetMessage retrieves a single chat message by the provided id.
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

// DeleteMessage deletes a chat message by ID.
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

// GetPinnedMessages retrieves pinned messages from a room. Callee must have permission to access the room where the message resides.
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
		url.Add("offset", strconv.Itoa(param.Offset))
	}
	if param.Count != 0 {
		url.Add("count", strconv.Itoa(param.Count))
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

// PinMessage pins a chat message to the message's channel.
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

// UnpinMessage unpins a chat message from the message's channel.
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
