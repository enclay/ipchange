# IPchange

* Allows to execute shell command when your ip changes
* Exports new public IP to environment
* Specify check interval with **--each** flag

### Usage:

```sh
ipchange --on-change "caddy reload" --each 1m
ipchange --on-change "caddy reload" --each 30s

```
