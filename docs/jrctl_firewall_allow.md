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
jrctl firewall allow -s nginx -a 1.1.1.1 -p 80 -p 443
jrctl firewall allow -s admin -a 1.1.1.1 -p 22 -b me
jrctl firewall allow -s mysql -a 1.1.1.1 -p 3306 -b me -c 'Office'
```

### Options

```
  -s, --service string   Specify deamon service selector
  -a, --address string   IP address to firewall
  -p, --port ints        port(s) to firewall
  -c, --comment string   add optional comment (default "none")
  -b, --blame string     specify blame entry (default "raffi")
  -h, --help             help for allow
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

