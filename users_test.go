package gorocket

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsersPresence(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"users":[{"_id":"rocket.cat","name":"Rocket.Cat","username":"rocket.cat","status":"online","utcOffset":0,"avatarETag":"5BB9B5ny5DkKdrwkq"},{"_id":"rocketchat.internal.admin.test","name":"RocketChat Internal Admin Test","username":"rocketchat.internal.admin.test","status":"online","utcOffset":-2,"avatarETag":"iEbEm4bTT327NJjXt"}],"full":true,"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersPresence("rocket")
	require.NoError(t, err)

	require.Equal(t, 2, len(resp.Users))
	require.Equal(t, "rocket.cat", resp.Users[0].ID)
	require.Equal(t, "Rocket.Cat", resp.Users[0].Name)
	require.Equal(t, "rocket.cat", resp.Users[0].Username)
	require.Equal(t, "online", resp.Users[0].Status)
	require.Equal(t, 0.0, resp.Users[0].UtcOffset)
	require.Equal(t, "5BB9B5ny5DkKdrwkq", resp.Users[0].AvatarETag)
	require.Equal(t, "rocketchat.internal.admin.test", resp.Users[1].ID)
	require.Equal(t, "RocketChat Internal Admin Test", resp.Users[1].Name)
	require.Equal(t, "rocketchat.internal.admin.test", resp.Users[1].Username)
	require.Equal(t, "online", resp.Users[1].Status)
	require.Equal(t, -2.0, resp.Users[1].UtcOffset)
	require.Equal(t, "iEbEm4bTT327NJjXt", resp.Users[1].AvatarETag)
	require.True(t, resp.Full)
	require.True(t, resp.Success)
}

func TestUsersCreate(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"user":{"_id":"BsNr28znDkG8aeo7W","createdAt":"2016-09-13T14:57:56.037Z","services":{"password":{"bcrypt":"$2a$i7BFS55uFYRf5TE4ErSUH8HymMNAbpMAvsOcl2C"}},"username":"uniqueusername","emails":[{"address":"email@user.tld","verified":false}],"type":"user","status":"offline","active":true,"roles":["user"],"_updatedAt":"2016-09-13T14:57:56.175Z","name":"name","settings":{}},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	payload := NewUser{
		Email:    "email@user.tld",
		Name:     "name",
		Password: "password",
		Username: "uniqueusername",
	}

	resp, err := client.UsersCreate(&payload)
	require.NoError(t, err)

	require.Equal(t, "BsNr28znDkG8aeo7W", resp.User.ID)
	require.Equal(t, "2016-09-13T14:57:56.037Z", resp.User.CreatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "2016-09-13T14:57:56.175Z", resp.User.UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "uniqueusername", resp.User.Username)
	require.Equal(t, "email@user.tld", resp.User.Emails[0].Address)
	require.Equal(t, false, resp.User.Emails[0].Verified)
	require.Equal(t, "user", resp.User.Type)
	require.Equal(t, "offline", resp.User.Status)
	require.Equal(t, true, resp.User.Active)
	require.Equal(t, "user", resp.User.Roles[0])
	require.Equal(t, "name", resp.User.Name)
	require.True(t, resp.Success)
}

func TestUsersDelete(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	req := UsersDelete{
		Username: "blabla",
	}

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersDelete(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestUsersCreateToken(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"data":{"authToken":"9HqLlyZOugoStsXCUfD_0YdwnNnunAJF8V47U3QHXSq","userId":"aobEdbYhXfu5hkeqG"},"success":true}`,
	}))
	defer server.Close()

	req := SimpleUserRequest{
		Username: "username",
	}

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersCreateToken(&req)
	require.NoError(t, err)

	require.Equal(t, "9HqLlyZOugoStsXCUfD_0YdwnNnunAJF8V47U3QHXSq", resp.Data.AuthToken)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Data.UserID)
	require.True(t, resp.Success)
}

func TestUsersDeactivateIdle(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"count": 1,"success":true}`,
	}))
	defer server.Close()

	req := DeactivateRequest{
		DaysIdle: "2",
	}

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersDeactivateIdle(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, 1, resp.Count)
}

func TestUsersDeleteOwnAccount(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersDeleteOwnAccount("password")
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestUsersForgotPassword(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersForgotPassword("rocket@cat.com")

	require.NoError(t, err)
	require.True(t, resp.Success)
}

func TestUsersGeneratePersonalAccessToken(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"token":"2jdk99wuSjXPO201XlAks9sjDjAhSJmskAKW301mSuj9Sk","success":true}`,
	}))
	defer server.Close()

	req := GetNewToken{
		Token: "test",
	}

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersGeneratePersonalAccessToken(&req)

	require.NoError(t, err)
	require.Equal(t, "2jdk99wuSjXPO201XlAks9sjDjAhSJmskAKW301mSuj9Sk", resp.Token)
	require.True(t, resp.Success)
}

func TestUsersGetStatus(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"message":"Latest status","connectionStatus":"online","status":"online","success":true}`,
	}))
	defer server.Close()

	req := SimpleUserRequest{
		Username: "rocket.cat",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersGetStatus(&req)

	require.NoError(t, err)
	require.Equal(t, "Latest status", resp.Message)
	require.Equal(t, "online", resp.ConnectionStatus)
	require.Equal(t, "online", resp.Status)
	require.True(t, resp.Success)
}

