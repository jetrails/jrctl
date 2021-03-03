## jrctl firewall list

List firewall entries

### Synopsis

List firewall entries. Ask daemon(s) for a list of firewall entries. Specifing a
tag selector will only query daemons with that tag. Not specifing any tag will
show query all configured daemons.

```
jrctl firewall list [flags]
```

### Examples

```
jrctl firewall list
jrctl firewall list -t admin
jrctl firewall list -t mysql
```

### Options

```
  -t, --tag string   specify daemon tag selector
  -h, --help         help for list
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

