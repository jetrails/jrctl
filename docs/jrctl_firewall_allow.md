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
  -a, --address string   IP address
  -c, --comment string   Add a comment to the firewall entry (optional)
  -h, --help             This help screen
  -p, --port int         Port to firewall, can be specified multiple times
  -t, --tag string       Specify cluster tier: mysql, www, admin
```

### Examples For A Standalone Server 

```
jrctl firewall allow -a 1.1.1.1 -p 22
jrctl firewall allow -a 1.1.1.1 -p 80 -p 443 -c "Dev Agency"
```

### Examples For Multi Server Clusters

```
jrctl firewall allow -t admin -a 1.1.1.1 -p 22
jrctl firewall allow -t mysql -a 1.1.1.1 -p 3306 -c 'Office VPN'
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

