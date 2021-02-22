## jrctl secret create

Create a new one-time secret

### Synopsis

Create a new one-time secret

```
jrctl secret create [flags]
```

### Examples

```
  jrctl secret create
  jrctl secret create -c -a
  jrctl secret create -c -t 60
  jrctl secret create -c -p secretpass
  jrctl secret create -c -f ~/.ssh/id_rsa.pub
```

### Options

```
  -a, --auto-generate     automatically generate password
  -c, --clipboard         copy secret url to clipboard
  -f, --file string       use file contents as secret data
  -h, --help              help for create
  -p, --password string   protect secret with a password
  -t, --ttl int           specify custom ttl in seconds (default 86400)
```

### SEE ALSO

* [jrctl secret](jrctl_secret.md)	 - Interact with our one-time secret service
