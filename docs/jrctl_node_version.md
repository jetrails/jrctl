## jrctl node version

Display daemon version running on configured nodes

### Synopsis

Display daemon version running on configured nodes. Specifing a node type will
only display results for nodes of that type.

```
jrctl node version [flags]
```

### Examples

```
jrctl node version
jrctl node version -t www
```

### Options

```
  -h, --help              help for version
  -q, --quiet             only display versions
  -t, --tag stringArray   filter nodes using tags
```

### SEE ALSO

* [jrctl node](jrctl_node.md)	 - Manage configured nodes

