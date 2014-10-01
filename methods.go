package methods

import "net/http"
import "strings"

// HTTP Methods
const (
	GET    = "GET"
	POST   = "POST"
	HEAD   = "HEAD"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

// Bouncer struct around list of allowed HTTP methods
type Bouncer struct {
	allow []string
}

func toupper(s ...string) []string {
	for i, v := range s {
		s[i] = strings.ToUpper(v)
	}

	return s
}

// Allow returns a new Bouncer with given methods (upcased)
func Allow(m ...string) *Bouncer {
	return &Bouncer{
		allow: toupper(m...),
	}
}

// Allowed returns true if request method is included in the allow list
func (b Bouncer) Allowed(req *http.Request) bool {
	m := req.Method
	for _, v := range b.allow {
		if m == v {
			return true
		}
	}

	return false
}
