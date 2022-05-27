## jrctl service enable

Enable specified services in deployment

### Synopsis

Enable specified services in deployment. Services can be repeated and execution
will happen in the order that is given. Specifing a server type will only
display results for servers of that type.

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
  -h, --help               help for enable
  -q, --quiet              display no output
  -t, --type stringArray   filter servers using type selectors (default [localhost])
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with services in deployment

