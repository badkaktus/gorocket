package gorocket

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArchiveGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	req := SimpleGroupId{
		RoomId: "GENERAL",
	}
	resp, err := client.ArchiveGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestCloseGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	req := SimpleGroupId{
		RoomId: "GENERAL",
	}
	resp, err := client.CloseGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestGroupCounters(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"joined":true,"members":1,"unreads":1,"unreadsFrom":"2018-02-18T21:51:20.091Z","msgs":1,"latest":"2018-02-23T17:20:17.345Z","userMentions":0,"success":true}`,
	}))
	defer server.Close()

	req := GroupCountersRequest{
		RoomId: "GENERAL",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.GroupCounters(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.True(t, resp.Joined)
	require.Equal(t, 1, resp.Members)
	require.Equal(t, 1, resp.Unreads)
	require.Equal(t, "2018-02-18T21:51:20.091Z", resp.UnreadsFrom.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, 1, resp.Msgs)
	require.Equal(t, "2018-02-23T17:20:17.345Z", resp.Latest.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, 0, resp.UserMentions)
}

func TestCreateGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"group":{"_id":"NtR6RQ7NvzA9ejecX","name":"testing","t":"p","msgs":0,"u":{"_id":"aobEdbYhXfu5hkeqG","username":"tester"},"ts":"2016-12-09T16:53:06.761Z","ro":false,"sysMes":true,"_updatedAt":"2016-12-09T16:53:06.761Z"},"success":true}`,
	}))
	defer server.Close()

	req := CreateGroupRequest{
		Name:     "GENERAL",
		ReadOnly: false,
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.CreateGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "NtR6RQ7NvzA9ejecX", resp.Group.ID)
	require.Equal(t, "testing", resp.Group.Name)
	require.Equal(t, "p", resp.Group.T)
	require.Equal(t, 0, resp.Group.Msgs)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Group.U.ID)
	require.Equal(t, "tester", resp.Group.U.Username)
	require.Equal(t, "2016-12-09T16:53:06.761Z", resp.Group.Ts.Format("2006-01-02T15:04:05.000Z"))
}

func TestDeleteGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	req := SimpleGroupId{
		RoomId: "GENERAL",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.DeleteGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestGroupInfo(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"group":{"_id":"ByehQjC44FwMeiLbX","name":"testing","fname":"testing","t":"p","msgs":0,"usersCount":2,"u":{"_id":"HKKPmF8rZh45GMHWH","username":"marcos.defendi"},"customFields":{},"broadcast":false,"encrypted":false,"ts":"2020-05-21T13:16:24.749Z","ro":false,"default":false,"sysMes":true,"_updatedAt":"2020-05-21T13:16:24.772Z"},"success":true}`,
	}))
	defer server.Close()

	req := SimpleGroupRequest{
		RoomId: "GENERAL",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.GroupInfo(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Group.ID)
	require.Equal(t, "testing", resp.Group.Name)
	require.Equal(t, "testing", resp.Group.Fname)
	require.Equal(t, "p", resp.Group.T)
	require.Equal(t, 0, resp.Group.Msgs)
	require.Equal(t, 2, resp.Group.UsersCount)
	require.Equal(t, "HKKPmF8rZh45GMHWH", resp.Group.U.ID)
	require.Equal(t, "marcos.defendi", resp.Group.U.Username)
	require.False(t, resp.Group.Broadcast)
	require.False(t, resp.Group.Encrypted)
	require.Equal(t, "2020-05-21T13:16:24.749Z", resp.Group.Ts.Format("2006-01-02T15:04:05.000Z"))
	require.False(t, resp.Group.Ro)
	require.False(t, resp.Group.Default)
	require.True(t, resp.Group.SysMes)
	require.Equal(t, "2020-05-21T13:16:24.772Z", resp.Group.UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
}

func TestGroupInvite(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"group":{"_id":"ByehQjC44FwMeiLbX","ts":"2016-11-30T21:23:04.737Z","t":"p","name":"testing","username":"testing","u":{"_id":"aobEdbYhXfu5hkeqG","username":"testing1"},"msgs":1,"_updatedAt":"2016-12-09T12:50:51.575Z","lm":"2016-12-09T12:50:51.555Z"},"success":true}`,
	}))
	defer server.Close()

	req := InviteGroupRequest{
		RoomId: "GENERAL",
		UserId: "1234",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.GroupInvite(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Group.ID)
	require.Equal(t, "2016-11-30T21:23:04.737Z", resp.Group.Ts.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "p", resp.Group.T)
	require.Equal(t, "testing", resp.Group.Name)
	require.Equal(t, "testing", resp.Group.Username)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Group.U.ID)
	require.Equal(t, "testing1", resp.Group.U.Username)
	require.Equal(t, 1, resp.Group.Msgs)
	require.Equal(t, "2016-12-09T12:50:51.575Z", resp.Group.UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "2016-12-09T12:50:51.555Z", resp.Group.Lm.Format("2006-01-02T15:04:05.000Z"))
}

