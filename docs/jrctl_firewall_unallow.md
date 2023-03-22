## jrctl firewall unallow

Deletes allow entry given a source IP address and a port number

### Synopsis

Removes a 'allow' entry. Able to control a single node as well as cluster of
nodes.

```
jrctl firewall unallow [flags]
```

### Examples

```
# Stand-Alone Server
jrctl firewall unallow -a 1.1.1.1 -p 22

# Multi-Server Cluster
jrctl firewall unallow -t db -a 1.1.1.1 -p 3306
jrctl firewall unallow -t admin -a 1.1.1.1 -p 22
```

### Options

```
  -a, --address string    ip address
  -f, --file string       use text file with line separated ips
  -h, --help              help for unallow
  -p, --port int          port to unallow
      --protocol string   specify 'tcp' or 'udp', default is 'tcp' (default "tcp")
  -q, --quiet             display no input
  -t, --tag stringArray   filter nodes using tags (default [default])
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with server firewall
* [jrctl firewall unallow cloudflare](jrctl_firewall_unallow_cloudflare.md)	 - Remove allow entries for Cloudflare IP addresses

