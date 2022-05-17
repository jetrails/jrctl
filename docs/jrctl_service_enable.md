## jrctl service enable

Enable specified service(s) running on configured server(s)

### Synopsis

Enable specified service(s) running on configured server(s). Services can be
repeated and execution will happen in the order that is given. Specifing a
server type will only display results for servers of that type.

```
jrctl service enable SERVICE... [flags]
```

### Examples

```
jrctl service enable nginx
jrctl service enable nginx varnish
jrctl service enable nginx varnish php-fpm
jrctl service enable nginx varnish php-fpm-7.2 nginx
```

### Options

```
  -h, --help           help for enable
  -q, --quiet          output as little information as possible
  -t, --type strings   specify server type(s), useful for cluster (default [localhost])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in configured deployment

