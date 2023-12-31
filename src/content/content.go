package content

const (
	// error
	InvalidCommandLine string = "Couldn't find a valid command to run. Use 'help' for help."
	ExpectArgument     string = "expects at least one argument. Use 'help' for help."
	CommandNotFound    string = "Command is not found."
	FileExisting       string = "file already existing."
	// success
	FileCreated string = "file successfully created."
	// commands
	NoWorkspaceAvailable   string = "No workspace available."
	WorkspaceNotFound      string = "workspace not found."
	NowUsingWorkspace      string = "Now using workspace"
	NoWorkspaceInUse       string = "No workspace in use"
	WorkspaceInUseNotFound string = "in use but this workspace does not exists."
	NothingToDo            string = "Nothing to do."
)
