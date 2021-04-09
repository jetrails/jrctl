## jrctl secret read

Display contents of secret

### Synopsis

Display contents of secret. Passing the secret identifier will allow us to
retrieve the contents of the secret and print it to STDOUT. Optionally, you can
copy the contents to your clipboard by passing the --clipboard flag! If the
secret's URL is passed, the identifier is extracted automatically.

```
jrctl secret read IDENTIFIER [flags]
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

* [jrctl secret](jrctl_secret.md)	 - Interact with one-time secret service

