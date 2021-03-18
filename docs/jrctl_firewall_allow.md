## jrctl firewall allow

Permanently allows a source IP address to a specific port

### Synopsis

Allows a specified IP address to bypass the local system firewall by creating an
'allow' entry into the permanent firewall config. Grants unprivileged users
ability to manipulate the firewall in a safe and controlled manner and keeps an
audit log. Able to control a single (localhost) system as well as clusters.

```
jrctl firewall allow [flags]
```

### Examples

```
# Stand-Alone Server
jrctl firewall allow -a 1.1.1.1 -p 80 -p 443

# Multi-Server Cluster
jrctl firewall allow -t db -a 1.1.1.1 -p 3306
jrctl firewall allow -t admin -a 1.1.1.1 -p 22 -c 'Office'
```

### Options

```
  -a, --address string   ip address
  -c, --comment string   add a comment to the firewall entry (optional) (default "none")
  -h, --help             help for allow
  -p, --port ints        port to allow, can be specified multiple times
  -t, --type string      specify deamon type selector, useful for cluster deployments (default "localhost")
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

