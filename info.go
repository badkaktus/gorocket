package gorocket

import (
	"fmt"
	"net/http"
	"time"
)

type RespInfo struct {
	Version string `json:"version"`
	Info    struct {
		Version               string `json:"version"`
		Build                 Build  `json:"build"`
		Commit                Commit `json:"commit"`
		MarketplaceAPIVersion string `json:"marketplaceApiVersion"`
	} `json:"info"`
	Success bool `json:"success"`
}

type Build struct {
	Date        time.Time `json:"date"`
	NodeVersion string    `json:"nodeVersion"`
	Arch        string    `json:"arch"`
	Platform    string    `json:"platform"`
	OsRelease   string    `json:"osRelease"`
	TotalMemory int64     `json:"totalMemory"`
	FreeMemory  int       `json:"freeMemory"`
	Cpus        int       `json:"cpus"`
}

type Commit struct {
	Hash    string `json:"hash"`
	Date    string `json:"date"`
	Author  string `json:"author"`
	Subject string `json:"subject"`
	Tag     string `json:"tag"`
	Branch  string `json:"branch"`
}

type RespDirectory struct {
	Result  []DirectoryResult `json:"result"`
	Count   int               `json:"count"`
	Offset  int               `json:"offset"`
	Total   int               `json:"total"`
	Success bool              `json:"success"`
}

type DirectoryResult struct {
	ID         string    `json:"_id"`
	CreatedAt  time.Time `json:"createdAt"`
	Emails     []Email   `json:"emails"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	AvatarETag string    `json:"avatarETag"`
}

type Result struct {
	ID          string      `json:"_id"`
	Ts          time.Time   `json:"ts"`
	T           string      `json:"t"`
	Name        string      `json:"name"`
	UsersCount  int         `json:"usersCount"`
	Default     bool        `json:"default"`
	LastMessage LastMessage `json:"lastMessage"`
}

type LastMessage struct {
	ID          string        `json:"_id"`
	Alias       string        `json:"alias"`
	Msg         string        `json:"msg"`
	Attachments []interface{} `json:"attachments"`
	ParseUrls   bool          `json:"parseUrls"`
	Groupable   bool          `json:"groupable"`
	Ts          time.Time     `json:"ts"`
	U           U             `json:"u"`
	Rid         string        `json:"rid"`
	UpdatedAt   time.Time     `json:"_updatedAt"`
	Mentions    []interface{} `json:"mentions"`
	Channels    []interface{} `json:"channels"`
}

type U struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Name     string `json:"name,omitempty"`
}

type RespSpotlight struct {
	Users   []UsersInfo `json:"users"`
	Rooms   []RoomsInfo `json:"rooms"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
}

