## jrctl server restart

Restart specified service(s) running on configured server(s)

### Synopsis

Restart specified service(s) running on configured server(s). In order to
successfully restart a service, the server first validates the respected
service's config file. Once deemed valid, the service is restarted. For a list
of available running services, run 'jrctl server list'. Services can be repeated
and execution will happen in the order that is given. Specifing a server type
will only display results for servers of that type.

```
jrctl server restart SERVICE... [flags]
```

### Examples

```
jrctl server restart nginx
jrctl server restart nginx varnish
jrctl server restart nginx varnish php-fpm
jrctl server restart nginx varnish php-fpm-7.2 nginx
```

### Options

```
  -h, --help          help for restart
  -q, --quiet         output as little information as possible
  -t, --type string   specify server type, useful for cluster (default "localhost")
```

### SEE ALSO

* [jrctl server](jrctl_server.md)	 - Interact with servers in configured deployment

