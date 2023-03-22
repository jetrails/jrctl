## jrctl firewall list

List firewall entries across configured nodes

### Synopsis

List firewall entries across configured nodes. Specifing a tag will display
nodes that have that tag.

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
  -h, --help              help for list
  -q, --quiet             display number of entries found for each matching server
  -t, --tag stringArray   filter nodes using tags (default [default])
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with server firewall

