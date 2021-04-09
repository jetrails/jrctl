## jrctl transfer receive

Download file from secure server

### Synopsis

Download file from secure server.

```
jrctl transfer receive [flags]
```

### Examples

```
  jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo
  jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -f
  jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o ./private/
  jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -n custom.name
  jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o ./private/ -n custom.name
```

### Options

```
  -f, --force            force download, overwrite existing file
  -h, --help             help for receive
  -n, --name string      specify file name
  -o, --out-dir string   specify download directory, default $PWD
```

### SEE ALSO

* [jrctl transfer](jrctl_transfer.md)	 - Securely transfer files from one machine to another

