package types

type Container struct {
	Name    string
	ID      string
	Image   string
	Ports   []ContainerNetwork
	State   string
	Status  string
	Command string
}

type ContainerNetwork struct {
	IP          string
	PrivatePort uint16
	PublicPort  uint16
	Type        string
}
