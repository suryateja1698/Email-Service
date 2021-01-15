package main

import (
	"emailservice/handler"
	"fmt"
	"net/http"
	"os"

	"github.com/go-delve/delve/pkg/config"
	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)

func main() {
	config.LoadConfig()
	router := handler.NewRouter()
	http.Handle("/", router)
	//	os.Setenv("PORT", "8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	n := negroni.Classic()

	n.UseHandler(handlers.CORS()(router))

	fmt.Println("Listening on port\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), n)
	if err != nil {
		panic(err.Error())
	}

}
