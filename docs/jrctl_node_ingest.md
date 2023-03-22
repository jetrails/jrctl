## jrctl node ingest

Ingest node token and save it to config

```
jrctl node ingest [flags]
```

### Examples

```
echo -n TOKEN | jrctl node ingest -t localhost
echo -n TOKEN | jrctl node ingest -t jump -e 10.10.10.7
echo -n TOKEN | jrctl node ingest -t web -e 10.10.10.6 -f
```

### Options

```
  -e, --endpoint string   filter nodes using this endpoint (default "127.0.0.1:27482")
  -f, --force             create new entry if no matching nodes were found
  -h, --help              help for ingest
  -q, --quiet             output only errors
  -t, --type strings      types to attach to found nodes
```

### SEE ALSO

* [jrctl node](jrctl_node.md)	 - Manage configured nodes

