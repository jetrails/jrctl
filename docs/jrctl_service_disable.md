## jrctl service disable

Disable specified services in deployment

### Synopsis

Disable specified services in deployment. Services can be repeated and execution
will happen in the order that is given. Specifing a tag will display nodes that
have that tag.

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
  -h, --help              help for disable
  -q, --quiet             display no output
  -t, --tag stringArray   filter nodes using tags (default [default])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in deployment

