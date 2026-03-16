package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tc1io/clean-arch-demo/adapter"

	"github.com/tc1io/clean-arch-demo/api"
	"github.com/tc1io/clean-arch-demo/ui"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// read env vars for config
	orderPlacer := adapter.OrderQueue{}

	app := httprouter.New()
	api.Register(orderPlacer, app)
	ui.Register(app)
	bindAddr := "127.0.0.1"
	httpPort := 8080
	addr := fmt.Sprintf("%s:%d", bindAddr, httpPort)
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(addr, app))
}
