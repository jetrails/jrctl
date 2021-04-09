## jrctl server list

List servers in configured deployment

### Synopsis

List servers in configured deployment. Specifing a server type will only display
results for servers of that type.

```
jrctl server list [flags]
```

### Examples

```
jrctl server list
jrctl server list -t admin
jrctl server list -t localhost
jrctl server list -t www
```

### Options

```
  -h, --help          help for list
  -t, --type string   specify server type selector
```

### SEE ALSO

* [jrctl server](jrctl_server.md)	 - Interact with servers in configured deployment

