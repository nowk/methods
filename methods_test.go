package methods

import "net/http"
import "testing"
import "github.com/nowk/assert"

func TestAllowMethods(t *testing.T) {
	m := Allow("get", "POST")

	for _, v := range []struct {
		method string
		b      bool
	}{
		{"GET", true},
		{"POST", true},
		{"DELETE", false},
	} {
		b := m.Allowed(&http.Request{
			Method: v.method,
		})
		assert.Equal(t, b, v.b)
	}
}
