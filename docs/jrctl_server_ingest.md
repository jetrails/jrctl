## jrctl server ingest

Ingest server token and save it to config

```
jrctl server ingest [flags]
```

### Examples

```
echo -n TOKEN | jrctl server ingest -t localhost
echo -n TOKEN | jrctl server ingest -t jump -e 10.10.10.7
echo -n TOKEN | jrctl server ingest -t web -e 10.10.10.6 -f
```

### Options

```
  -e, --endpoint string   server endpoint used for new entries only (default "127.0.0.1:27482")
  -f, --force             create new entry even if matching entries exist
  -h, --help              help for ingest
  -q, --quiet             output as little information as possible
  -t, --type strings      filter servers using type selectors, all must match
```

### SEE ALSO

* [jrctl server](jrctl_server.md)	 - Manage servers in deployment

