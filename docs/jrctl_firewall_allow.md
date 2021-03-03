## jrctl firewall allow

Add entry to firewall

### Synopsis

Add entry to firewall. Ask the daemon(s) to create an allow entry to their host
system's firewall.

```
jrctl firewall allow [flags]
```

### Examples

```
jrctl firewall allow -t nginx -a 1.1.1.1 -p 80 -p 443
jrctl firewall allow -t admin -a 1.1.1.1 -p 22 -b me
jrctl firewall allow -t mysql -a 1.1.1.1 -p 3306 -b me -c 'Office'
```

### Options

```
  -t, --tag string       Specify deamon tag selector
  -a, --address string   IP address to firewall
  -p, --port ints        port(s) to firewall
  -c, --comment string   add optional comment (default "none")
  -b, --blame string     specify blame entry (default "raffi")
  -h, --help             help for allow
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

