## jrctl website switch-php

Switch php-fpm version for website

```
jrctl website switch-php WEBSITE_NAME PHP_VERSION [flags]
```

### Examples

```
jrctl website switch-php example.com php-fpm-7.4
jrctl website switch-php example.com php-fpm-7.4 -q
```

### Options

```
  -h, --help               help for switch-php
  -q, --quiet              display no output
  -t, --type stringArray   filter servers using type selectors (default [localhost])
```

### SEE ALSO

* [jrctl website](jrctl_website.md)	 - Manage websites in deployment

