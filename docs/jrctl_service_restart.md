## jrctl service restart

Restart apache, nginx, mysql, or varnish service

### Synopsis

Restart apache, nginx, mysql, varnish, or php-fpm-* service. Valid entries for
php-fpm services would be prefixed with 'php-fpm-' and followed by a version
number. Ask the daemon(s) to restart a given service. In order to successfully
restart it, the daemon first validates the respected service's configuration.
Services can be repeated and execution will happen in the order that is given.

```
jrctl service restart SERVICE... [flags]
```

### Examples

```
jrctl service restart nginx
jrctl service restart nginx varnish
jrctl service restart nginx varnish php-fpm-7.2
jrctl service restart nginx varnish php-fpm-7.2 nginx
```

### Options

```
  -h, --help   help for restart
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with system services

