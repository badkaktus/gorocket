package gorocket

import (
	"fmt"
	"net/http"
)

type SupportedLanguageResp struct {
	ErrStatus
	Languages []language `json:"languages"`
}

type language struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}

func (c *Client) GetSupportedLanguage(query string) (*SupportedLanguageResp, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/autotranslate.getSupportedLanguages", c.baseURL, c.apiVersion), nil)
	if err != nil {
		return nil, err
	}

	if query != "" {
		q := req.URL.Query()
		q.Add("targetLanguage", query)
		req.URL.RawQuery = q.Encode()
	}

	res := SupportedLanguageResp{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
