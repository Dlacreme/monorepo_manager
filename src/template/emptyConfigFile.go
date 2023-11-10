package template

func EmptyConfigFile() []byte {
	return []byte(`###
### Monorepo Manager config file
###

title = "My Monorepo App"
`)
}
