## jrctl firewall list

List firewall entries

### Synopsis

List firewall entries. Ask the daemon for a list of firewall entries. Specifing
the service will only return results with that service. Not specifing any
service will show everything available.

```
jrctl firewall list [flags]
```

### Examples

```
jrctl firewall list
jrctl firewall list -s admin
jrctl firewall list -s mysql
```

### Options

```
  -s, --service string   filter by service
  -h, --help             help for list
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

