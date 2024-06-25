package handlers

import (
	"github.com/docker/docker/client"
	"github.com/qqMelon/mynotory-minecraft/internal"
	"github.com/qqMelon/mynotory-minecraft/types"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, c *client.Client) {
	containers, err := internal.ListContainers(c)
	if err != nil {
		log.Fatal("Unable to retrieve containers data from internal.ListContainers")
	}
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
	err = tmpl.Execute(w, struct {
		Containers []types.Container
	}{
		Containers: containers,
	})
	if err != nil {
		return
	}
}
