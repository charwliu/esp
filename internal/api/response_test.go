package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"go.vixal.xyz/esp/internal/i18n"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	t.Run("already exists", func(t *testing.T) {
		resp := NewResponse(http.StatusConflict, i18n.ErrAlreadyExists, "A cat")
		assert.Equal(t, http.StatusConflict, resp.Code)
		assert.Equal(t, "A cat already exists", resp.Err)
		assert.Equal(t, "", resp.Message)
	})

	t.Run("unexpected error", func(t *testing.T) {
		resp := NewResponse(http.StatusInternalServerError, i18n.ErrUnexpected, "A cat")
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, "Unexpected error, please try again", resp.Err)
		assert.Equal(t, "", resp.Message)
	})

	t.Run("changes saved", func(t *testing.T) {
		resp := NewResponse(http.StatusOK, i18n.MsgChangesSaved)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "", resp.Err)
		assert.Equal(t, "Changes successfully saved", resp.Message)

		if s, err := json.Marshal(resp); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, `{"code":200,"message":"Changes successfully saved"}`, string(s))
		}
	})
}

func TestResponse_String(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		resp := Response{Code: 404, Err: "Not found", Message: "page not found", Details: "xyz"}
		assert.Equal(t, "Not found", resp.String())
	})

	t.Run("no error", func(t *testing.T) {
		t.Run("error", func(t *testing.T) {
			resp := Response{Code: 200, Message: "Ok", Details: "xyz"}
			assert.Equal(t, "Ok", resp.String())
		})
	})
}

func TestResponse_LowerString(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		resp := Response{Code: 404, Err: "Not found", Message: "page not found", Details: "xyz"}
		assert.Equal(t, "not found", resp.LowerString())
	})

	t.Run("no error", func(t *testing.T) {
		t.Run("error", func(t *testing.T) {
			resp := Response{Code: 200, Message: "Ok", Details: "xyz"}
			assert.Equal(t, "ok", resp.LowerString())
		})
	})
}

func TestResponse_Error(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		resp := Response{Code: 404, Err: "Not found", Message: "page not found", Details: "xyz"}
		assert.Equal(t, "Not found", resp.Error())
	})

	t.Run("no error", func(t *testing.T) {
		t.Run("error", func(t *testing.T) {
			resp := Response{Code: 200, Message: "Ok", Details: "xyz"}
			assert.Equal(t, "", resp.Error())
		})
	})
}

func TestResponse_Success(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		resp := Response{Code: 404, Err: "Not found", Message: "page not found", Details: "xyz"}
		assert.Equal(t, false, resp.Success())
	})

	t.Run("no error", func(t *testing.T) {
		t.Run("error", func(t *testing.T) {
			resp := Response{Code: 200, Message: "Ok", Details: "xyz"}
			assert.Equal(t, true, resp.Success())
		})
	})
}
