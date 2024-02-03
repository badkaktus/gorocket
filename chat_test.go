package gorocket

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostMessage(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"message":{"rid":"GENERAL","msg":"123456789","ts":"2018-03-01T18:02:26.825Z","u":{"_id":"i5FdM4ssFgAcQP62k","username":"rocket.cat","name":"test"},"unread":true,"mentions":[],"channels":[],"_updatedAt":"2018-03-01T18:02:26.828Z","_id":"LnCSJxxNkCy6K9X8X"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	payload := Message{
		Text: "Hello",
	}
	resp, err := client.PostMessage(&payload)
	require.NoError(t, err)

	require.Equal(t, "GENERAL", resp.Message.Rid)
	require.Equal(t, "123456789", resp.Message.Msg)
	require.Equal(t, "2018-03-01T18:02:26.825Z", resp.Message.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "i5FdM4ssFgAcQP62k", resp.Message.U.ID)
	require.Equal(t, "rocket.cat", resp.Message.U.Username)
	require.Equal(t, "2018-03-01T18:02:26.828Z", resp.Message.UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "LnCSJxxNkCy6K9X8X", resp.Message.ID)
	require.True(t, resp.Success)
}

func TestGetMessage(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"message":{"_id":"7aDSXtjMA3KPLxLjt","rid":"GENERAL","msg":"This is a test!","ts":"2016-12-14T20:56:05.117Z","u":{"_id":"y65tAmHs93aDChMWu","username":"graywolf336"}},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SingleMessageId{
		MessageId: "7aDSXtjMA3KPLxLjt",
	}

	resp, err := client.GetMessage(&req)
	require.NoError(t, err)

	require.Equal(t, "7aDSXtjMA3KPLxLjt", resp.Message.ID)
	require.Equal(t, "GENERAL", resp.Message.Rid)
	require.Equal(t, "This is a test!", resp.Message.Msg)
	require.Equal(t, "2016-12-14T20:56:05.117Z", resp.Message.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "y65tAmHs93aDChMWu", resp.Message.U.ID)
	require.Equal(t, "graywolf336", resp.Message.U.Username)
	require.True(t, resp.Success)
}

func TestDeleteMessage(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"_id":"jEnjsxuoDJamGjbH2","ts":1696533809813,"message":{"_id":"jEnjsxuoDJamGjbH2","rid":"6GFJ3tbmHiyHbahmC","u":{"_id":"5fRTXMt7DMJbpPJfh","username":"test.funke","name":"TestFunke"}},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := DeleteMessageRequest{
		RoomID: "6GFJ3tbmHiyHbahmC",
		MsgID:  "jEnjsxuoDJamGjbH2",
	}

	resp, err := client.DeleteMessage(&req)
	require.NoError(t, err)

	require.Equal(t, "jEnjsxuoDJamGjbH2", resp.ID)
	require.Equal(t, int64(1696533809813), resp.Ts)
	require.True(t, resp.Success)
}

func TestGetPinnedMessages(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"messages":[{"_id":"Srhca3mgthgjkEisJ","rid":"ByehQjC44FwMeiLbX","msg":"I pinned this message","ts":"2019-03-23T00:53:24.388Z","u":{"_id":"aobEdbYhXfu5hkeqG","username":"user","name":"User"},"mentions":[],"channels":[],"_updatedAt":"2019-03-23T00:53:28.813Z","pinned":true,"pinnedAt":"2019-03-23T00:53:28.813Z","pinnedBy":{"_id":"aobEdbYhXfu5hkeqG","username":"user"}},{"_id":"m3AZcKrvayKEZSKJN","rid":"GENERAL","msg":"Ola","ts":"2019-03-23T00:53:50.974Z","u":{"_id":"aobEdbYhXfu5hkeqG","username":"user","name":"user"},"mentions":[],"channels":[],"_updatedAt":"2019-03-23T00:53:53.649Z","pinned":true,"pinnedAt":"2019-03-23T00:53:53.649Z","pinnedBy":{"_id":"aobEdbYhXfu5hkeqG","username":"user"}}],"count":2,"offset":0,"total":2,"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := GetPinnedMsgRequest{
		RoomId: "GENERAL",
		Count:  2,
		Offset: 0,
	}

	resp, err := client.Sort(map[string]int{"ts": -1}).GetPinnedMessages(&req)
	require.NoError(t, err)

	require.Equal(t, 2, resp.Count)
	require.Equal(t, 0, resp.Offset)
	require.Equal(t, 2, resp.Total)
	require.True(t, resp.Success)
	require.Equal(t, "Srhca3mgthgjkEisJ", resp.Messages[0].ID)
	require.Equal(t, "ByehQjC44FwMeiLbX", resp.Messages[0].Rid)
	require.Equal(t, "I pinned this message", resp.Messages[0].Msg)
	require.Equal(t, "2019-03-23T00:53:24.388Z", resp.Messages[0].Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Messages[0].U.ID)
	require.Equal(t, "user", resp.Messages[0].U.Username)
	require.Equal(t, "User", resp.Messages[0].U.Name)
	require.Equal(t, 0, len(resp.Messages[0].Mentions))
	require.Equal(t, 0, len(resp.Messages[0].Channels))
	require.Equal(t, "2019-03-23T00:53:28.813Z", resp.Messages[0].UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.True(t, resp.Messages[0].Pinned)
	require.Equal(t, "2019-03-23T00:53:28.813Z", resp.Messages[0].PinnedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Messages[0].PinnedBy.ID)
	require.Equal(t, "user", resp.Messages[0].PinnedBy.Username)
	require.Equal(t, "m3AZcKrvayKEZSKJN", resp.Messages[1].ID)
	require.Equal(t, "GENERAL", resp.Messages[1].Rid)
	require.Equal(t, "Ola", resp.Messages[1].Msg)
	require.Equal(t, "2019-03-23T00:53:50.974Z", resp.Messages[1].Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Messages[1].U.ID)
	require.Equal(t, "user", resp.Messages[1].U.Username)
	require.Equal(t, "user", resp.Messages[1].U.Name)
	require.Equal(t, 0, len(resp.Messages[1].Mentions))
	require.Equal(t, 0, len(resp.Messages[1].Channels))
	require.Equal(t, "2019-03-23T00:53:53.649Z", resp.Messages[1].UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.True(t, resp.Messages[1].Pinned)
	require.Equal(t, "2019-03-23T00:53:53.649Z", resp.Messages[1].PinnedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "aobEdbYhXfu5hkeqG", resp.Messages[1].PinnedBy.ID)
	require.Equal(t, "user", resp.Messages[1].PinnedBy.Username)
}

func TestPinMessage(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"message":{"t":"message_pinned","rid":"GENERAL","ts":"2017-09-27T20:39:57.921Z","msg":"","u":{"_id":"Z3cpiYN6CNK2oXWKv","username":"graywolf336"},"groupable":false,"attachments":[{"text":"Hello","author_name":"graywolf336","author_icon":"/avatar/graywolf336?_dc=0","ts":"2017-09-27T19:36:01.683Z"}],"_updatedAt":"2017-09-27T20:39:57.921Z","_id":"hmzxXKSWmMkoQyiAd"},"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SingleMessageId{
		MessageId: "jEnjsxuoDJamGjbH2",
	}

	resp, err := client.PinMessage(&req)
	require.NoError(t, err)

	require.Equal(t, "message_pinned", resp.Message.T)
	require.Equal(t, "GENERAL", resp.Message.Rid)
	require.Equal(t, "2017-09-27T20:39:57.921Z", resp.Message.Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "", resp.Message.Msg)
	require.Equal(t, "Z3cpiYN6CNK2oXWKv", resp.Message.U.ID)
	require.Equal(t, "graywolf336", resp.Message.U.Username)
	require.False(t, resp.Message.Groupable)
	require.Equal(t, "Hello", resp.Message.Attachments[0].Text)
	require.Equal(t, "graywolf336", resp.Message.Attachments[0].AuthorName)
	require.Equal(t, "/avatar/graywolf336?_dc=0", resp.Message.Attachments[0].AuthorIcon)
	require.Equal(t, "2017-09-27T19:36:01.683Z", resp.Message.Attachments[0].Ts.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "2017-09-27T20:39:57.921Z", resp.Message.UpdatedAt.Format("2006-01-02T15:04:05.999Z"))
	require.Equal(t, "hmzxXKSWmMkoQyiAd", resp.Message.ID)
	require.True(t, resp.Success)
}

func TestUnpinMessage(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)

	req := SingleMessageId{
		MessageId: "jEnjsxuoDJamGjbH2",
	}

	resp, err := client.UnpinMessage(&req)
	require.NoError(t, err)

	require.True(t, resp.Success)
}
