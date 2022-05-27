## jrctl service list

List services with their statuses and abilities.

### Synopsis

List services with their statuses and abilities.. Specifing a server type will
only display results for servers of that type. Specifing the service will filter
the list of services to include those services.

```
jrctl service list [flags]
```

### Examples

```
jrctl service list
jrctl service list -t admin
jrctl service list -t localhost
jrctl service list -t www
```

### Options

```
  -h, --help               help for list
  -q, --quiet              display unique list of found services
  -s, --service strings    filter by service
  -t, --type stringArray   filter servers using type selectors (default [localhost])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in deployment

