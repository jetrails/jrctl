## jrctl report audit

Month-to-date security audit to ensure access is properly limited

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
  -f, --format string      specify 'table' or 'json' (default "table")
  -h, --help               help for audit
  -t, --type stringArray   filter servers using type selectors
```

### SEE ALSO

* [jrctl report](jrctl_report.md)	 - Generate reports for deployment

