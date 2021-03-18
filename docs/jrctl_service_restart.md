## jrctl service restart

Restart apache, nginx, mysql, or varnish service

### Synopsis

Restart apache, nginx, mysql, varnish, or php-fpm* service. Valid entries for
php-fpm services would be prefixed with 'php-fpm' and followed by a version
number. Ask the server(s) to restart a given service. In order to successfully
restart it, the server first validates the respected service's configuration.
Services can be repeated and execution will happen in the order that is given.

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
  -h, --help          help for restart
  -t, --type string   specify deamon type selector, useful for cluster deployments (default "localhost")
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with system services

