package main

import (
	"net/http"

	"github.com/timothykimble/gettingstartedgo/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
