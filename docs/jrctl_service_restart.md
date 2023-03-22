## jrctl service restart

Restart specified services in deployment

### Synopsis

Restart specified services in deployment. In order to successfully restart a
service, the server first validates the respected service's config file. Once
deemed valid, the service is restarted. Services can be repeated and execution
will happen in the order that is given. Specifing a tag will display nodes that
have that tag.

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
  -h, --help              help for restart
  -q, --quiet             display no output
  -t, --tag stringArray   filter nodes using tags (default [localhost])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in deployment

