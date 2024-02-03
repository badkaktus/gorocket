package gorocket

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"status":"success","data":{"authToken":"9HqLlyZOugoStsXCUfD_0YdwnNnunAJF8V47U3QHXSq","userId":"aobEdbYhXfu5hkeqG","me":{"_id":"aYjNnig8BEAWeQzMh","name":"Rocket Cat","emails":[{"address":"rocket.cat@rocket.chat","verified":false}],"status":"offline","statusConnection":"offline","username":"rocket.cat","utcOffset":-3,"active":true,"roles":["admin"],"settings":{"preferences":{}},"avatarUrl":"http://localhost:3000/avatar/test"}}}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	payload := LoginPayload{
		User:     "username",
		Password: "password",
	}
	resp, err := client.Login(&payload)
	require.NoError(t, err)

	require.Equal(t, "success", resp.Status)
	require.Equal(t, "9HqLlyZOugoStsXCUfD_0YdwnNnunAJF8V47U3QHXSq", resp.Data.AuthToken)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Data.UserID)
	require.Equal(t, "aYjNnig8BEAWeQzMh", resp.Data.Me.ID)
	require.Equal(t, "Rocket Cat", resp.Data.Me.Name)
	require.Equal(t, "rocket.cat@rocket.chat", resp.Data.Me.Emails[0].Address)
	require.Equal(t, false, resp.Data.Me.Emails[0].Verified)
	require.Equal(t, "offline", resp.Data.Me.Status)
	require.Equal(t, "offline", resp.Data.Me.StatusConnection)
	require.Equal(t, "rocket.cat", resp.Data.Me.Username)
	require.Equal(t, -3.0, resp.Data.Me.UtcOffset)
	require.Equal(t, true, resp.Data.Me.Active)
	require.Equal(t, "admin", resp.Data.Me.Roles[0])
	require.IsType(t, Settings{}, resp.Data.Me.Settings)
	require.IsType(t, Preferences{}, resp.Data.Me.Settings.Preferences)
	require.Equal(t, "http://localhost:3000/avatar/test", resp.Data.Me.AvatarURL)
}

func TestMe(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"_id":"aobEdbYhXfu5hkeqG","name":"Example User","emails":[{"address":"example@example.com","verified":true}],"status":"offline","statusConnection":"offline","username":"example","utcOffset":0,"active":true,"roles":["user","admin"],"settings":{"preferences":{"enableAutoAway":false,"idleTimeoutLimit":300,"desktopNotificationDuration":0,"audioNotifications":"mentions","desktopNotifications":"mentions","mobileNotifications":"mentions","unreadAlert":true,"useEmojis":true,"convertAsciiEmoji":true,"autoImageLoad":true,"saveMobileBandwidth":true,"collapseMediaByDefault":false,"hideUsernames":false,"hideRoles":false,"hideFlexTab":false,"hideAvatars":false,"roomsListExhibitionMode":"category","sidebarViewMode":"medium","sidebarHideAvatar":false,"sidebarShowUnread":false,"sidebarShowFavorites":true,"sendOnEnter":"normal","messageViewMode":0,"emailNotificationMode":"all","roomCounterSidebar":false,"newRoomNotification":"door","newMessageNotification":"chime","muteFocusedConversations":true,"notificationsSoundVolume":100}},"customFields":{"twitter":"@userstwi"},"avatarUrl":"http://localhost:3000/avatar/test","success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.Me()
	require.NoError(t, err)

	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.ID)
	require.Equal(t, "Example User", resp.Name)
	require.Equal(t, "example@example.com", resp.Emails[0].Address)
	require.True(t, resp.Emails[0].Verified)
	require.Equal(t, "offline", resp.Status)
	require.Equal(t, "offline", resp.StatusConnection)
	require.Equal(t, "example", resp.Username)
	require.Equal(t, 0.0, resp.UtcOffset)
	require.True(t, resp.Active)
	require.Equal(t, "user", resp.Roles[0])
	require.Equal(t, "admin", resp.Roles[1])
	require.IsType(t, Settings{}, resp.Settings)
	require.IsType(t, Preferences{}, resp.Settings.Preferences)
	require.Equal(t, "http://localhost:3000/avatar/test", resp.AvatarURL)
	require.True(t, resp.Success)
}

func TestLogout(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"status":"success","data":{"message":"You've been logged out!"}}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.Logout()
	require.NoError(t, err)

	require.Equal(t, "success", resp.Status)
	require.Equal(t, "You've been logged out!", resp.Data.Message)
}
