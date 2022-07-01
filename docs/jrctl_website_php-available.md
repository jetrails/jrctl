## jrctl website php-available

List available php-fpm versions that are available for websites to use

```
jrctl website php-available [flags]
```

### Examples

```
jrctl website php-available
jrctl website php-available -q
```

### Options

```
  -h, --help               help for php-available
  -q, --quiet              display only available php-fpm versions
  -t, --type stringArray   filter servers using type selectors (default [localhost])
```

### SEE ALSO

* [jrctl website](jrctl_website.md)	 - Manage websites in deployment
