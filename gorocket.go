package gorocket

import (
	"context"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"log"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	userID     string
	xToken     string
	apiVersion string
	HTTPClient *http.Client

	timeout time.Duration
}

type PaginationStruct struct {
	Count  int    `url:"count,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Sort   string `url:"sort,omitempty"`
}

var pagination PaginationStruct

// NewClient creates new rocket.chat client with given API key
func NewClient(url string) *Client {
	return &Client{
		//userID: user,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		//xToken:     token,
		baseURL:    url,
		apiVersion: "api/v1",
	}
}

// NewWithOptions creates new rocket.chat client with options
func NewWithOptions(url string, opts ...Option) *Client {
	c := &Client{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		baseURL:    url,
		apiVersion: "api/v1",
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

type Option func(*Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.timeout = d
	}
}

func WithUserID(userID string) Option {
	return func(c *Client) {
		c.userID = userID
	}
}

func WithXToken(xtoken string) Option {
	return func(c *Client) {
		c.xToken = xtoken
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("X-Auth-Token", c.xToken)
	req.Header.Add("X-User-Id", c.userID)

	if c.timeout > 0 {
		ctx, cancel := context.WithTimeout(req.Context(), c.timeout)
		defer cancel()

		req = req.WithContext(ctx)
	}

	res, err := c.HTTPClient.Do(c.addQueryParams(req))
	if err != nil {
		log.Println(err)
		return err
	}

	c.cleanup()

	defer res.Body.Close()

	resp := v
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}

	return nil
}

func (c *Client) Count(val int) *Client {
	pagination.Count = val
	return c
}

func (c *Client) Offset(val int) *Client {
	pagination.Offset = val
	return c
}

func (c *Client) Sort(val map[string]int) *Client {
	byteJson, err := json.Marshal(val)
	if err != nil {
		log.Printf("cant create sort. error: %s", err)
		return c
	}
	pagination.Sort = string(byteJson)

	return c
}

func (c *Client) addQueryParams(req *http.Request) *http.Request {
	v, err := query.Values(pagination)
	if err != nil {
		log.Printf("error create query string: %s", err)
		return req
	}
	q := req.URL.Query()
	for k := range v {
		q.Add(k, v.Get(k))
	}
	req.URL.RawQuery = q.Encode()
	return req
}

func (c *Client) cleanup() {
	pagination.Sort = ""
	pagination.Offset = 0
	pagination.Count = 0
}
