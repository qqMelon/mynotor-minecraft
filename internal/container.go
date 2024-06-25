package internal

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/qqMelon/mynotory-minecraft/types"
	"log"
)

func ListContainers(c *client.Client) ([]types.Container, error) {
	filterArgs := filters.NewArgs()
	filterArgs.Add("name", "minecraft")
	containers, err := c.ContainerList(context.Background(), container.ListOptions{
		Size:    false,
		All:     false,
		Latest:  false,
		Since:   "",
		Before:  "",
		Limit:   0,
		Filters: filterArgs,
	})
	if err != nil {
		log.Fatal("Unable to get containers, please make sure that docker daemon is up and running")
	}

	var containersList []types.Container
	for _, t := range containers {

		var networks []types.ContainerNetwork
		for _, p := range t.Ports {
			networks = append(networks, types.ContainerNetwork{
				IP:          p.IP,
				PrivatePort: p.PrivatePort,
				PublicPort:  p.PublicPort,
				Type:        p.Type,
			})
		}
		containersList = append(containersList, types.Container{
			Name:    t.Names[0],
			ID:      t.ID,
			Image:   t.Image,
			Ports:   networks,
			Command: t.Command,
			State:   t.State,
			Status:  t.Status,
		})
	}

	return containersList, nil
}
