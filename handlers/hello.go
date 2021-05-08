package hello

import (
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"net/http"
)

var err error

type Hello struct {
}


// Say 无参数解析
func (h *Hello) Say(b *rf.Context) {
	err = b.Write([]byte("Hello Chassis"))
	if err != nil {
		openlog.Fatal("h.Say Error " + err.Error())
		return
	}
}

func (h *Hello) URLPatterns() []rf.Route {
	return []rf.Route{
		{
			Method:       http.MethodGet,
			Path:         "/",
			ResourceFunc: h.Say,
			Returns:      []*rf.Returns{{Code: http.StatusOK}},
		},
	}
}