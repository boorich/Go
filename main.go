package main

import (
	"net/http"

	"github.com/empea-careercriminal/Go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
