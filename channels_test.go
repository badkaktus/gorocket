package gorocket

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddAllToChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"channel":{"_id":"ByehQjC44FwMeiLbX","name":"channelname","t":"c","usernames":["example","rocket.cat"],"msgs":0,"u":{"_id":"aobEdbYhXfu5hkeqG","username":"example"},"ts":"2016-05-30T13:42:25.304Z"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := AddAllRequest{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.AddAllToChannel(&req)
	require.NoError(t, err)

	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Channel.ID)
	require.Equal(t, "channelname", resp.Channel.Name)
	require.Equal(t, "c", resp.Channel.T)
	require.Equal(t, "example", resp.Channel.Usernames[0])
	require.Equal(t, "rocket.cat", resp.Channel.Usernames[1])
	require.Equal(t, 0, resp.Channel.Msgs)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Channel.U.ID)
	require.Equal(t, "example", resp.Channel.U.Username)
	require.Equal(t, "2016-05-30T13:42:25.304Z", resp.Channel.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.True(t, resp.Success)
}

func TestArchiveChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SimpleChannelId{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.ArchiveChannel(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestCloseChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SimpleChannelId{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.CloseChannel(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestChannelCounters(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"joined":true,"members":78,"unreads":2,"unreadsFrom":"2018-02-23T17:15:51.907Z","msgs":304,"latest":"2018-02-23T17:17:03.110Z","userMentions":0,"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := ChannelCountersRequest{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.ChannelCounters(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.True(t, resp.Joined)
	require.Equal(t, 78, resp.Members)
	require.Equal(t, 2, resp.Unreads)
	require.Equal(t, "2018-02-23T17:15:51.907Z", resp.UnreadsFrom.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, 304, resp.Msgs)
	require.Equal(t, "2018-02-23T17:17:03.11Z", resp.Latest.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, 0, resp.UserMentions)
}

func TestCreateChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"channel":{"_id":"ByehQjC44FwMeiLbX","name":"channelname","t":"c","usernames":["example"],"msgs":0,"u":{"_id":"aobEdbYhXfu5hkeqG","username":"example"},"ts":"2016-05-30T13:42:25.304Z"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := CreateChannelRequest{
		Name:    "channelname",
		Members: []string{"example"},
	}
	resp, err := client.CreateChannel(&req)
	require.NoError(t, err)

	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Channel.ID)
	require.Equal(t, "channelname", resp.Channel.Name)
	require.Equal(t, "c", resp.Channel.T)
	require.Equal(t, "example", resp.Channel.Usernames[0])
	require.Equal(t, 0, resp.Channel.Msgs)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Channel.U.ID)
	require.Equal(t, "example", resp.Channel.U.Username)
	require.Equal(t, "2016-05-30T13:42:25.304Z", resp.Channel.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.True(t, resp.Success)
}

func TestDeleteChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SimpleChannelRequest{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.DeleteChannel(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestGetChannelInfo(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"channel":{"_id":"ByehQjC44FwMeiLbX","name":"testing","fname":"testing","t":"c","msgs":0,"usersCount":2,"u":{"_id":"HKKPmF8rZh45GMHWH","username":"marcos.defendi"},"customFields":{},"broadcast":false,"encrypted":false,"ts":"2020-05-21T13:14:07.070Z","ro":false,"default":false,"sysMes":true,"_updatedAt":"2020-05-21T13:14:07.096Z"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SimpleChannelRequest{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.ChannelInfo(&req)
	require.NoError(t, err)

	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Channel.ID)
	require.Equal(t, "testing", resp.Channel.Name)
	require.Equal(t, "testing", resp.Channel.Fname)
	require.Equal(t, "c", resp.Channel.T)
	require.Equal(t, 0, resp.Channel.Msgs)
	require.Equal(t, 2, resp.Channel.UsersCount)
	require.Equal(t, "HKKPmF8rZh45GMHWH", resp.Channel.U.ID)
	require.Equal(t, "marcos.defendi", resp.Channel.U.Username)
	require.Equal(t, false, resp.Channel.Broadcast)
	require.Equal(t, false, resp.Channel.Encrypted)
	require.Equal(t, "2020-05-21T13:14:07.07Z", resp.Channel.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, false, resp.Channel.Ro)
	require.Equal(t, false, resp.Channel.Default)
	require.Equal(t, true, resp.Channel.SysMes)
	require.Equal(t, "2020-05-21T13:14:07.096Z", resp.Channel.UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.True(t, resp.Success)
}

func TestChannelInvite(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"channel":{"_id":"ByehQjC44FwMeiLbX","ts":"2016-11-30T21:23:04.737Z","t":"c","name":"testing","usernames":["testing"],"msgs":1,"_updatedAt":"2016-12-09T12:50:51.575Z","lm":"2016-12-09T12:50:51.555Z"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := InviteChannelRequest{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.ChannelInvite(&req)
	require.NoError(t, err)

	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Channel.ID)
	require.Equal(t, "2016-11-30T21:23:04.737Z", resp.Channel.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "c", resp.Channel.T)
	require.Equal(t, "testing", resp.Channel.Name)
	require.Equal(t, "testing", resp.Channel.Usernames[0])
	require.Equal(t, 1, resp.Channel.Msgs)
	require.Equal(t, "2016-12-09T12:50:51.575Z", resp.Channel.UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "2016-12-09T12:50:51.555Z", resp.Channel.Lm.Format("2006-01-02T15:04:05.999Z"))
	require.True(t, resp.Success)
}

func TestChannelKick(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"channel":{"_id":"ByehQjC44FwMeiLbX","name":"invite-me","t":"c","usernames":["testing1"],"msgs":0,"u":{"_id":"aobEdbYhXfu5hkeqG","username":"testing1"},"ts":"2016-12-09T15:08:58.042Z","ro":false,"sysMes":true,"_updatedAt":"2016-12-09T15:22:40.656Z"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := InviteChannelRequest{
		RoomId: "ByehQjC44FwMeiLbX",
		UserId: "testing1",
	}
	resp, err := client.ChannelKick(&req)
	require.NoError(t, err)

	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Channel.ID)
	require.Equal(t, "invite-me", resp.Channel.Name)
	require.Equal(t, "c", resp.Channel.T)
	require.Equal(t, "testing1", resp.Channel.Usernames[0])
	require.Equal(t, 0, resp.Channel.Msgs)
	require.Equal(t, "2016-12-09T15:08:58.042Z", resp.Channel.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "2016-12-09T15:22:40.656Z", resp.Channel.UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.True(t, resp.Success)
}

func TestChannelList(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"channels":[{"_id":"ByehQjC44FwMeiLbX","name":"test-test","t":"c","usernames":["testing1"],"msgs":0,"u":{"_id":"aobEdbYhXfu5hkeqG","username":"testing1"},"ts":"2016-12-09T15:08:58.042Z","ro":false,"sysMes":true,"_updatedAt":"2016-12-09T15:22:40.656Z"},{"_id":"t7qapfhZjANMRAi5w","name":"testing","t":"c","usernames":["testing2"],"msgs":0,"u":{"_id":"y65tAmHs93aDChMWu","username":"testing2"},"ts":"2016-12-01T15:08:58.042Z","ro":false,"sysMes":true,"_updatedAt":"2016-12-09T15:22:40.656Z"}],"offset":0,"count":1,"total":1,"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	resp, err := client.Count(2).ChannelList()
	require.NoError(t, err)

	require.Equal(t, 2, len(resp.Channels))
	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Channels[0].ID)
	require.Equal(t, "test-test", resp.Channels[0].Name)
	require.Equal(t, "c", resp.Channels[0].T)
	require.Equal(t, "testing1", resp.Channels[0].Usernames[0])
	require.Equal(t, 0, resp.Channels[0].Msgs)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Channels[0].U.ID)
	require.Equal(t, "testing1", resp.Channels[0].U.Username)
	require.Equal(t, "2016-12-09T15:08:58.042Z", resp.Channels[0].Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, false, resp.Channels[0].Ro)
	require.Equal(t, true, resp.Channels[0].SysMes)
	require.Equal(t, "2016-12-09T15:22:40.656Z", resp.Channels[0].UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "t7qapfhZjANMRAi5w", resp.Channels[1].ID)
	require.Equal(t, "testing", resp.Channels[1].Name)
	require.Equal(t, "c", resp.Channels[1].T)
	require.Equal(t, "testing2", resp.Channels[1].Usernames[0])
	require.Equal(t, 0, resp.Channels[1].Msgs)
	require.Equal(t, "y65tAmHs93aDChMWu", resp.Channels[1].U.ID)
	require.Equal(t, "testing2", resp.Channels[1].U.Username)
	require.Equal(t, "2016-12-01T15:08:58.042Z", resp.Channels[1].Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, false, resp.Channels[1].Ro)
	require.Equal(t, true, resp.Channels[1].SysMes)
	require.Equal(t, "2016-12-09T15:22:40.656Z", resp.Channels[1].UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, 0, resp.Offset)
	require.Equal(t, 1, resp.Count)
	require.Equal(t, 1, resp.Total)
	require.True(t, resp.Success)
}

func TestChannelMembers(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"members":[{"_id":"Loz7qh9ChSqHMPymx","username":"customField_apiuser.test.1529436896005","name":"customField_apiuser.test.1529436896005","status":"offline"},{"_id":"Zc3Y3cRW7ZtS7Y8Hk","username":"customField_apiuser.test.1529436997563","name":"customField_apiuser.test.1529436997563","status":"offline"}],"count":2,"offset":0,"total":35,"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SimpleChannelRequest{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.Offset(10).ChannelMembers(&req)
	require.NoError(t, err)

	require.Equal(t, 2, len(resp.Members))
	require.Equal(t, "Loz7qh9ChSqHMPymx", resp.Members[0].ID)
	require.Equal(t, "customField_apiuser.test.1529436896005", resp.Members[0].Username)
	require.Equal(t, "customField_apiuser.test.1529436896005", resp.Members[0].Name)
	require.Equal(t, "offline", resp.Members[0].Status)
	require.Equal(t, "Zc3Y3cRW7ZtS7Y8Hk", resp.Members[1].ID)
	require.Equal(t, "customField_apiuser.test.1529436997563", resp.Members[1].Username)
	require.Equal(t, "customField_apiuser.test.1529436997563", resp.Members[1].Name)
	require.Equal(t, "offline", resp.Members[1].Status)
	require.Equal(t, 2, resp.Count)
	require.Equal(t, 0, resp.Offset)
	require.Equal(t, 35, resp.Total)
	require.True(t, resp.Success)
}

func TestOpenChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SimpleChannelId{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.OpenChannel(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestRenameChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"channel":{"_id":"ByehQjC44FwMeiLbX","name":"new-name","t":"c","usernames":["testing1"],"msgs":4,"u":{"_id":"aobEdbYhXfu5hkeqG","username":"testing1"},"ts":"2016-12-09T15:08:58.042Z","ro":false,"sysMes":true,"_updatedAt":"2016-12-09T15:57:44.686Z"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := RenameChannelRequest{
		RoomId:  "ByehQjC44FwMeiLbX",
		NewName: "new-name",
	}
	resp, err := client.RenameChannel(&req)
	require.NoError(t, err)

	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Channel.ID)
	require.Equal(t, "new-name", resp.Channel.Name)
	require.Equal(t, "c", resp.Channel.T)
	require.Equal(t, "testing1", resp.Channel.Usernames[0])
	require.Equal(t, 4, resp.Channel.Msgs)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Channel.U.ID)
	require.Equal(t, "testing1", resp.Channel.U.Username)
	require.Equal(t, "2016-12-09T15:08:58.042Z", resp.Channel.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, false, resp.Channel.Ro)
	require.Equal(t, true, resp.Channel.SysMes)
	require.Equal(t, "2016-12-09T15:57:44.686Z", resp.Channel.UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.True(t, resp.Success)
}

func TestSetAnnouncementChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"announcement":"Test out everything","success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SetAnnouncementRequest{
		RoomId:       "ByehQjC44FwMeiLbX",
		Announcement: "Test out everything",
	}
	resp, err := client.SetAnnouncementChannel(&req)
	require.NoError(t, err)

	require.Equal(t, "Test out everything", resp.Announcement)
	require.True(t, resp.Success)
}

func TestSetDescriptionChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"description":"Test out everything","success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SetDescriptionRequest{
		RoomId:      "ByehQjC44FwMeiLbX",
		Description: "Test out everything",
	}
	resp, err := client.SetDescriptionChannel(&req)
	require.NoError(t, err)

	require.Equal(t, "Test out everything", resp.Description)
	require.True(t, resp.Success)
}

func TestSetTopicChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"topic":"Test out everything","success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SetTopicRequest{
		RoomId: "ByehQjC44FwMeiLbX",
		Topic:  "Test out everything",
	}
	resp, err := client.SetTopicChannel(&req)
	require.NoError(t, err)

	require.Equal(t, "Test out everything", resp.Topic)
	require.True(t, resp.Success)
}

func TestUnarchiveChannel(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SimpleChannelId{
		RoomId: "ByehQjC44FwMeiLbX",
	}
	resp, err := client.UnarchiveChannel(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}
