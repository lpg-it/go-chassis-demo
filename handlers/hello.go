package hello

import (
	"github.com/go-chassis/go-chassis/server/restful"
	"net/http"
)

type Hello struct {
}

func (s *Hello) SayHello(b *restful.Context) {
	b.Write([]byte("get user id: " + b.ReadPathParameter("userid")))
}

func (s *Hello) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodPost, Path: "/sayhello/{userid}", ResourceFunc: s.SayHello,
			Returns: []*restful.Returns{{Code: 200}}}}
}