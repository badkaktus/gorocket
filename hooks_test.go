package gorocket

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHooks(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	msg := HookMessage{
		Text: "Hello",
		Attachments: []HookAttachment{
			{
				Title:     "Title",
				TitleLink: "http://example.com",
				Text:      "Text",
				ImageURL:  "http://example.com/image.png",
				Color:     "#ff0000",
			},
		},
	}

	resp, err := client.Hooks(&msg, "token")
	require.NoError(t, err)

	require.True(t, resp.Success)
}
