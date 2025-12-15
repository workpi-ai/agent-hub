package registry

import (
	"embed"
)

const (
	AgentsDir   = "agents"
	CommandsDir = "commands"
)

//go:embed agents
var Agents embed.FS

//go:embed commands
var Commands embed.FS

//go:embed agents commands
var All embed.FS