func TestGroupKick(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"group":{"_id":"ByehQjC44FwMeiLbX","ts":"2016-11-30T21:23:04.737Z","t":"p","name":"testing","username":"testing","u":{"_id":"aobEdbYhXfu5hkeqG","username":"testing1"},"msgs":1,"_updatedAt":"2016-12-09T12:50:51.575Z","lm":"2016-12-09T12:50:51.555Z"},"success":true}`,
	}))
	defer server.Close()

	req := InviteGroupRequest{
		RoomId: "GENERAL",
		UserId: "1234",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.GroupKick(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Group.ID)
	require.Equal(t, "2016-11-30T21:23:04.737Z", resp.Group.Ts.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "p", resp.Group.T)
	require.Equal(t, "testing", resp.Group.Name)
	require.Equal(t, "testing", resp.Group.Username)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Group.U.ID)
	require.Equal(t, "testing1", resp.Group.U.Username)
	require.Equal(t, 1, resp.Group.Msgs)
	require.Equal(t, "2016-12-09T12:50:51.575Z", resp.Group.UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "2016-12-09T12:50:51.555Z", resp.Group.Lm.Format("2006-01-02T15:04:05.000Z"))
}

func TestGroupList(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"groups":[{"_id":"ByehQjC44FwMeiLbX","name":"test-test","t":"p","msgs":0,"u":{"_id":"aobEdbYhXfu5hkeqG","username":"testing1"},"ts":"2016-12-09T15:08:58.042Z","ro":false,"sysMes":true,"_updatedAt":"2016-12-09T15:22:40.656Z"},{"_id":"t7qapfhZjANMRAi5w","name":"testing","t":"p","msgs":0,"u":{"_id":"y65tAmHs93aDChMWu","username":"testing2"},"ts":"2016-12-01T15:08:58.042Z","ro":false,"sysMes":true,"_updatedAt":"2016-12-09T15:22:40.656Z"}],"offset":0,"count":1,"total":1,"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.GroupList()
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, 0, resp.Offset)
	require.Equal(t, 1, resp.Count)
	require.Equal(t, 1, resp.Total)
	require.Equal(t, 2, len(resp.Groups))
	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Groups[0].ID)
	require.Equal(t, "test-test", resp.Groups[0].Name)
	require.Equal(t, "p", resp.Groups[0].T)
	require.Equal(t, 0, resp.Groups[0].Msgs)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Groups[0].U.ID)
	require.Equal(t, "testing1", resp.Groups[0].U.Username)
	require.Equal(t, "2016-12-09T15:08:58.042Z", resp.Groups[0].Ts.Format("2006-01-02T15:04:05.000Z"))
	require.False(t, resp.Groups[0].Ro)
	require.True(t, resp.Groups[0].SysMes)
	require.Equal(t, "2016-12-09T15:22:40.656Z", resp.Groups[0].UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "t7qapfhZjANMRAi5w", resp.Groups[1].ID)
	require.Equal(t, "testing", resp.Groups[1].Name)
	require.Equal(t, "p", resp.Groups[1].T)
	require.Equal(t, 0, resp.Groups[1].Msgs)
	require.Equal(t, "y65tAmHs93aDChMWu", resp.Groups[1].U.ID)
	require.Equal(t, "testing2", resp.Groups[1].U.Username)
	require.Equal(t, "2016-12-01T15:08:58.042Z", resp.Groups[1].Ts.Format("2006-01-02T15:04:05.000Z"))
	require.False(t, resp.Groups[1].Ro)
	require.True(t, resp.Groups[1].SysMes)
	require.Equal(t, "2016-12-09T15:22:40.656Z", resp.Groups[1].UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
}

func TestGroupMembers(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"members":[{"_id":"Q4GkX6RMepGDdQ7YJ","status":"online","name":"Marcos Defendi","utcOffset":-3,"username":"marcos.defendi"},{"_id":"rocket.cat","name":"Rocket.Cat","username":"rocket.cat","status":"online","utcOffset":0}],"count":2,"offset":0,"total":2,"success":true}`,
	}))
	defer server.Close()

	req := SimpleGroupRequest{
		RoomId: "GENERAL",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.GroupMembers(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, 2, resp.Count)
	require.Equal(t, 0, resp.Offset)
	require.Equal(t, 2, resp.Total)
	require.Equal(t, 2, len(resp.Members))
	require.Equal(t, "Q4GkX6RMepGDdQ7YJ", resp.Members[0].ID)
	require.Equal(t, "online", resp.Members[0].Status)
	require.Equal(t, "Marcos Defendi", resp.Members[0].Name)
	require.Equal(t, "marcos.defendi", resp.Members[0].Username)
	require.Equal(t, "rocket.cat", resp.Members[1].ID)
	require.Equal(t, "Rocket.Cat", resp.Members[1].Name)
	require.Equal(t, "rocket.cat", resp.Members[1].Username)
	require.Equal(t, "online", resp.Members[1].Status)
}

