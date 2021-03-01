## jrctl service restart

Restart apache, nginx, mysql, or varnish service

### Synopsis

Restart apache, nginx, mysql, or varnish service. Ask the daemon to restart a
given service. In order to successfully restart it, the daemon first validates
the respected service's configuration.

The following environmental variables can be passed in place of the 'endpoint'
and 'token' flags: JR_DAEMON_ENDPOINT, JR_DAEMON_TOKEN.

```
jrctl service restart SERVICE [flags]
```

### Examples

```
jrctl service restart apache
jrctl service restart nginx
jrctl service restart mysql
jrctl service restart varnish
```

### Options

```
  -e, --endpoint string   specify endpoint hostname (default "localhost:27482")
  -h, --help              help for restart
  -t, --token string      specify auth token
```

### SEE ALSO

* [jrctl service](jrctl_service.md)	 - Interact with system services

