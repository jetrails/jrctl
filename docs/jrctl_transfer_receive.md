## jrctl transfer receive

Download file from secure server

### Synopsis

Download file from secure server. If no output path is specified, then the file
is stored in the current directory and will be named after the file identifier.

The following environmental variables can be used: JR_PUBLIC_API_ENDPOINT.

```
jrctl transfer receive [flags]
```

### Examples

```
jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo
jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o private.png
```

### Options

```
  -h, --help            help for receive
  -o, --output string   specify output file path
```

### SEE ALSO

* [jrctl transfer](jrctl_transfer.md)	 - Securely transfer files from one machine to another

