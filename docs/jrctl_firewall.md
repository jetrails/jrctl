## jrctl firewall

Interact with firewall to whitelist IP addresses/ports

### Examples

```
  jrctl firewall allow -h
  jrctl firewall list -h
```

### Options

```
  -h, --help   help for firewall
```

### SEE ALSO

* [jrctl](jrctl.md)	 - Command line tool to help interact with the >jetrails_ API.
* [jrctl firewall allow](jrctl_firewall_allow.md)	 - Permanently allows a source IP address to a specific port
* [jrctl firewall deny](jrctl_firewall_deny.md)	 - Permanently denies a source IP address to a specific port
* [jrctl firewall list](jrctl_firewall_list.md)	 - List firewall entries across configured servers
* [jrctl firewall unallow](jrctl_firewall_unallow.md)	 - Deletes allow entry given a source IP address and a port number
* [jrctl firewall undeny](jrctl_firewall_undeny.md)	 - Deletes deny entry given a source IP address and a port number