func TestUsersInfo(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"user":{"_id":"5fRTXMt7DMJbpPJfh","createdAt":"2023-07-10T16:44:58.548Z","services":{"password":true,"email2fa":{"enabled":true,"changedAt":"2023-07-10T16:44:58.546Z"},"resume":{"loginTokens":[{"when":"2023-10-05T18:55:02.996Z","hashedToken":"..."},{"when":"2023-10-05T19:09:30.415Z","hashedToken":"....."},{"when":"2023-10-10T23:40:46.098Z","hashedToken":"...."}]}},"username":"test.john","emails":[{"address":"test.john@test.com","verified":true}],"type":"user","status":"offline","active":true,"roles":["user","admin"],"name":"Test John","requirePasswordChange":false,"lastLogin":"2023-10-10T23:40:46.093Z","statusConnection":"offline","utcOffset":1,"statusText":"","avatarETag":"GFoEi6wv3uAxnzDcD","nickname":"tesuser2","canViewAllInfo":true},"success":true}`,
	}))
	defer server.Close()

	req := SimpleUserRequest{
		Username: "test.john",
	}

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersInfo(&req)

	require.NoError(t, err)
	require.Equal(t, "5fRTXMt7DMJbpPJfh", resp.User.ID)
	require.Equal(t, "2023-07-10T16:44:58.548Z", resp.User.CreatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "test.john", resp.User.Username)
	require.Equal(t, "user", resp.User.Type)
	require.Equal(t, "offline", resp.User.Status)
	require.Equal(t, true, resp.User.Active)
	require.Equal(t, "Test John", resp.User.Name)
	require.Equal(t, 1.0, resp.User.UtcOffset)
	require.Equal(t, "GFoEi6wv3uAxnzDcD", resp.User.AvatarETag)
	require.True(t, resp.Success)
}

func TestUsersRegister(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"user":{"_id":"nSYqWzZ4GsKTX4dyK","type":"user","status":"offline","active":true,"name":"Example User","utcOffset":0,"username":"example"},"success":true}`,
	}))
	defer server.Close()

	req := UserRegisterRequest{
		Email:    "rocket@chat.com",
		Name:     "Example User",
		Pass:     "password",
		Username: "example",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersRegister(&req)

	require.NoError(t, err)
	require.Equal(t, "nSYqWzZ4GsKTX4dyK", resp.User.ID)
	require.Equal(t, "user", resp.User.Type)
	require.Equal(t, "offline", resp.User.Status)
	require.Equal(t, true, resp.User.Active)
	require.Equal(t, "Example User", resp.User.Name)
	require.Equal(t, 0.0, resp.User.UtcOffset)
	require.Equal(t, "example", resp.User.Username)
	require.True(t, resp.Success)
}

func TestUsersSetStatus(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	req := SetStatus{
		Message: "I am online",
		Status:  "online",
	}

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersSetStatus(&req)

	require.NoError(t, err)
	require.True(t, resp.Success)
}

func TestUsersUpdate(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"user":{"_id":"BsNr28znDkG8aeo7W","createdAt":"2016-09-13T14:57:56.037Z","services":{"password":{"bcrypt":"$2a$10$5I5nUzqNEs8jKhi7BFS55uFYRf5TE4ErSUH8HymMNAbpMAvsOcl2C"}},"username":"uniqueusername","emails":[{"address":"newemail@user.tld","verified":false}],"type":"user","status":"offline","active":true,"roles":["user"],"_updatedAt":"2016-09-13T14:57:56.175Z","name":"new name","customFields":{"twitter":"userstwitter"}},"success":true}`,
	}))
	defer server.Close()

	req := UserUpdateRequest{
		UserId: "BsNr28znDkG8aeo7W",
		Data:   UserUpdateData{Name: "new name"},
	}

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UsersUpdate(&req)

	require.NoError(t, err)
	require.Equal(t, "BsNr28znDkG8aeo7W", resp.User.ID)
	require.Equal(t, "2016-09-13T14:57:56.037Z", resp.User.CreatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "2016-09-13T14:57:56.175Z", resp.User.UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "uniqueusername", resp.User.Username)
	require.Equal(t, "newemail@user.tld", resp.User.Emails[0].Address)
	require.Equal(t, false, resp.User.Emails[0].Verified)
	require.Equal(t, "user", resp.User.Roles[0])
	require.Equal(t, "new name", resp.User.Name)
	require.True(t, resp.Success)
}
