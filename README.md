# Monorepo Manager (mm)

An easy to use application to help you setup & navigate your monorepo. It's a standalone binary that requires nothing but itself.

## Usage

### Commands


| Command (short name) | Description                           | Example               |
| -------------------- | ------------------------------------- | --------------------- |
| init (i)             | Create a default 'monorepo.toml' file | mm init               |
| help (h)             | Display more information about 'mm'   | mm h                  |
| list (l)             | List all workspace available          | mm l                  |
| use (u)              | Set the workspace to use              | mm u {workspace_name} |
| workspace (w)        | Get the workspace currently in use    | mm w                  |

*See `mm help` for more information.*

### Setup

Create a default 'monorepo.toml' file with
```shell
$ mm init
$ mm help
```

### Workspaces

You can define as many workspaces as you wish in your 'monorepo.toml' file. See **config** section for more information about the configuration of workspace.

Once your workspace have been properly defined, you can use `list` & `use` to set the workspace you wish to work with and `workspace` to check the workspace currently in use.

The **`_root`** workspace refers to the root of your monorepo and can be used to run commands that are common to all your workspaces.

```shell
$ mm list # display all workspace available
 - _root
 - hello_work
 - bye_work
$ mm use hello_work # set the workspace to use
 now using 'hello_world'
$ mm workspace # display the workspace currently in use
 using 'hello_world'
$ mm setup # run the 'setup' command for 'hello_world'
# ...
```

You can use `{workspace_name}${command}` to target a specific workspace without having to change your current workspace. If you don't pass a workspace name before '$' the _root workspace will be used.

```shell
$ mm workspace
 using 'hello_world'
$ mm $setup # run 'setup' from the _root workspace
 # ...
$ mm bye_work$setup # run 'setup' from the 'bye_work' workspace
```

*You can override the workspace currently in use by setting manually the `MM_WORKSPACE` ENV variable.*

## Install

Clone this repository and then simply run:
```shell
$ make install
```

It will compile and move the binary to your ~/.local/bin/ folder.

*If you are not using ~/.local/bin simply use `mv mm /your/desired/path`*
