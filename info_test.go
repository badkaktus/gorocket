package gorocket

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"version":"6.5.2","info":{"version":"6.5.2","build":{"date":"2024-01-03T03:33:06.277Z","nodeVersion":"v14.21.4","arch":"x64","platform":"linux","osRelease":"5.15.0-1053-azure","totalMemory":16768364544,"freeMemory":812339200,"cpus":4},"marketplaceApiVersion":"1.41.0","commit":{"hash":"3ebc8e0868c859d2d8e636787645c29a89dea1e5","date":"Tue Jan 2 23:38:47 2024 -0300","author":"Diego Sampaio","subject":"chore: fix dep version","tag":"6.5.2","branch":"HEAD"}},"minimumClientVersions":{"desktop":"3.9.6","mobile":"4.39.0"},"supportedVersions":{"signed":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.Info()
	require.NoError(t, err)

	require.Equal(t, "6.5.2", resp.Version)
	require.Equal(t, "6.5.2", resp.Info.Version)
	require.Equal(t, "2024-01-03T03:33:06.277Z", resp.Info.Build.Date.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "v14.21.4", resp.Info.Build.NodeVersion)
	require.Equal(t, "x64", resp.Info.Build.Arch)
	require.Equal(t, "linux", resp.Info.Build.Platform)
	require.Equal(t, "5.15.0-1053-azure", resp.Info.Build.OsRelease)
	require.Equal(t, int64(16768364544), resp.Info.Build.TotalMemory)
	require.Equal(t, 812339200, resp.Info.Build.FreeMemory)
	require.Equal(t, 4, resp.Info.Build.Cpus)
	require.Equal(t, "1.41.0", resp.Info.MarketplaceAPIVersion)
	require.Equal(t, "3ebc8e0868c859d2d8e636787645c29a89dea1e5", resp.Info.Commit.Hash)
	require.Equal(t, "Tue Jan 2 23:38:47 2024 -0300", resp.Info.Commit.Date)
	require.Equal(t, "Diego Sampaio", resp.Info.Commit.Author)
	require.Equal(t, "chore: fix dep version", resp.Info.Commit.Subject)
	require.Equal(t, "6.5.2", resp.Info.Commit.Tag)
	require.Equal(t, "HEAD", resp.Info.Commit.Branch)
}

func TestDirectory(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"result":[{"_id":"jRca8kibJx8NkLJxt","createdAt":"2018-04-13T12:46:26.517Z","emails":[{"address":"user.test.1523623548558@rocket.chat","verified":false}],"name":"EditedRealNameuser.test.1523623548558","username":"editedusernameuser.test.1523623548558","avatarETag":"6YbLtc4v9b4conXon"}],"count":1,"offset":0,"total":1,"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.Directory()
	require.NoError(t, err)

	require.Equal(t, 1, resp.Count)
	require.Equal(t, 0, resp.Offset)
	require.Equal(t, 1, resp.Total)
	require.Equal(t, "jRca8kibJx8NkLJxt", resp.Result[0].ID)
	require.Equal(t, "2018-04-13T12:46:26.517Z", resp.Result[0].CreatedAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "user.test.1523623548558@rocket.chat", resp.Result[0].Emails[0].Address)
	require.Equal(t, false, resp.Result[0].Emails[0].Verified)
	require.Equal(t, "EditedRealNameuser.test.1523623548558", resp.Result[0].Name)
	require.Equal(t, "editedusernameuser.test.1523623548558", resp.Result[0].Username)
	require.Equal(t, "6YbLtc4v9b4conXon", resp.Result[0].AvatarETag)
}

func TestSpotlight(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"users":[{"_id":"rocket.cat","name":"Rocket.Cat","username":"rocket.cat","status":"online","avatarETag":"5BB9B5ny5DkKdrwkq"}],"rooms":[],"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.Spotlight("rocket")
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Users))
	require.Equal(t, "rocket.cat", resp.Users[0].ID)
	require.Equal(t, "Rocket.Cat", resp.Users[0].Name)
	require.Equal(t, "rocket.cat", resp.Users[0].Username)
	require.Equal(t, "online", resp.Users[0].Status)
	require.Equal(t, "5BB9B5ny5DkKdrwkq", resp.Users[0].AvatarETag)
	require.Equal(t, 0, len(resp.Rooms))
}

