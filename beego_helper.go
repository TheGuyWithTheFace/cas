package cas

import (
	"net/http"

	"github.com/golang/glog"
)

type BeegoCASData struct {
	userName string
}

func (bcd *BeegoCASData) GetUserName() string {
	return bcd.userName
}

func ServeBeego(w http.ResponseWriter, r *http.Request, c *Client) *BeegoCASData {
	if glog.V(2) {
		glog.Infof("cas: handling %v request for %v", r.Method, r.URL)
	}

	setClient(r, c)
	defer clear(r)

	c.getSession(w, r)

	if !IsAuthenticated(r) {
		RedirectToLogin(w, r)
		return nil
	}

	if r.URL.Path == "/logout" {
		RedirectToLogout(w, r)
		return nil
	}

	return &BeegoCASData{
		userName: Username(r),
	}
}
