## jrctl utility mkpass

Generate random passwords

### Synopsis

Generate random passwords. If count is greater than 1, then each password will
be new-line separated.

```
jrctl utility mkpass [flags]
```

### Examples

```
jrctl utility mkpass
```

### Options

```
  -c, --count int      number of passwords to generate (default 1)
  -h, --help           help for mkpass
  -l, --length int     length of password (default 32)
  -L, --no-lowercase   do not include lowercase chars in password
  -N, --no-numbers     do not include numbers in password
  -S, --no-symbols     do not include symbols in password
  -U, --no-uppercase   do not include uppercase chars in password
```

### SEE ALSO

* [jrctl utility](jrctl_utility.md)	 - Auxiliary command-line tools

