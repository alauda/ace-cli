# alauda-cli
The command-line interface for Alauda.io.

## Usage
```
$ alauda
Alauda CLI

Usage:
  alauda [command]

Available Commands:
  create      Create a new service
  help        Help about any command
  login       Log onto the Alauda platform
  logout      Log out of the Alauda platform
  ps          List services
  rm          Remove a service
  run         Create and start a new service
  service     Manage services
  start       Start a service
  stop        Stop a service
  version     Display version of Alauda CLI

Flags:
      --config string   config file (default is $HOME/.alauda.yml)
  -h, --help            help for alauda

Use "alauda [command] --help" for more information about a command.
```

## Running Tests
1. Use `alauda login` to log into an Alauda account.
2. Add the following settings to the config file (default at `$HOME/.alauda.yml`):
```
test:
    cluster: <mycluster>
    space: <myspace>
    name: <mytestservicename>
    image: <mytestimage>
```
3. Run `go test`.