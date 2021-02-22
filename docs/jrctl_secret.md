## jrctl secret

Interact with our one-time secret service

### Synopsis

Interact with our one-time secret service

### Examples

```
  jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq
  jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c
  jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c -p secretpass
  jrctl secret delete 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq
  jrctl secret create
  jrctl secret create -c -a
  jrctl secret create -c -t 60
  jrctl secret create -c -p secretpass
  jrctl secret create -c -f ~/.ssh/id_rsa.pub
```

### Options

```
  -h, --help   help for secret
```

### SEE ALSO

* [jrctl](jrctl.md)	 - Command line tool to help interact with the [32m>[0mjetrails[32m_[0m API
* [jrctl secret create](jrctl_secret_create.md)	 - Create a new one-time secret
* [jrctl secret delete](jrctl_secret_delete.md)	 - Delete secret without viewing contents
* [jrctl secret read](jrctl_secret_read.md)	 - Display contents of secret

