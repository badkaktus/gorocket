package gorocket

import (
	"fmt"
	"strings"
)

type Response interface {
	OK(statusCode int, debug bool) error
}

// ErrStatus handles error responses.
type ErrStatus struct {
	Success   bool   `json:"success"`
	ErrorMsg  string `json:"error,omitempty"`
	ErrorType string `json:"errorType,omitempty"`

	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Details any    `json:"details,omitempty"`

	statusCode int
	debug      bool
}

func (s ErrStatus) Error() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "rocket.chat response error (status code %d): success:%t", s.statusCode, s.Success)

	if len(s.ErrorMsg) > 0 {
		fmt.Fprintf(&sb, ", error:%s", s.ErrorMsg)
	}
	if len(s.ErrorType) > 0 {
		fmt.Fprintf(&sb, ", errorType:%s", s.ErrorType)
	}
	if len(s.Status) > 0 {
		fmt.Fprintf(&sb, ", status:%s", s.Status)
	}
	if len(s.Message) > 0 {
		fmt.Fprintf(&sb, ", message:%s", s.Message)
	}

	if s.debug && s.Details != nil {
		fmt.Fprintf(&sb, ", details:%s", s.Details)
	}

	return sb.String()
}

func (s ErrStatus) OK(statusCode int, debug bool) error {
	s.statusCode = statusCode
	s.debug = debug

	// Only assume the request was successful if success is true, and/or status is success.
	if s.Success {
		return nil
	}

	if s.Status == "success" {
		return nil
	}

	return s
}
