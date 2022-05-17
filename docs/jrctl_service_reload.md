## jrctl service reload

Reload specified service(s) running on configured server(s)

### Synopsis

Reload specified service(s) running on configured server(s). In order to
successfully reload a service, the server first validates the respected
service's config file. Once deemed valid, the service is reloaded. Is passed
service does not support reloading, then a restart will happen instead. Services
can be repeated and execution will happen in the order that is given. Specifing
a server type will only display results for servers of that type.

```
jrctl service reload SERVICE... [flags]
```

### Examples

```
jrctl service reload nginx
jrctl service reload nginx varnish
jrctl service reload nginx varnish php-fpm
jrctl service reload nginx varnish php-fpm-7.2 nginx
```

### Options

```
  -h, --help           help for reload
  -q, --quiet          output as little information as possible
  -t, --type strings   specify server type(s), useful for cluster (default [localhost])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in configured deployment

