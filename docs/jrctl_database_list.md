## jrctl database list

Display databases in deployment

```
jrctl database list [flags]
```

### Examples

```
jrctl database list
jrctl database list -q
jrctl database list -t db
```

### Options

```
  -h, --help               help for list
  -q, --quiet              only display database names
  -t, --type stringArray   filter servers using type selectors (default [localhost])
```

### SEE ALSO

* [jrctl database](jrctl_database.md)	 - Manage databases in deployment

