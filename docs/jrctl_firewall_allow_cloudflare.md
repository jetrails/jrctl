## jrctl firewall allow cloudflare

Whitelist Cloudflare IP addresses

### Synopsis

Whitelist Cloudflare IP addresses.

```
jrctl firewall allow cloudflare [flags]
```

### Examples

```
jrctl firewall allow cloudflare
jrctl firewall allow cloudflare -t www
```

### Options

```
  -h, --help              help for cloudflare
  -q, --quiet             display no output
  -t, --tag stringArray   filter nodes using tags (default [localhost])
```

### SEE ALSO

* [jrctl firewall allow](jrctl_firewall_allow.md)	 - Permanently allows a source IP address to a specific port

