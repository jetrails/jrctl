## jrctl firewall allow

Add entry to firewall

### Synopsis

Add entry to firewall. Ask the daemon to create an allow entry in the system's
firewall.

The following environmental variables can be passed in place of the 'endpoint'
and 'token' flags: JR_DAEMON_ENDPOINT, JR_DAEMON_TOKEN.

```
jrctl firewall allow [flags]
```

### Examples

```
jrctl firewall allow -a 1.1.1.1 -p 80 -p 443
jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me
jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me -c 'Office'
```

### Options

```
  -a, --address string    IP address to firewall
  -b, --blame string      specify blame entry (default "raffi")
  -c, --comment string    add optional comment (default "none")
  -e, --endpoint string   specify endpoint hostname (default "localhost:27482")
  -h, --help              help for allow
  -p, --port ints         port(s) to firewall
  -t, --token string      specify auth token
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

