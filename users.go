package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UsersPresenceResponse struct {
	Users   []user `json:"users"`
	Full    bool   `json:"full"`
	Success bool   `json:"success"`
}

type user struct {
	ID         string `json:"_id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Status     string `json:"status"`
	UtcOffset  int    `json:"utcOffset"`
	AvatarETag string `json:"avatarETag"`
}

type NewUser struct {
	Email                 string   `json:"email"`
	Name                  string   `json:"name"`
	Password              string   `json:"password"`
	Username              string   `json:"username"`
	Active                bool     `json:"active,omitempty"`
	Roles                 []string `json:"roles,omitempty"`
	JoinDefaultChannels   bool     `json:"joinDefaultChannels,omitempty"`
	RequirePasswordChange bool     `json:"requirePasswordChange,omitempty"`
	SendWelcomeEmail      bool     `json:"sendWelcomeEmail,omitempty"`
	Verified              bool     `json:"verified,omitempty"`
	CustomFields          string   `json:"customFields,omitempty"`
}

type UserCreateResponse struct {
	User    userCreateInfo `json:"user"`
	Success bool           `json:"success"`
}

type userCreateInfo struct {
	ID        string       `json:"_id"`
	CreatedAt time.Time    `json:"createdAt"`
	Services  userServices `json:"services"`
	Username  string       `json:"username"`
	Emails    []email      `json:"emails"`
	Type      string       `json:"type"`
	Status    string       `json:"status"`
	Active    bool         `json:"active"`
	Roles     []string     `json:"roles"`
	UpdatedAt time.Time    `json:"_updatedAt"`
	Name      string       `json:"name"`
	Settings  struct {
	} `json:"settings"`
}

type userServices struct {
	Password struct {
		Bcrypt string `json:"bcrypt"`
	} `json:"password"`
}

type UsersDelete struct {
	Username string `json:"username"`
}

type SimpleSuccessResponse struct {
	Success bool `json:"success"`
}

// Gets all connected users presence
func (c *Client) UsersPresence(query string) (*UsersPresenceResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/users.presence?from=%s", c.baseURL, c.apiVersion, query), nil)

	if err != nil {
		return nil, err
	}

	res := UsersPresenceResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Create a new user. Requires create-user permission.
func (c *Client) UsersCreate(user *NewUser) (*UserCreateResponse, error) {
	opt, _ := json.Marshal(user)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.create", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := UserCreateResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Deletes an existing user. Requires delete-user permission.
func (c *Client) UsersDelete(user *UsersDelete) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(user)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.delete", c.baseURL, c.apiVersion),
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
