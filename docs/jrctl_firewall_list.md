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
  -h, --help               help for list
  -q, --quiet              display number of entries found for each matching server
  -t, --type stringArray   filter servers using type selectors (default [localhost])
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with server firewall

