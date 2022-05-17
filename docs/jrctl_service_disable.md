## jrctl service disable

Disable specified service(s) running on configured server(s)

### Synopsis

Disable specified service(s) running on configured server(s). Services can be
repeated and execution will happen in the order that is given. Specifing a
server type will only display results for servers of that type.

```
jrctl service disable SERVICE... [flags]
```

### Examples

```
jrctl service disable nginx
jrctl service disable nginx varnish
jrctl service disable nginx varnish php-fpm
jrctl service disable nginx varnish php-fpm-7.2 nginx
```

### Options

```
  -h, --help           help for disable
  -q, --quiet          output as little information as possible
  -t, --type strings   specify server type(s), useful for cluster (default [localhost])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in configured deployment

