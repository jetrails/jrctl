## jrctl firewall undeny

Deletes deny entry given a source IP address and a port number

### Synopsis

Removes a 'deny' entry. Able to control a single node as well as cluster of
nodes.

```
jrctl firewall undeny [flags]
```

### Examples

```
# Stand-Alone Server
jrctl firewall undeny -a 1.1.1.1 -p 22

# Multi-Server Cluster
jrctl firewall undeny -t db -a 1.1.1.1 -p 3306
jrctl firewall undeny -t admin -a 1.1.1.1 -p 22
jrctl firewall undeny -t admin -a 1.1.1.1 -p 22,2223
```

### Options

```
  -a, --address string    ip address
  -f, --file string       use text file with line separated ips
  -h, --help              help for undeny
  -p, --port ints         port to undeny, can be specified multiple times
      --protocol string   specify 'tcp' or 'udp', default is 'tcp' (default "tcp")
  -q, --quiet             display no input
  -t, --tag stringArray   filter nodes using tags (default [default])
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with server firewall

