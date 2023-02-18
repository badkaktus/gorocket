package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LoginPayload struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Resume   string `json:"resume,omitempty"`
}

type LoginResponse struct {
	ErrStatus
	Data    DataLogin `json:"data"`
	Message string    `json:"message,omitempty"`
}

type DataLogin struct {
	UserID    string `json:"userId"`
	AuthToken string `json:"authToken"`
	Me        Me     `json:"me"`
}

type Me struct {
	ID                    string    `json:"_id"`
	Services              Services  `json:"services"`
	Emails                []Email   `json:"emails"`
	Status                string    `json:"status"`
	Active                bool      `json:"active"`
	UpdatedAt             time.Time `json:"_updatedAt"`
	Roles                 []string  `json:"roles"`
	Name                  string    `json:"name"`
	StatusConnection      string    `json:"statusConnection"`
	Username              string    `json:"username"`
	UtcOffset             float64   `json:"utcOffset"`
	StatusText            string    `json:"statusText"`
	Settings              Settings  `json:"settings"`
	AvatarOrigin          string    `json:"avatarOrigin"`
	RequirePasswordChange bool      `json:"requirePasswordChange"`
	Language              string    `json:"language"`
	Email                 string    `json:"email"`
	AvatarURL             string    `json:"avatarUrl"`
}

type Services struct {
	Password Password `json:"password"`
}

type Password struct {
	Bcrypt string `json:"bcrypt"`
}

type Email struct {
	Address  string `json:"address"`
	Verified bool   `json:"verified"`
}

type Settings struct {
	Preferences Preferences `json:"preferences"`
}

type DontAskAgainList struct {
	Action string `json:"action"`
	Label  string `json:"label"`
}

type Preferences struct {
	EnableAutoAway                        bool               `json:"enableAutoAway"`
	IdleTimeLimit                         int                `json:"idleTimeLimit"`
	AudioNotifications                    string             `json:"audioNotifications"`
	DesktopNotifications                  string             `json:"desktopNotifications"`
	MobileNotifications                   string             `json:"mobileNotifications"`
	UnreadAlert                           bool               `json:"unreadAlert"`
	UseEmojis                             bool               `json:"useEmojis"`
	ConvertASCIIEmoji                     bool               `json:"convertAsciiEmoji"`
	AutoImageLoad                         bool               `json:"autoImageLoad"`
	SaveMobileBandwidth                   bool               `json:"saveMobileBandwidth"`
	CollapseMediaByDefault                bool               `json:"collapseMediaByDefault"`
	HideUsernames                         bool               `json:"hideUsernames"`
	HideRoles                             bool               `json:"hideRoles"`
	HideFlexTab                           bool               `json:"hideFlexTab"`
	HideAvatars                           bool               `json:"hideAvatars"`
	SidebarGroupByType                    bool               `json:"sidebarGroupByType"`
	SidebarViewMode                       string             `json:"sidebarViewMode"`
	SidebarHideAvatar                     bool               `json:"sidebarHideAvatar"`
	SidebarShowUnread                     bool               `json:"sidebarShowUnread"`
	SidebarShowFavorites                  bool               `json:"sidebarShowFavorites"`
	SendOnEnter                           string             `json:"sendOnEnter"`
	MessageViewMode                       int                `json:"messageViewMode"`
	EmailNotificationMode                 string             `json:"emailNotificationMode"`
	NewRoomNotification                   string             `json:"newRoomNotification"`
	NewMessageNotification                string             `json:"newMessageNotification"`
	MuteFocusedConversations              bool               `json:"muteFocusedConversations"`
	NotificationsSoundVolume              int                `json:"notificationsSoundVolume"`
	SidebarShowDiscussion                 bool               `json:"sidebarShowDiscussion"`
	DesktopNotificationRequireInteraction bool               `json:"desktopNotificationRequireInteraction"`
	SidebarSortby                         string             `json:"sidebarSortby"`
	DesktopNotificationDuration           int                `json:"desktopNotificationDuration"`
	DontAskAgainList                      []DontAskAgainList `json:"dontAskAgainList"`
	Highlights                            []interface{}      `json:"highlights"`
	Language                              string             `json:"language"`
}

type LogoutResponse struct {
	ErrStatus
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

type MeResponse struct {
	ErrStatus
	ID                    string    `json:"_id"`
	Services              Services  `json:"services"`
	Emails                []Email   `json:"emails"`
	Status                string    `json:"status"`
	Active                bool      `json:"active"`
	UpdatedAt             time.Time `json:"_updatedAt"`
	Roles                 []string  `json:"roles"`
	Name                  string    `json:"name"`
	StatusConnection      string    `json:"statusConnection"`
	Username              string    `json:"username"`
	UtcOffset             float64   `json:"utcOffset"`
	StatusText            string    `json:"statusText"`
	Settings              Settings  `json:"settings"`
	AvatarOrigin          string    `json:"avatarOrigin"`
	RequirePasswordChange bool      `json:"requirePasswordChange"`
	Language              string    `json:"language"`
	Email                 string    `json:"email"`
	AvatarURL             string    `json:"avatarUrl"`
}

func (c *Client) Login(login *LoginPayload) (*LoginResponse, error) {

	opt, _ := json.Marshal(login)
	url := fmt.Sprintf("%s/%s/login", c.baseURL, c.apiVersion)

	req, err := http.NewRequest("POST",
		url,
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := LoginResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	// success auth
	if res.Status == "success" {
		c.userID = res.Data.UserID
		c.xToken = res.Data.AuthToken
	}

	return &res, nil
}

func (c *Client) Logout() (*LogoutResponse, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/logout", c.baseURL, c.apiVersion), nil)
	if err != nil {
		return nil, err
	}

	res := LogoutResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Me() (*MeResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/me", c.baseURL, c.apiVersion), nil)
	if err != nil {
		return nil, err
	}

	res := MeResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
