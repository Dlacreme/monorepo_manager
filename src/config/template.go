package config

func EmptyConfigFile() []byte {
	return []byte(`###
### Monorepo Manager config file
###

name = "My Monorepo App"
`)
}
