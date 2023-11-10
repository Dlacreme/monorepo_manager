package command

type Command string

const (
	Init Command = "init"
)

func Run(cmd Command, params []string) {

}
