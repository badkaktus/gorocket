package gorocket

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSupportedLanguage(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"languages":[{"language":"af","name":"Africâner"},{"language":"sq","name":"Albanês"},{"language":"de","name":"Alemão"},{"language":"am","name":"Amárico"}],"success":true}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	resp, err := client.GetSupportedLanguage("")
	require.NoError(t, err)

	require.Equal(t, "af", resp.Languages[0].Language)
	require.Equal(t, "Africâner", resp.Languages[0].Name)
	require.Equal(t, "sq", resp.Languages[1].Language)
	require.Equal(t, "Albanês", resp.Languages[1].Name)
	require.Equal(t, "de", resp.Languages[2].Language)
	require.Equal(t, "Alemão", resp.Languages[2].Name)
	require.Equal(t, "am", resp.Languages[3].Language)
	require.Equal(t, "Amárico", resp.Languages[3].Name)
	require.True(t, resp.Success)
}
