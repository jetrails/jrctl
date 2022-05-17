## jrctl service restart

Restart specified service(s) running on configured server(s)

### Synopsis

Restart specified service(s) running on configured server(s). In order to
successfully restart a service, the server first validates the respected
service's config file. Once deemed valid, the service is restarted. Services can
be repeated and execution will happen in the order that is given. Specifing a
server type will only display results for servers of that type.

```
jrctl service restart SERVICE... [flags]
```

### Examples

```
jrctl service restart nginx
jrctl service restart nginx varnish
jrctl service restart nginx varnish php-fpm
jrctl service restart nginx varnish php-fpm-7.2 nginx
```

### Options

```
  -h, --help           help for restart
  -q, --quiet          output as little information as possible
  -t, --type strings   specify server type(s), useful for cluster (default [localhost])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in configured deployment

