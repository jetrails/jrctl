## jrctl firewall allow

Add entry to firewall

### Synopsis

Add entry to firewall. Ask the daemon(s) to create an allow entry to their host
system's firewall. The tag flag is useful for cluster deployments and allows
executing command on daemons that are tagged a certain way.

```
jrctl firewall allow [flags]
```

### Options

```
  -t, --tag string       specify deamon tag selector, useful for cluster deployments (default "localhost")
  -a, --address string   ip address to firewall
  -p, --port ints        port(s) to firewall
  -c, --comment string   add optional comment (default "none")
  -h, --help             help for allow
```

### Examples For A Standalone Server 

```
jrctl firewall allow -a 1.1.1.1 -p 80 -p 443
jrctl firewall allow -a 1.1.1.1 -p 80 -p 443 -c "Dev Agency"
```

### Examples For Multi Server Clusters

```
jrctl firewall allow -t admin -a 1.1.1.1 -p 22
jrctl firewall allow -t mysql -a 1.1.1.1 -p 3306 -c 'Office VPN'
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

