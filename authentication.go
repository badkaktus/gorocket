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
	Status  string `json:"status"`
	Data    data   `json:"data"`
	Message string `json:"message,omitempty"`
}

type data struct {
	UserID    string `json:"userId"`
	AuthToken string `json:"authToken"`
	Me        me     `json:"me"`
}

type me struct {
	ID                    string    `json:"_id"`
	Services              services  `json:"services"`
	Emails                []email   `json:"emails"`
	Status                string    `json:"status"`
	Active                bool      `json:"active"`
	UpdatedAt             time.Time `json:"_updatedAt"`
	Roles                 []string  `json:"roles"`
	Name                  string    `json:"name"`
	StatusConnection      string    `json:"statusConnection"`
	Username              string    `json:"username"`
	UtcOffset             int       `json:"utcOffset"`
	StatusText            string    `json:"statusText"`
	Settings              settings  `json:"settings"`
	AvatarOrigin          string    `json:"avatarOrigin"`
	RequirePasswordChange bool      `json:"requirePasswordChange"`
	Language              string    `json:"language"`
	Email                 string    `json:"email"`
	AvatarURL             string    `json:"avatarUrl"`
}

type services struct {
	Password password `json:"password"`
}

type password struct {
	Bcrypt string `json:"bcrypt"`
}

type email struct {
	Address  string `json:"address"`
	Verified bool   `json:"verified"`
}

type settings struct {
	Preferences preferences `json:"preferences"`
}

type dontAskAgainList struct {
	Action string `json:"action"`
	Label  string `json:"label"`
}

type preferences struct {
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
	DontAskAgainList                      []dontAskAgainList `json:"dontAskAgainList"`
	Highlights                            []interface{}      `json:"highlights"`
	Language                              string             `json:"language"`
}

type LogoutResponse struct {
	Status string `json:"status"`
	Data   struct {
		Message string `json:"message"`
	} `json:"data"`
}

type MeResponse struct {
	ID                    string    `json:"_id"`
	Services              services  `json:"services"`
	Emails                []email   `json:"emails"`
	Status                string    `json:"status"`
	Active                bool      `json:"active"`
	UpdatedAt             time.Time `json:"_updatedAt"`
	Roles                 []string  `json:"roles"`
	Name                  string    `json:"name"`
	StatusConnection      string    `json:"statusConnection"`
	Username              string    `json:"username"`
	UtcOffset             int       `json:"utcOffset"`
	StatusText            string    `json:"statusText"`
	Settings              settings  `json:"settings"`
	AvatarOrigin          string    `json:"avatarOrigin"`
	RequirePasswordChange bool      `json:"requirePasswordChange"`
	Language              string    `json:"language"`
	Email                 string    `json:"email"`
	AvatarURL             string    `json:"avatarUrl"`
	Success               bool      `json:"success"`
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
