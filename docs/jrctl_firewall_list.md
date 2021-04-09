## jrctl firewall list

List firewall entries across configured servers

### Synopsis

List firewall entries across configured servers. Specifing a server type will
only display results for servers of that type.

```
jrctl firewall list [flags]
```

### Examples

```
  jrctl firewall list
  jrctl firewall list -t admin
  jrctl firewall list -t db
  jrctl firewall list -t www
```

### Options

```
  -h, --help          help for list
  -t, --type string   specify server type selector
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

