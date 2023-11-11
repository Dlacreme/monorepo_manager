package config

func EmptyConfigFile() []byte {
	return []byte(`###
### Monorepo Manager config file
###

name = "My Monorepo App"

[[workspaces]]
name = "$root"
folder = "."
setup = []
start = ["docker compose up -d", "source .env"]
stop = []

# first workspace is defined using only the required fields
[[workspaces]]
name = "workspace1"

# second workspace is defined using all fields available
[[workspaces]]
name = "workspace2"
folder = "second_workspace"
`)
}
