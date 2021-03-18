## jrctl firewall list

List firewall entries

### Synopsis

List firewall entries. Ask server(s) for a list of firewall entries. Specifing a
type selector will only query servers with that type. Not specifing any type
will show query all configured servers.

```
jrctl firewall list [flags]
```

### Examples

```
jrctl firewall list
jrctl firewall list -t admin
jrctl firewall list -t db
```

### Options

```
  -h, --help          help for list
  -t, --type string   specify server type selector
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

