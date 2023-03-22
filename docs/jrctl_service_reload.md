## jrctl service reload

Reload specified services in deployment

### Synopsis

Reload specified services in deployment. In order to successfully reload a
service, the server first validates the respected service's config file. Once
deemed valid, the service is reloaded. Is passed service does not support
reloading, then a restart will happen instead. Services can be repeated and
execution will happen in the order that is given. Specifing a tag will display
nodes that have that tag.

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
  -h, --help              help for reload
  -q, --quiet             display no output
  -t, --tag stringArray   filter nodes using tags (default [default])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in deployment