func TestStatistics(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"_id":"wufRdmSrjmSMhBdTN","uniqueId":"wD4EP3M7FeFzJZgk9","installedAt":"2018-02-18T19:40:45.369Z","version":"0.61.0-develop","totalUsers":88,"activeUsers":88,"nonActiveUsers":0,"onlineUsers":0,"awayUsers":1,"offlineUsers":87,"totalRooms":81,"totalChannels":41,"totalPrivateGroups":37,"totalDirect":3,"totlalLivechat":0,"totalMessages":2408,"totalChannelMessages":730,"totalPrivateGroupMessages":1869,"totalDirectMessages":25,"totalLivechatMessages":0,"lastLogin":"2018-02-24T12:44:45.045Z","lastMessageSentAt":"2018-02-23T18:14:03.490Z","lastSeenSubscription":"2018-02-23T17:58:54.779Z","os":{"type":"Linux","platform":"linux","arch":"x64","release":"4.13.0-32-generic","uptime":76242,"loadavg":[0.0576171875,0.04638671875,0.00439453125],"totalmem":5787901952,"freemem":1151168512,"cpus":[{"model":"Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz","speed":2405,"times":{"user":6437000,"nice":586500,"sys":1432200,"idle":750117500,"irq":0}},{"model":"Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz","speed":2405,"times":{"user":7319700,"nice":268800,"sys":1823600,"idle":747642700,"irq":0}},{"model":"Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz","speed":2405,"times":{"user":7484600,"nice":1003500,"sys":1446000,"idle":748873400,"irq":0}},{"model":"Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz","speed":2405,"times":{"user":8378200,"nice":548500,"sys":1443200,"idle":747053300,"irq":0}}]},"process":{"nodeVersion":"v8.9.4","pid":11736,"uptime":16265.506},"deploy":{"method":"tar","platform":"selfinstall"},"migration":{"_id":"control","version":106,"locked":false,"lockedAt":"2018-02-23T18:13:13.948Z","buildAt":"2018-02-18T17:22:51.212Z"},"instanceCount":1,"createdAt":"2018-02-24T13:13:00.236Z","_updatedAt":"2018-02-24T13:13:00.236Z","success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.Statistics()
	require.NoError(t, err)

	require.Equal(t, "wufRdmSrjmSMhBdTN", resp.ID)
	require.Equal(t, "wD4EP3M7FeFzJZgk9", resp.UniqueID)
	require.Equal(t, "2018-02-18T19:40:45.369Z", resp.InstalledAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "0.61.0-develop", resp.Version)
	require.Equal(t, 88, resp.TotalUsers)
	require.Equal(t, 88, resp.ActiveUsers)
	require.Equal(t, 0, resp.NonActiveUsers)
	require.Equal(t, 0, resp.OnlineUsers)
	require.Equal(t, 1, resp.AwayUsers)
	require.Equal(t, 87, resp.OfflineUsers)
	require.Equal(t, 81, resp.TotalRooms)
	require.Equal(t, 41, resp.TotalChannels)
	require.Equal(t, 37, resp.TotalPrivateGroups)
	require.Equal(t, 3, resp.TotalDirect)
	require.Equal(t, 0, resp.TotalLivechat)
	require.Equal(t, 2408, resp.TotalMessages)
	require.Equal(t, 730, resp.TotalChannelMessages)
	require.Equal(t, 1869, resp.TotalPrivateGroupMessages)
	require.Equal(t, 25, resp.TotalDirectMessages)
	require.Equal(t, 0, resp.TotalLivechatMessages)
	require.Equal(t, "2018-02-24T12:44:45.045Z", resp.LastLogin.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "2018-02-23T18:14:03.490Z", resp.LastMessageSentAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "2018-02-23T17:58:54.779Z", resp.LastSeenSubscription.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "Linux", resp.Os.Type)
	require.Equal(t, "linux", resp.Os.Platform)
	require.Equal(t, "x64", resp.Os.Arch)
	require.Equal(t, "4.13.0-32-generic", resp.Os.Release)
	require.Equal(t, 76242, resp.Os.Uptime)
	require.Equal(t, int64(5787901952), resp.Os.Totalmem)
	require.Equal(t, 1151168512, resp.Os.Freemem)
	require.Equal(t, 4, len(resp.Os.Cpus))
	require.Equal(t, "Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz", resp.Os.Cpus[0].Model)
	require.Equal(t, 2405, resp.Os.Cpus[0].Speed)
	require.Equal(t, 6437000, resp.Os.Cpus[0].Times.User)
	require.Equal(t, 586500, resp.Os.Cpus[0].Times.Nice)
	require.Equal(t, 1432200, resp.Os.Cpus[0].Times.Sys)
	require.Equal(t, 750117500, resp.Os.Cpus[0].Times.Idle)
	require.Equal(t, 0, resp.Os.Cpus[0].Times.Irq)
	require.Equal(t, "Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz", resp.Os.Cpus[1].Model)
	require.Equal(t, 2405, resp.Os.Cpus[1].Speed)
	require.Equal(t, 7319700, resp.Os.Cpus[1].Times.User)
	require.Equal(t, 268800, resp.Os.Cpus[1].Times.Nice)
	require.Equal(t, 1823600, resp.Os.Cpus[1].Times.Sys)
	require.Equal(t, 747642700, resp.Os.Cpus[1].Times.Idle)
	require.Equal(t, 0, resp.Os.Cpus[1].Times.Irq)
	require.Equal(t, "Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz", resp.Os.Cpus[2].Model)
	require.Equal(t, 2405, resp.Os.Cpus[2].Speed)
	require.Equal(t, 7484600, resp.Os.Cpus[2].Times.User)
	require.Equal(t, 1003500, resp.Os.Cpus[2].Times.Nice)
	require.Equal(t, 1446000, resp.Os.Cpus[2].Times.Sys)
	require.Equal(t, 748873400, resp.Os.Cpus[2].Times.Idle)
	require.Equal(t, 0, resp.Os.Cpus[2].Times.Irq)
	require.Equal(t, "Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz", resp.Os.Cpus[3].Model)
	require.Equal(t, 2405, resp.Os.Cpus[3].Speed)
	require.Equal(t, 8378200, resp.Os.Cpus[3].Times.User)
	require.Equal(t, 548500, resp.Os.Cpus[3].Times.Nice)
	require.Equal(t, 1443200, resp.Os.Cpus[3].Times.Sys)
	require.Equal(t, 747053300, resp.Os.Cpus[3].Times.Idle)
	require.Equal(t, 0, resp.Os.Cpus[3].Times.Irq)
	require.Equal(t, "v8.9.4", resp.Process.NodeVersion)
	require.Equal(t, 11736, resp.Process.Pid)
	require.Equal(t, 16265.506, resp.Process.Uptime)
	require.Equal(t, "tar", resp.Deploy.Method)
	require.Equal(t, "selfinstall", resp.Deploy.Platform)
	require.Equal(t, "control", resp.Migration.ID)
	require.Equal(t, 106, resp.Migration.Version)
	require.Equal(t, false, resp.Migration.Locked)
	require.Equal(t, 1, resp.InstanceCount)
	require.Equal(t, "2018-02-24T13:13:00.236Z", resp.CreatedAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "2018-02-24T13:13:00.236Z", resp.UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
}

func TestStatisticsList(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"statistics":[{"_id":"v3D4mvobwfznKozH8","uniqueId":"wD4EP3M7FeFzJZgk9","installedAt":"2018-02-18T19:40:45.369Z","version":"0.61.0-develop","totalUsers":88,"activeUsers":88,"nonActiveUsers":0,"onlineUsers":0,"awayUsers":1,"offlineUsers":87,"totalRooms":81,"totalChannels":41,"totalPrivateGroups":37,"totalDirect":3,"totlalLivechat":0,"totalMessages":2408,"totalChannelMessages":730,"totalPrivateGroupMessages":1869,"totalDirectMessages":25,"totalLivechatMessages":0,"lastLogin":"2018-02-24T12:44:45.045Z","lastMessageSentAt":"2018-02-23T18:14:03.490Z","lastSeenSubscription":"2018-02-23T17:58:54.779Z","instanceCount":1,"createdAt":"2018-02-24T15:13:00.312Z","_updatedAt":"2018-02-24T15:13:00.312Z"}],"count":1,"offset":0,"total":1,"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.StatisticsList()
	require.NoError(t, err)

	require.Equal(t, 1, resp.Count)
	require.Equal(t, 0, resp.Offset)
	require.Equal(t, 1, resp.Total)
	require.Equal(t, "v3D4mvobwfznKozH8", resp.Statistics[0].ID)
	require.Equal(t, "wD4EP3M7FeFzJZgk9", resp.Statistics[0].UniqueID)
	require.Equal(t, "2018-02-18T19:40:45.369Z", resp.Statistics[0].InstalledAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "0.61.0-develop", resp.Statistics[0].Version)
	require.Equal(t, 88, resp.Statistics[0].TotalUsers)
	require.Equal(t, 88, resp.Statistics[0].ActiveUsers)
	require.Equal(t, 0, resp.Statistics[0].NonActiveUsers)
	require.Equal(t, 0, resp.Statistics[0].OnlineUsers)
	require.Equal(t, 1, resp.Statistics[0].AwayUsers)
	require.Equal(t, 87, resp.Statistics[0].OfflineUsers)
	require.Equal(t, 81, resp.Statistics[0].TotalRooms)
	require.Equal(t, 41, resp.Statistics[0].TotalChannels)
	require.Equal(t, 37, resp.Statistics[0].TotalPrivateGroups)
	require.Equal(t, 3, resp.Statistics[0].TotalDirect)
	require.Equal(t, 0, resp.Statistics[0].TotalLivechat)
	require.Equal(t, 2408, resp.Statistics[0].TotalMessages)
	require.Equal(t, 730, resp.Statistics[0].TotalChannelMessages)
	require.Equal(t, 1869, resp.Statistics[0].TotalPrivateGroupMessages)
	require.Equal(t, 25, resp.Statistics[0].TotalDirectMessages)
	require.Equal(t, 0, resp.Statistics[0].TotalLivechatMessages)
	require.Equal(t, "2018-02-24T12:44:45.045Z", resp.Statistics[0].LastLogin.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "2018-02-23T18:14:03.490Z", resp.Statistics[0].LastMessageSentAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "2018-02-23T17:58:54.779Z", resp.Statistics[0].LastSeenSubscription.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, 1, resp.Statistics[0].InstanceCount)
	require.Equal(t, "2018-02-24T15:13:00.312Z", resp.Statistics[0].CreatedAt.Format("2006-01-02T15:04:05.000Z"))
	require.Equal(t, "2018-02-24T15:13:00.312Z", resp.Statistics[0].UpdatedAt.Format("2006-01-02T15:04:05.000Z"))
}
