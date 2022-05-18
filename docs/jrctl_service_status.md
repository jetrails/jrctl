## jrctl service status

List services with their statuses and abilities.

### Synopsis

List services with their statuses and abilities.. Specifing a server type will
only display results for servers of that type. Specifing the service will filter
the list of services to include those services.

```
jrctl service status [flags]
```

### Examples

```
jrctl service status
jrctl service status -t admin
jrctl service status -t localhost
jrctl service status -t www
```

### Options

```
  -h, --help              help for status
  -s, --service strings   filter by service
  -t, --type strings      specify server type(s) selector
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in configured deployment