type UsersInfo struct {
	ID         string `json:"_id"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	StatusText string `json:"statusText"`
	AvatarETag string `json:"avatarETag,omitempty"`
}

type RoomsInfo struct {
	ID          string      `json:"_id"`
	Name        string      `json:"name"`
	T           string      `json:"t"`
	LastMessage LastMessage `json:"lastMessage"`
}

type RespStatistics struct {
	ID                        string       `json:"_id"`
	Wizard                    Wizard       `json:"wizard"`
	UniqueID                  string       `json:"uniqueId"`
	InstalledAt               time.Time    `json:"installedAt"`
	Version                   string       `json:"version"`
	TotalUsers                int          `json:"totalUsers"`
	ActiveUsers               int          `json:"activeUsers"`
	ActiveGuests              int          `json:"activeGuests"`
	NonActiveUsers            int          `json:"nonActiveUsers"`
	AppUsers                  int          `json:"appUsers"`
	OnlineUsers               int          `json:"onlineUsers"`
	AwayUsers                 int          `json:"awayUsers"`
	TotalConnectedUsers       int          `json:"totalConnectedUsers"`
	OfflineUsers              int          `json:"offlineUsers"`
	TotalRooms                int          `json:"totalRooms"`
	TotalChannels             int          `json:"totalChannels"`
	TotalPrivateGroups        int          `json:"totalPrivateGroups"`
	TotalDirect               int          `json:"totalDirect"`
	TotalLivechat             int          `json:"totalLivechat"`
	TotalDiscussions          int          `json:"totalDiscussions"`
	TotalThreads              int          `json:"totalThreads"`
	TotalLivechatVisitors     int          `json:"totalLivechatVisitors"`
	TotalLivechatAgents       int          `json:"totalLivechatAgents"`
	LivechatEnabled           bool         `json:"livechatEnabled"`
	TotalChannelMessages      int          `json:"totalChannelMessages"`
	TotalPrivateGroupMessages int          `json:"totalPrivateGroupMessages"`
	TotalDirectMessages       int          `json:"totalDirectMessages"`
	TotalLivechatMessages     int          `json:"totalLivechatMessages"`
	TotalMessages             int          `json:"totalMessages"`
	FederatedServers          int          `json:"federatedServers"`
	FederatedUsers            int          `json:"federatedUsers"`
	LastLogin                 time.Time    `json:"lastLogin"`
	LastMessageSentAt         time.Time    `json:"lastMessageSentAt"`
	LastSeenSubscription      time.Time    `json:"lastSeenSubscription"`
	Os                        Os           `json:"os"`
	Process                   Process      `json:"process"`
	Deploy                    Deploy       `json:"deploy"`
	EnterpriseReady           bool         `json:"enterpriseReady"`
	UploadsTotal              int          `json:"uploadsTotal"`
	UploadsTotalSize          int          `json:"uploadsTotalSize"`
	Migration                 Migration    `json:"migration"`
	InstanceCount             int          `json:"instanceCount"`
	OplogEnabled              bool         `json:"oplogEnabled"`
	MongoVersion              string       `json:"mongoVersion"`
	MongoStorageEngine        string       `json:"mongoStorageEngine"`
	UniqueUsersOfYesterday    Stats        `json:"uniqueUsersOfYesterday"`
	UniqueUsersOfLastMonth    Stats        `json:"uniqueUsersOfLastMonth"`
	UniqueDevicesOfYesterday  Stats        `json:"uniqueDevicesOfYesterday"`
	UniqueDevicesOfLastMonth  Stats        `json:"uniqueDevicesOfLastMonth"`
	UniqueOSOfYesterday       Stats        `json:"uniqueOSOfYesterday"`
	UniqueOSOfLastMonth       Stats        `json:"uniqueOSOfLastMonth"`
	Apps                      Apps         `json:"apps"`
	Integrations              Integrations `json:"integrations"`
	PushQueue                 int          `json:"pushQueue"`
	CreatedAt                 time.Time    `json:"createdAt"`
	UpdatedAt                 time.Time    `json:"_updatedAt"`
	Success                   bool         `json:"success"`
}

type Wizard struct {
	OrganizationType string `json:"organizationType"`
	Industry         string `json:"industry"`
	Size             string `json:"size"`
	Country          string `json:"country"`
	Language         string `json:"language"`
	ServerType       string `json:"serverType"`
	RegisterServer   bool   `json:"registerServer"`
}

type Os struct {
	Type     string    `json:"type"`
	Platform string    `json:"platform"`
	Arch     string    `json:"arch"`
	Release  string    `json:"release"`
	Uptime   int       `json:"uptime"`
	Loadavg  []float64 `json:"loadavg"`
	Totalmem int64     `json:"totalmem"`
	Freemem  int       `json:"freemem"`
	Cpus     []Cpus    `json:"cpus"`
}

type Cpus struct {
	Model string `json:"model"`
	Speed int    `json:"speed"`
	Times Times  `json:"times"`
}

type Times struct {
	User int `json:"user"`
	Nice int `json:"nice"`
	Sys  int `json:"sys"`
	Idle int `json:"idle"`
	Irq  int `json:"irq"`
}

type Process struct {
	NodeVersion string  `json:"nodeVersion"`
	Pid         int     `json:"pid"`
	Uptime      float64 `json:"uptime"`
}

type Apps struct {
	EngineVersion  string `json:"engineVersion"`
	Enabled        bool   `json:"enabled"`
	TotalInstalled int    `json:"totalInstalled"`
	TotalActive    int    `json:"totalActive"`
}

type Integrations struct {
	TotalIntegrations      int `json:"totalIntegrations"`
	TotalIncoming          int `json:"totalIncoming"`
	TotalIncomingActive    int `json:"totalIncomingActive"`
	TotalOutgoing          int `json:"totalOutgoing"`
	TotalOutgoingActive    int `json:"totalOutgoingActive"`
	TotalWithScriptEnabled int `json:"totalWithScriptEnabled"`
}

type Stats struct {
	Year  int           `json:"year"`
	Month int           `json:"month"`
	Day   int           `json:"day"`
	Data  []interface{} `json:"data"`
}

type Migration struct {
	ID      string `json:"_id"`
	Locked  bool   `json:"locked"`
	Version int    `json:"version"`
}

type Deploy struct {
	Method   string `json:"method"`
	Platform string `json:"platform"`
}

type RespStatisticsList struct {
	Statistics []struct {
		ID                        string       `json:"_id"`
		Wizard                    Wizard       `json:"wizard"`
		UniqueID                  string       `json:"uniqueId"`
		InstalledAt               time.Time    `json:"installedAt"`
		Version                   string       `json:"version"`
		TotalUsers                int          `json:"totalUsers"`
		ActiveUsers               int          `json:"activeUsers"`
		ActiveGuests              int          `json:"activeGuests"`
		NonActiveUsers            int          `json:"nonActiveUsers"`
		AppUsers                  int          `json:"appUsers"`
		OnlineUsers               int          `json:"onlineUsers"`
		AwayUsers                 int          `json:"awayUsers"`
		TotalConnectedUsers       int          `json:"totalConnectedUsers"`
		OfflineUsers              int          `json:"offlineUsers"`
		TotalRooms                int          `json:"totalRooms"`
		TotalChannels             int          `json:"totalChannels"`
		TotalPrivateGroups        int          `json:"totalPrivateGroups"`
		TotalDirect               int          `json:"totalDirect"`
		TotalLivechat             int          `json:"totalLivechat"`
		TotalDiscussions          int          `json:"totalDiscussions"`
		TotalThreads              int          `json:"totalThreads"`
		TotalLivechatVisitors     int          `json:"totalLivechatVisitors"`
		TotalLivechatAgents       int          `json:"totalLivechatAgents"`
		TotalChannelMessages      int          `json:"totalChannelMessages"`
		TotalPrivateGroupMessages int          `json:"totalPrivateGroupMessages"`
		TotalDirectMessages       int          `json:"totalDirectMessages"`
		TotalLivechatMessages     int          `json:"totalLivechatMessages"`
		TotalMessages             int          `json:"totalMessages"`
		FederatedServers          int          `json:"federatedServers"`
		FederatedUsers            int          `json:"federatedUsers"`
		Os                        Os           `json:"os"`
		Process                   Process      `json:"process"`
		Deploy                    Deploy       `json:"deploy"`
		EnterpriseReady           bool         `json:"enterpriseReady"`
		UploadsTotal              int          `json:"uploadsTotal"`
		UploadsTotalSize          int          `json:"uploadsTotalSize"`
		Migration                 Migration    `json:"migration"`
		InstanceCount             int          `json:"instanceCount"`
		OplogEnabled              bool         `json:"oplogEnabled"`
		MongoVersion              string       `json:"mongoVersion"`
		MongoStorageEngine        string       `json:"mongoStorageEngine"`
		UniqueUsersOfYesterday    Stats        `json:"uniqueUsersOfYesterday"`
		UniqueUsersOfLastMonth    Stats        `json:"uniqueUsersOfLastMonth"`
		UniqueDevicesOfYesterday  Stats        `json:"uniqueDevicesOfYesterday"`
		UniqueDevicesOfLastMonth  Stats        `json:"uniqueDevicesOfLastMonth"`
		UniqueOSOfYesterday       Stats        `json:"uniqueOSOfYesterday"`
		UniqueOSOfLastMonth       Stats        `json:"uniqueOSOfLastMonth"`
		Apps                      Apps         `json:"apps"`
		Integrations              Integrations `json:"integrations"`
		PushQueue                 int          `json:"pushQueue"`
		CreatedAt                 time.Time    `json:"createdAt"`
		UpdatedAt                 time.Time    `json:"_updatedAt"`
		LivechatEnabled           bool         `json:"livechatEnabled,omitempty"`
		LastLogin                 time.Time    `json:"lastLogin,omitempty"`
		LastMessageSentAt         time.Time    `json:"lastMessageSentAt,omitempty"`
		LastSeenSubscription      time.Time    `json:"lastSeenSubscription,omitempty"`
	} `json:"statistics"`
	Count   int  `json:"count"`
	Offset  int  `json:"offset"`
	Total   int  `json:"total"`
	Success bool `json:"success"`
}

// Info returns information about the server
func (c *Client) Info() (*RespInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/info", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	res := RespInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Directory returns a list of channels
func (c *Client) Directory() (*RespDirectory, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/directory", c.baseURL, c.apiVersion), nil)
	if err != nil {
		return nil, err
	}

	res := RespDirectory{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Spotlight returns a list of users and rooms that match the provided query
func (c *Client) Spotlight(query string) (*RespSpotlight, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/spotlight?query=%s", c.baseURL, c.apiVersion, query), nil)

	if err != nil {
		return nil, err
	}

	res := RespSpotlight{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Statistics returns statistics about the server
func (c *Client) Statistics() (*RespStatistics, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/statistics", c.baseURL, c.apiVersion), nil)
	if err != nil {
		return nil, err
	}

	res := RespStatistics{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// StatisticsList returns a list of statistics
func (c *Client) StatisticsList() (*RespStatisticsList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/statistics.list", c.baseURL, c.apiVersion), nil)
	if err != nil {
		return nil, err
	}

	res := RespStatisticsList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
