## jrctl report audit

Month-to-date security audit to ensure access is properly limited

### Synopsis

Month-to-date security audit to ensure access is properly limited.

```
jrctl report audit [flags]
```

### Examples

```
jrctl report audit
jrctl report audit -t www
jrctl report audit -o json
```

### Options

```
  -h, --help            help for audit
  -o, --output string   specify 'table' or 'json' (default "table")
  -t, --type string     specify server type selector
```

### SEE ALSO

* [jrctl report](jrctl_report.md)	 - Generate server reports

