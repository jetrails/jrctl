## jrctl firewall allow

### Name

```
jrctl firewall allow -- permanently allows a source IP address to a specific local port
```
  

### Description

Allows a specified IP address to bypass the local system firewall by creating an "allow" entry into the permanent firewall config.

Grants unprivileged users ability to manipulate the firewall in a safe and controlled manner and keeps an audit log.

Able to control a single (localhost) system as well as clusters.  

```
jrctl firewall allow [flags]
```

### Options

```
  -a, --address string   IP address
  -c, --comment string   Add a comment to the firewall entry (optional)
  -h, --help             This help screen
  -p, --port int         Port to firewall, can be specified multiple times
  -t, --tag string       Specify cluster tier: mysql, www, admin.  Tags are preconfigured by local sysadmin during provisioning.
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

