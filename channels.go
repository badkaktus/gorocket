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
	Channel Channel `json:"channel"`
	Success bool    `json:"success"`
}

type Channel struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	T         string    `json:"t"`
	Usernames []string  `json:"usernames"`
	Msgs      int       `json:"msgs"`
	U         U         `json:"u"`
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
	Channel Channel `json:"channel"`
	Success bool    `json:"success"`
}

type SimpleChannelRequest struct {
	RoomId   string `json:"roomId,omitempty"`
	RoomName string `json:"roomName,omitempty"`
}

type ChannelHistoryRequest struct {
	RoomId    string
	Latest    time.Time
	Oldest    time.Time
	Inclusive bool
	Offset    int
	Count     int
	Unreads   bool
}

type ChannelInfoResponse struct {
	Channel ChannelInfo `json:"channel"`
	Success bool        `json:"success"`
}

type ChannelInfo struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	Fname        string `json:"fname"`
	T            string `json:"t"`
	Msgs         int    `json:"msgs"`
	UsersCount   int    `json:"usersCount"`
	U            UChat  `json:"u"`
	CustomFields struct {
	} `json:"customFields"`
	Broadcast bool      `json:"broadcast"`
	Encrypted bool      `json:"encrypted"`
	Ts        time.Time `json:"ts"`
	Ro        bool      `json:"ro"`
	Default   bool      `json:"default"`
	SysMes    bool      `json:"sysMes"`
	UpdatedAt time.Time `json:"_updatedAt"`
}

type InviteChannelRequest struct {
	RoomId string `json:"roomId"`
	UserId string `json:"userId"`
}

type InviteChannelResponse struct {
	Channel struct {
		ID        string    `json:"_id"`
		Ts        time.Time `json:"ts"`
		T         string    `json:"t"`
		Name      string    `json:"name"`
		Usernames []string  `json:"usernames"`
		Msgs      int       `json:"msgs"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Lm        time.Time `json:"lm"`
	} `json:"channel"`
	Success bool `json:"success"`
}

type ChannelListResponse struct {
	Channels []ChannelList `json:"channels"`
	Offset   int           `json:"offset"`
	Count    int           `json:"count"`
	Total    int           `json:"total"`
	Success  bool          `json:"success"`
}

type ChannelList struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	T         string    `json:"t"`
	Usernames []string  `json:"usernames"`
	Msgs      int       `json:"msgs"`
	U         UChat     `json:"u"`
	Ts        time.Time `json:"ts"`
	Ro        bool      `json:"ro"`
	SysMes    bool      `json:"sysMes"`
	UpdatedAt time.Time `json:"_updatedAt"`
}

type ChannelMembersResponse struct {
	Members []Member `json:"members"`
	Count   int      `json:"count"`
	Offset  int      `json:"offset"`
	Total   int      `json:"total"`
	Success bool     `json:"success"`
}

type Member struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}

type RenameChannelRequest struct {
	RoomId  string `json:"roomId"`
	NewName string `json:"name"`
}

type RenameChannelResponse struct {
	Channel ChannelList `json:"channel"`
	Success bool        `json:"success"`
}

type SetAnnouncementRequest struct {
	RoomId       string `json:"roomId"`
	Announcement string `json:"announcement"`
}

type SetAnnouncementResponse struct {
	Announcement string `json:"announcement"`
	Success      bool   `json:"success"`
}

type SetDescriptionRequest struct {
	RoomId      string `json:"roomId"`
	Description string `json:"description"`
}

type SetDescriptionResponse struct {
	Description string `json:"description"`
	Success     bool   `json:"success"`
}

type SetTopicRequest struct {
	RoomId string `json:"roomId"`
	Topic  string `json:"topic"`
}

type SetTopicResponse struct {
	Topic   string `json:"topic"`
	Success bool   `json:"success"`
}

// AddAllToChannel adds all of the users on the server to a channel.
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

// ArchiveChannel archives a channel.
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

// CloseChannel removes the channel from the user's list of channels.
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

// ChannelCounters gets channel counters.
func (c *Client) ChannelCounters(param *ChannelCountersRequest) (*ChannelCountersResponse, error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/channels.counters", c.baseURL, c.apiVersion),
		nil)

	if param.RoomName == "" && param.RoomId == "" {
		return nil, fmt.Errorf("false parameters")
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

// CreateChannel creates a new channel.
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

// DeleteChannel delete channel.
func (c *Client) DeleteChannel(param *SimpleChannelRequest) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.delete", c.baseURL, c.apiVersion),
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

// ChannelInfo get channel info.
func (c *Client) ChannelInfo(param *SimpleChannelRequest) (*ChannelInfoResponse, error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/channels.info", c.baseURL, c.apiVersion),
		nil)

	if param.RoomName == "" && param.RoomId == "" {
		return nil, fmt.Errorf("false parameters")
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

	res := ChannelInfoResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ChannelInvite adds a user to the channel.
func (c *Client) ChannelInvite(param *InviteChannelRequest) (*InviteChannelResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.invite", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := InviteChannelResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ChannelKick kick a user from the channel.
func (c *Client) ChannelKick(param *InviteChannelRequest) (*InviteChannelResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.kick", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := InviteChannelResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ChannelList get channels list
func (c *Client) ChannelList() (*ChannelListResponse, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/channels.list", c.baseURL, c.apiVersion),
		nil)

	if err != nil {
		return nil, err
	}

	res := ChannelListResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ChannelMembers gets channel members
func (c *Client) ChannelMembers(param *SimpleChannelRequest) (*ChannelMembersResponse, error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/channels.members", c.baseURL, c.apiVersion),
		nil)

	if param.RoomName == "" && param.RoomId == "" {
		return nil, fmt.Errorf("false parameters")
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

	res := ChannelMembersResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// OpenChannel adds the channel back to the user's list of channels.
func (c *Client) OpenChannel(param *SimpleChannelId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.open", c.baseURL, c.apiVersion),
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

// RenameChannel changes a channel's name.
func (c *Client) RenameChannel(param *RenameChannelRequest) (*RenameChannelResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.rename", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := RenameChannelResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SetAnnouncementChannel sets the announcement for the channel.
func (c *Client) SetAnnouncementChannel(param *SetAnnouncementRequest) (*SetAnnouncementResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.setAnnouncement", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := SetAnnouncementResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SetDescriptionChannel sets the Description for the channel.
func (c *Client) SetDescriptionChannel(param *SetDescriptionRequest) (*SetDescriptionResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.setDescription", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := SetDescriptionResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SetTopicChannel sets the topic for the channel.
func (c *Client) SetTopicChannel(param *SetTopicRequest) (*SetTopicResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.setTopic", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := SetTopicResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// UnarchiveChannel unarchive a channel.
func (c *Client) UnarchiveChannel(param *SimpleChannelId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/channels.unarchive", c.baseURL, c.apiVersion),
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
