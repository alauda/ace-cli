# alauda-cli
The command-line interface for Alauda.io.

## Usage
```
$ alauda
Alauda CLI

Usage:
  alauda [command]

Available Commands:
  cluster     Manage clusters
  compose     Manage application compose
  config      Manage configurations
  create      Create a new service
  help        Help about any command
  image       Manage images
  images      List images
  inspect     Inspect a service
  lb          Manage load balancers
  login       Log onto the Alauda platform
  logout      Log out of the Alauda platform
  node        Manage nodes
  nodes       List nodes
  ps          List services
  registry    Manage registries
  restart     Restart a service
  rm          Remove a service
  run         Create and start a new service
  scale       Scale a service to the specified number of instances
  service     Manage services
  space       Manage spaces
  start       Start a service
  stop        Stop a service
  template    Manage application templates
  update      Update a service
  version     Display version of Alauda CLI
  volume      Manage volumes

Flags:
      --config string   config file (default: $HOME/.alauda.yml)
  -h, --help            help for alauda

Use "alauda [command] --help" for more information about a command.
```

## Running Tests
1. Use `alauda login` to log into an Alauda account.
2. Add the following settings to the config file (default at `$HOME/.alauda.yml`):
```
test:
  app: <app name>
  cluster: <cluster name>
  config: <config name>
  image: <image name>
  lb: <loadbalancer name>
  registry: <registry name>
  registryproject: <registry project name>
  repo: <registry repository name>
  service: <service name>
  space: <space name>
  template: <template name>
  volume: <volume name>
```
3. Run `go test`.