package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AddAllRequest struct {
	RoomId          string `json:"roomId"`
	ActiveUsersOnly bool   `json:"activeUsersOnly,omitempty"`
}

type AddAllResponse struct {
	Channel channel `json:"channel"`
	Success bool    `json:"success"`
}

type channel struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	T         string    `json:"t"`
	Usernames []string  `json:"usernames"`
	Msgs      int       `json:"msgs"`
	U         u         `json:"u"`
	Ts        time.Time `json:"ts"`
}

type SimpleChannelId struct {
	RoomId string `json:"roomId,omitempty"`
}

type ChannelCountersRequest struct {
	RoomId   string
	RoomName string
	UserId   string
}

type ChannelCountersResponse struct {
	Joined       bool      `json:"joined"`
	Members      int       `json:"members"`
	Unreads      int       `json:"unreads"`
	UnreadsFrom  time.Time `json:"unreadsFrom"`
	Msgs         int       `json:"msgs"`
	Latest       time.Time `json:"latest"`
	UserMentions int       `json:"userMentions"`
	Success      bool      `json:"success"`
}

type CreateChannelRequest struct {
	Name     string   `json:"name"`
	Members  []string `json:"members,omitempty"`
	ReadOnly bool     `json:"readOnly,omitempty"`
}

type CreateChannelResponse struct {
	Channel channel `json:"channel"`
	Success bool    `json:"success"`
}

// Adds all of the users on the server to a channel.
func (c *Client) AddAllToChannel(params *AddAllRequest) (*AddAllResponse, error) {
	opt, _ := json.Marshal(params)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.addAll", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := AddAllResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Archives a channel.
func (c *Client) ArchiveChannel(param *SimpleChannelId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.archive", c.baseURL, c.apiVersion),
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

// Removes the channel from the user's list of channels.
func (c *Client) CloseChannel(param *SimpleChannelId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.close", c.baseURL, c.apiVersion),
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

// Gets channel counters.
func (c *Client) ChannelCounters(param *ChannelCountersRequest) (*ChannelCountersResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/channels.counters", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if param.RoomName == "" && param.RoomId == "" {
		return nil, fmt.Errorf("False parameters")
	}

	url := req.URL.Query()
	if param.RoomName != "" {
		url.Add("roomName", param.RoomName)
	}
	if param.RoomId != "" {
		url.Add("roomId", param.RoomId)
	}
	req.URL.RawQuery = url.Encode()

	if err != nil {
		return nil, err
	}

	res := ChannelCountersResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Removes the channel from the user's list of channels.
func (c *Client) CreateChannel(param *CreateChannelRequest) (*CreateChannelResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.create", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := CreateChannelResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
