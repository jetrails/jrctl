## jrctl secret delete

Delete secret without viewing contents

### Synopsis

Delete secret without viewing contents. Passing the secret identifier will make
a request to destroy the secret without displaying the secret's contents. If the
secret's URL is passed, the identifier is extracted automatically.

```
jrctl secret delete IDENTIFIER [flags]
```

### Examples

```
  jrctl secret delete 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq
```

### Options

```
  -h, --help   help for delete
```

### SEE ALSO

* [jrctl secret](jrctl_secret.md)	 - Interact with one-time secret service

