package routes

import "github.com/cubeee/steamtracker/web/bootstrap"

func Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", GetIndexHandler)
}
