package gorocket

import (
	"fmt"
	"net/http"
)

func (c *Client) Info() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/info", c.baseURL), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	//if err := c.sendRequest(req); err != nil {
	//	return err
	//}

	return nil
}
