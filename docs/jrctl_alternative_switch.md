## jrctl alternative switch

Switch current version of program

```
jrctl alternative switch PROGRAM [flags]
```

### Examples

```
jrctl alternative switch php-cli -v php-cli-8.0
jrctl alternative switch php-cli -v php-cli-8.0 -q
jrctl alternative switch php-cli -v php-cli-8.0 -t www
```

### Options

```
  -h, --help              help for switch
  -q, --quiet             display no output
  -t, --tag stringArray   filter nodes using tags (default [default])
  -v, --version string    version to switch to
```

### SEE ALSO

* [jrctl alternative](jrctl_alternative.md)	 - Manage alternative programs

