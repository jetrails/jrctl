## jrctl firewall

Interact with firewall to whitelist IP addresses/ports

### Synopsis

Interact with firewall to whitelist IP addresses/ports

### Examples

```
  jrctl firewall list
  jrctl firewall allow -a 1.1.1.1 -p 80 -p 443
  jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me
  jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me -c 'Office'
```

### Options

```
  -h, --help   help for firewall
```

### SEE ALSO

* [jrctl](jrctl.md)	 - Command line tool to help interact with the [32m>[0mjetrails[32m_[0m API
* [jrctl firewall allow](jrctl_firewall_allow.md)	 - Add entry to firewall
* [jrctl firewall list](jrctl_firewall_list.md)	 - List firewall firewall entries
