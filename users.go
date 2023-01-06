package gorocket

import (
	"bytes"
	"crypto/sha256"
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
	ID         string  `json:"_id"`
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	Status     string  `json:"status"`
	UtcOffset  float64 `json:"utcOffset"`
	AvatarETag string  `json:"avatarETag"`
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
	Emails    []Email      `json:"emails"`
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

type SimpleUserRequest struct {
	UserId   string `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
}

type CreateTokenResponse struct {
	Data struct {
		UserID    string `json:"userId"`
		AuthToken string `json:"authToken"`
	} `json:"data"`
	Success bool `json:"success"`
}

type DeactivateRequest struct {
	DaysIdle string `json:"daysIdle"`
	Role     string `json:"role,omitempty"`
}

type DeactivateResponse struct {
	Count   int  `json:"count"`
	Success bool `json:"success"`
}

type GetNewToken struct {
	Token     string `json:"tokenName"`
	TwoFactor bool   `json:"bypassTwoFactor,omitempty"`
}

type NewTokenResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}

type GetStatusResponse struct {
	Message          string `json:"message"`
	ConnectionStatus string `json:"connectionStatus"`
	Status           string `json:"status"`
	Success          bool   `json:"success"`
}

type UsersInfoResponse struct {
	User    singleUserInfo `json:"user"`
	Success bool           `json:"success"`
}

type singleUserInfo struct {
	ID         string  `json:"_id"`
	Type       string  `json:"type"`
	Status     string  `json:"status"`
	Active     bool    `json:"active"`
	Name       string  `json:"name"`
	UtcOffset  float64 `json:"utcOffset"`
	Username   string  `json:"username"`
	AvatarETag string  `json:"avatarETag,omitempty"`
}

type UserRegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Pass      string `json:"pass"`
	Name      string `json:"name"`
	SecretURL string `json:"secret_url,omitempty"`
}

type SetStatus struct {
	Message string `json:"message"`
	Status  string `json:"status,omitempty"`
}

type UserUpdateRequest struct {
	UserId string         `json:"userId"`
	Data   UserUpdateData `json:"data,omitempty"`
}

type UserUpdateData struct {
	Email                 string   `json:"email,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Password              string   `json:"password,omitempty"`
	Username              string   `json:"username,omitempty"`
	Active                bool     `json:"active,omitempty"`
	Roles                 []string `json:"roles,omitempty"`
	RequirePasswordChange bool     `json:"requirePasswordChange,omitempty"`
	SendWelcomeEmail      bool     `json:"sendWelcomeEmail,omitempty"`
	Verified              bool     `json:"verified,omitempty"`
}

type UserUpdateResponse struct {
	User    userUpdateInfo `json:"user"`
	Success bool           `json:"success"`
}

type userUpdateInfo struct {
	ID        string    `json:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	Services  struct {
		Password struct {
			Bcrypt string `json:"bcrypt"`
		} `json:"password"`
	} `json:"services"`
	Username  string    `json:"username"`
	Emails    []Email   `json:"emails"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	Active    bool      `json:"active"`
	Roles     []string  `json:"roles"`
	UpdatedAt time.Time `json:"_updatedAt"`
	Name      string    `json:"name"`
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

// Create a user authentication token.
func (c *Client) UsersCreateToken(user *SimpleUserRequest) (*CreateTokenResponse, error) {
	opt, _ := json.Marshal(user)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.createToken", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := CreateTokenResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Deactivate Idle users. Requires edit-other-user-active-status permission.
func (c *Client) UsersDeactivateIdle(params *DeactivateRequest) (*DeactivateResponse, error) {
	opt, _ := json.Marshal(params)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.createToken", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := DeactivateResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Deletes your own user. Requires Allow Users to Delete Own Account enabled. Accessible from Administration -> Accounts.
func (c *Client) UsersDeleteOwnAccount(pass string) (*SimpleSuccessResponse, error) {

	param := struct {
		password string `json:"password"`
	}{}

	param.password = fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))

	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.deleteOwnAccount", c.baseURL, c.apiVersion),
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

// Send email to reset your password.
func (c *Client) UsersForgotPassword(email string) (*SimpleSuccessResponse, error) {
	param := struct {
		email string `json:"email"`
	}{
		email: email,
	}

	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.forgotPassword", c.baseURL, c.apiVersion),
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

// Generate Personal Access Token. Requires create-personal-access-tokens permission.
func (c *Client) UsersGeneratePersonalAccessToken(params *GetNewToken) (*NewTokenResponse, error) {
	opt, _ := json.Marshal(params)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.generatePersonalAccessToken", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := NewTokenResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Gets a user's Status if the query string userId or username is provided, otherwise it gets the callee's.
func (c *Client) UsersGetStatus(user *SimpleUserRequest) (*GetStatusResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/users.getStatus", c.baseURL, c.apiVersion), nil)

	if user.Username == "" && user.UserId == "" {
		return nil, fmt.Errorf("False parameters")
	}

	url := req.URL.Query()
	if user.Username != "" {
		url.Add("username", user.Username)
	}
	if user.UserId != "" {
		url.Add("userId", user.UserId)
	}
	req.URL.RawQuery = url.Encode()

	if err != nil {
		return nil, err
	}

	res := GetStatusResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Retrieves information about a user, the result is only limited to what the callee has access to view.
func (c *Client) UsersInfo(user *SimpleUserRequest) (*UsersInfoResponse, error) {
	opt, _ := json.Marshal(user)

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/users.info", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if user.Username == "" && user.UserId == "" {
		return nil, fmt.Errorf("False parameters")
	}

	url := req.URL.Query()
	if user.Username != "" {
		url.Add("username", user.Username)
	}
	if user.UserId != "" {
		url.Add("userId", user.UserId)
	}
	req.URL.RawQuery = url.Encode()

	if err != nil {
		return nil, err
	}

	res := UsersInfoResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Register a new user.
func (c *Client) UsersRegister(user *UserRegisterRequest) (*UsersInfoResponse, error) {
	opt, _ := json.Marshal(user)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.register", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := UsersInfoResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Sets a user Status when the status message and state is given.
func (c *Client) UsersSetStatus(status *SetStatus) (*SimpleSuccessResponse, error) {
	opt, _ := json.Marshal(status)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.setStatus", c.baseURL, c.apiVersion),
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

// Update an existing user.
func (c *Client) UsersUpdate(user *UserUpdateRequest) (*UserUpdateResponse, error) {
	opt, _ := json.Marshal(user)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/users.update", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := UserUpdateResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
