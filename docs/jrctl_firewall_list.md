## jrctl firewall list

List firewall entries

### Synopsis

List firewall entries. Ask daemon(s) for a list of firewall entries. Specifing a
type selector will only query daemons with that type. Not specifing any type
will show query all configured daemons.

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
  -h, --help          help for list
  -t, --type string   specify daemon type selector
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

