package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SimpleGroupId struct {
	RoomId string `json:"roomId,omitempty"`
}

type GroupCountersRequest struct {
	RoomId   string
	RoomName string
	UserId   string
}

type GroupCountersResponse struct {
	ErrStatus
	Joined       bool      `json:"joined"`
	Members      int       `json:"members"`
	Unreads      int       `json:"unreads"`
	UnreadsFrom  time.Time `json:"unreadsFrom"`
	Msgs         int       `json:"msgs"`
	Latest       time.Time `json:"latest"`
	UserMentions int       `json:"userMentions"`
}

type CreateGroupRequest struct {
	Name     string   `json:"name"`
	Members  []string `json:"members,omitempty"`
	ReadOnly bool     `json:"readOnly,omitempty"`
}

type CreateGroupResponse struct {
	ErrStatus
	Group Channel `json:"group"`
}

type SimpleGroupRequest struct {
	RoomId   string `json:"roomId,omitempty"`
	RoomName string `json:"roomName,omitempty"`
}

type GroupInfoResponse struct {
	ErrStatus
	Group groupInfo `json:"group"`
}

type groupInfo struct {
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

type InviteGroupRequest struct {
	RoomId string `json:"roomId"`
	UserId string `json:"userId"`
}

type InviteGroupResponse struct {
	ErrStatus
	Group struct {
		ID        string    `json:"_id"`
		Ts        time.Time `json:"ts"`
		T         string    `json:"t"`
		Name      string    `json:"name"`
		Usernames []string  `json:"usernames"`
		Msgs      int       `json:"msgs"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Lm        time.Time `json:"lm"`
	} `json:"group"`
}

type GroupListResponse struct {
	ErrStatus
	Groups []groupList `json:"groups"`
	Offset int         `json:"offset"`
	Count  int         `json:"count"`
	Total  int         `json:"total"`
}

type groupList struct {
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

type GroupMembersResponse struct {
	ErrStatus
	Members []Member `json:"members"`
	Count   int      `json:"count"`
	Offset  int      `json:"offset"`
	Total   int      `json:"total"`
}

type GroupMessage struct {
	ID        string        `json:"_id"`
	Rid       string        `json:"rid"`
	Msg       string        `json:"msg"`
	Ts        time.Time     `json:"ts"`
	U         U             `json:"u"`
	UpdatedAt time.Time     `json:"_updatedAt"`
	Reactions []interface{} `json:"reactions"`
	Mentions  []U           `json:"mentions"`
	Channels  []interface{} `json:"channels"`
	Starred   []interface{} `json:"starred"`
}

type GroupMessagesResponse struct {
	ErrStatus
	Messages []GroupMessage `json:"messages"`
	Count    int            `json:"count"`
	Offset   int            `json:"offset"`
	Total    int            `json:"total"`
}

type RenameGroupRequest struct {
	RoomId  string `json:"roomId"`
	NewName string `json:"name"`
}

type RenameGroupResponse struct {
	ErrStatus
	Group groupList `json:"group"`
}

type AddGroupPermissionRequest struct {
	RoomId string `json:"roomId"`
	UserId string `json:"userId"`
}

// ArchiveGroup archives a group.
func (c *Client) ArchiveGroup(param *SimpleGroupId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.archive", c.baseURL, c.apiVersion),
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

// CloseGroup removes the group from the user's list of groups.
func (c *Client) CloseGroup(param *SimpleGroupId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.close", c.baseURL, c.apiVersion),
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

// GroupCounters gets group counters.
func (c *Client) GroupCounters(param *GroupCountersRequest) (*GroupCountersResponse, error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/groups.counters", c.baseURL, c.apiVersion),
		nil)

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

	res := GroupCountersResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateGroup creates a new group.
func (c *Client) CreateGroup(param *CreateGroupRequest) (*CreateGroupResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.create", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := CreateGroupResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteGroup deletes a group.
func (c *Client) DeleteGroup(param *SimpleGroupId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.delete", c.baseURL, c.apiVersion),
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

// GroupInfo gets group info.
func (c *Client) GroupInfo(param *SimpleGroupRequest) (*GroupInfoResponse, error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/groups.info", c.baseURL, c.apiVersion),
		nil)

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

	res := GroupInfoResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GroupInvite adds a user to the group.
func (c *Client) GroupInvite(param *InviteGroupRequest) (*InviteGroupResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.invite", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := InviteGroupResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GroupKick kicks a user from the group.
func (c *Client) GroupKick(param *InviteGroupRequest) (*InviteGroupResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.kick", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := InviteGroupResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GroupList get groups list
func (c *Client) GroupList() (*GroupListResponse, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/groups.list", c.baseURL, c.apiVersion),
		nil)

	if err != nil {
		return nil, err
	}

	res := GroupListResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GroupMembers gets group members
func (c *Client) GroupMembers(param *SimpleGroupRequest) (*GroupMembersResponse, error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/groups.members", c.baseURL, c.apiVersion),
		nil)

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

	res := GroupMembersResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GroupMessages gets group messages
func (c *Client) GroupMessages(param *SimpleGroupRequest) (*GroupMessagesResponse, error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/groups.messages", c.baseURL, c.apiVersion),
		nil)

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

	res := GroupMessagesResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// OpenGroup adds the group back to the user's list of groups.
func (c *Client) OpenGroup(param *SimpleGroupId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.open", c.baseURL, c.apiVersion),
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

// RenameGroup changes a group's name.
func (c *Client) RenameGroup(param *RenameGroupRequest) (*RenameGroupResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.rename", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := RenameGroupResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// AddLeaderGroup adds leader for the group.
func (c *Client) AddLeaderGroup(param *AddGroupPermissionRequest) (*SimpleSuccessResponse, error) {
	if param.UserId == "" && param.RoomId == "" {
		return nil, fmt.Errorf("False parameters")
	}

	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.addLeader", c.baseURL, c.apiVersion),
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

// AddOwnerGroup adds owner for the group.
func (c *Client) AddOwnerGroup(param *AddGroupPermissionRequest) (*SimpleSuccessResponse, error) {
	if param.UserId == "" && param.RoomId == "" {
		return nil, fmt.Errorf("False parameters")
	}

	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.addOwner", c.baseURL, c.apiVersion),
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

// SetAnnouncementGroup sets the announcement for the group.
func (c *Client) SetAnnouncementGroup(param *SetAnnouncementRequest) (*SetAnnouncementResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.setAnnouncement", c.baseURL, c.apiVersion),
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

// SetDescriptionGroup sets the description for the group.
func (c *Client) SetDescriptionGroup(param *SetDescriptionRequest) (*SetDescriptionResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.setDescription", c.baseURL, c.apiVersion),
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

// SetTopicGroup sets the topic for the group.
func (c *Client) SetTopicGroup(param *SetTopicRequest) (*SetTopicResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.setTopic", c.baseURL, c.apiVersion),
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

// UnarchiveGroup unarchives a group.
func (c *Client) UnarchiveGroup(param *SimpleGroupId) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/groups.unarchive", c.baseURL, c.apiVersion),
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
