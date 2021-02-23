## jrctl firewall list

List firewall entries

### Synopsis

List firewall entries. Ask the daemon for a list of firewall entries.

The following environmental variables can be passed in place of the 'endpoint'
and 'token' flags: JR_DAEMON_ENDPOINT, JR_DAEMON_TOKEN.

```
jrctl firewall list [flags]
```

### Examples

```
jrctl firewall list
```

### Options

```
  -e, --endpoint string   specify endpoint hostname (default "localhost:27482")
  -h, --help              help for list
  -t, --token string      specify auth token
```

### SEE ALSO

* [jrctl firewall](jrctl_firewall.md)	 - Interact with firewall to whitelist IP addresses/ports

