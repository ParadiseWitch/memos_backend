package router

import (
	"memos/common/global"
)

func InitRouter() {
	r := global.Engine
	v1 := r.Group("/api/v1")
	for _, f := range global.Routers {
		f(v1)
	}
	r.Run()
}
