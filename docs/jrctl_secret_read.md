## jrctl secret read

Display contents of secret

### Synopsis

Display contents of secret

```
jrctl secret read <identifier> [flags]
```

### Examples

```
  jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq
  jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c
  jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c -p secretpass
```

### Options

```
  -c, --clipboard         copy contents to clipboard
  -h, --help              help for read
  -p, --password string   password to access secret
```

### SEE ALSO

* [jrctl secret](jrctl_secret.md)	 - Interact with our one-time secret service

