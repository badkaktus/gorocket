package gorocket

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type HandlerHelper struct {
	Code         int
	ResponseBody string
}

func getHandler(t *testing.T, param *HandlerHelper) http.HandlerFunc {
	httpStatus := param.Code
	if httpStatus == 0 {
		httpStatus = http.StatusOK
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		_, err := w.Write([]byte(param.ResponseBody))
		require.NoError(t, err)
	})
}
