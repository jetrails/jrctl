## jrctl firewall unallow cloudflare

Remove allow entries for Cloudflare IP addresses

### Synopsis

Remove allow entries for Cloudflare IP addresses.

```
jrctl firewall unallow cloudflare [flags]
```

### Examples

```
jrctl firewall unallow cloudflare
jrctl firewall unallow cloudflare -t www
```

### Options

```
  -h, --help               help for cloudflare
  -q, --quiet              display no output
  -t, --type stringArray   filter servers using type selectors (default [localhost])
```

### SEE ALSO

* [jrctl firewall unallow](jrctl_firewall_unallow.md)	 - Deletes allow entry given a source IP address and a port number

