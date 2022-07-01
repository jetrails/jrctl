## jrctl website php-switch

Switch php-fpm version for website

```
jrctl website php-switch WEBSITE_NAME PHP_VERSION [flags]
```

### Examples

```
jrctl website php-switch example.com php-fpm-7.4
jrctl website php-switch example.com php-fpm-7.4 -q
```

### Options

```
  -h, --help               help for php-switch
  -q, --quiet              display no output
  -t, --type stringArray   filter servers using type selectors (default [localhost])
```

### SEE ALSO

* [jrctl website](jrctl_website.md)	 - Manage websites in deployment