func TestGroupMessages(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"messages":[{"_id":"CeXwh5eBbdrtvnqG6","rid":"agh2Sucgb54RQ8dDo","msg":"s","ts":"2018-10-05T13:48:21.616Z","u":{"_id":"KPkEYwKKBKZnEEPpt","username":"marcos.defendi","name":"Marcos Defendi"},"_updatedAt":"2018-10-05T13:48:49.535Z","reactions":{":frowning2:":{"usernames":["marcos.defendi"]}},"mentions":[],"channels":[],"starred":{"_id":"KPkEYwKKBKZnEEPpt"}},{"_id":"MrAeupRiF9TvhMesK","t":"room_changed_privacy","rid":"agh2Sucgb54RQ8dDo","ts":"2018-10-05T00:11:16.998Z","msg":"Private Group","u":{"_id":"rocketchat.internal.admin.test","username":"rocketchat.internal.admin.test"},"groupable":false,"_updatedAt":"2018-10-05T00:11:16.998Z"}],"count":2,"offset":0,"total":2,"success":true}`,
	}))
	defer server.Close()

	req := SimpleGroupRequest{
		RoomId: "GENERAL",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.GroupMessages(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, 2, resp.Count)
	require.Equal(t, 0, resp.Offset)
	require.Equal(t, 2, resp.Total)
	require.Equal(t, 2, len(resp.Messages))
	require.Equal(t, "CeXwh5eBbdrtvnqG6", resp.Messages[0].ID)
	require.Equal(t, "agh2Sucgb54RQ8dDo", resp.Messages[0].Rid)
	require.Equal(t, "s", resp.Messages[0].Msg)
	require.Equal(t, "2018-10-05T13:48:21.616Z", resp.Messages[0].Ts.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "KPkEYwKKBKZnEEPpt", resp.Messages[0].U.ID)
	require.Equal(t, "marcos.defendi", resp.Messages[0].U.Username)
	require.Equal(t, "Marcos Defendi", resp.Messages[0].U.Name)
	require.Equal(t, "2018-10-05T13:48:49.535Z", resp.Messages[0].UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
	// require.Equal(t, 1, len(resp.Messages[0].Reactions))
	// require.Equal(t, 0, len(resp.Messages[0].Mentions))
	// require.Equal(t, 0, len(resp.Messages[0].Channels))
	require.Equal(t, "MrAeupRiF9TvhMesK", resp.Messages[1].ID)
	require.Equal(t, "agh2Sucgb54RQ8dDo", resp.Messages[1].Rid)
	require.Equal(t, "2018-10-05T00:11:16.998Z", resp.Messages[1].Ts.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "rocketchat.internal.admin.test", resp.Messages[1].U.ID)
	require.Equal(t, "rocketchat.internal.admin.test", resp.Messages[1].U.Username)
	require.Equal(t, "Private Group", resp.Messages[1].Msg)
	require.Equal(t, "2018-10-05T00:11:16.998Z", resp.Messages[1].UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
}

func TestOpenGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	req := SimpleGroupId{
		RoomId: "GENERAL",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.OpenGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestRenameGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"group":{"_id":"ByehQjC44FwMeiLbX","name":"new-name","t":"p","usernames":["testing1"],"msgs":4,"u":{"_id":"aobEdbYhXfu5hkeqG","username":"testing1"},"ts":"2016-12-09T15:08:58.042Z","ro":false,"sysMes":true,"_updatedAt":"2016-12-09T15:57:44.686Z"},"success":true}`,
	}))
	defer server.Close()

	req := RenameGroupRequest{
		RoomId:  "GENERAL",
		NewName: "new-name",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.RenameGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Group.ID)
	require.Equal(t, "new-name", resp.Group.Name)
	require.Equal(t, "p", resp.Group.T)
	require.Equal(t, 1, len(resp.Group.Usernames))
	require.Equal(t, "testing1", resp.Group.Usernames[0])
	require.Equal(t, 4, resp.Group.Msgs)
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Group.U.ID)
	require.Equal(t, "testing1", resp.Group.U.Username)
	require.Equal(t, "2016-12-09T15:08:58.042Z", resp.Group.Ts.Format("2006-01-02T15:04:05.000Z"))
	require.False(t, resp.Group.Ro)
	require.True(t, resp.Group.SysMes)
	require.Equal(t, "2016-12-09T15:57:44.686Z", resp.Group.UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
}

func TestAddLeaderGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	req := AddGroupPermissionRequest{
		RoomId: "GENERAL",
		UserId: "1234",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.AddLeaderGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestAddOwnerGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	req := AddGroupPermissionRequest{
		RoomId: "GENERAL",
		UserId: "1234",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.AddOwnerGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}

func TestSetAnnouncementGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"announcement": "Test out everything","success":true}`,
	}))
	defer server.Close()

	req := SetAnnouncementRequest{
		RoomId:       "GENERAL",
		Announcement: "Test out everything",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.SetAnnouncementGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "Test out everything", resp.Announcement)
}

func TestSetDescriptionGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"description": "Test out everything","success":true}`,
	}))
	defer server.Close()

	req := SetDescriptionRequest{
		RoomId:      "GENERAL",
		Description: "Test out everything",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.SetDescriptionGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "Test out everything", resp.Description)
}

func TestSetTopicGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"topic": "Test out everything","success":true}`,
	}))
	defer server.Close()

	req := SetTopicRequest{
		RoomId: "GENERAL",
		Topic:  "Test out everything",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.SetTopicGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
	require.Equal(t, "Test out everything", resp.Topic)
}

func TestUnarchiveGroup(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	req := SimpleGroupId{
		RoomId: "GENERAL",
	}
	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.UnarchiveGroup(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}
