## jrctl firewall undeny

Deletes deny entry given a source IP address and a port number

### Synopsis

Denies a specified IP address to bypass the local system firewall by creating an
'deny' entry into the permanent firewall config. Grants unprivileged users
ability to manipulate the firewall in a safe and controlled manner and keeps an
audit log. Able to control a single (localhost) server as well as cluster of
servers.

```
jrctl firewall undeny [flags]
```

### Examples

```
  # Stand-Alone Server
  jrctl firewall deny -a 1.1.1.1 -p 80 -p 443
  
  # Multi-Server Cluster
  jrctl firewall deny -t db -a 1.1.1.1 -p 3306
  jrctl firewall deny -t admin -a 1.1.1.1 -p 22 -c 'Office'
```

### Options

```
  -a, --address string    ip address
  -h, --help              help for undeny
  -p, --port int          port to undeny
      --protocol string   specify 'tcp' or 'udp', default is 'tcp' (default "tcp")
  -t, --type string       specify server type, useful for cluster (default "localhost")
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

