package main

import (
	"github.com/docker/docker/client"
	"github.com/qqMelon/mynotory-minecraft/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Unabel to create docker client, please make sure that docker is installed\n%s", err.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IndexHandler(w, r, cli)
	})
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}
