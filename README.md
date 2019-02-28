# alauda-cli
The command-line interface for Alauda.io.

## Usage
```
$ alauda
Alauda CLI

Usage:
  alauda [command]

Available Commands:
  app         Manage applications
  apps        List applications
  cluster     Manage clusters
  clusters    List clusters
  help        Help about any command
  image       Manage images
  images      List images
  inspect     Inspect an application
  login       Log onto the Alauda platform
  logout      Log out of the Alauda platform
  ls          List applications
  namespace   Manage namespaces
  namespaces  List namespaces
  project     Manage projects
  projects    List projects
  registries  List registries
  registry    Manage registries
  rm          Remove an application
  run         Run an application
  space       Manage spaces
  spaces      List spaces
  start       Start an application
  stop        Stop an application
  version     Display version of Alauda CLI
  yaml        Retrieve the YAML of an application

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