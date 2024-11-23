# CodexCli


```
:'######:::'#######::'########::'########:'##::::'##:
'##... ##:'##.... ##: ##.... ##: ##.....::. ##::'##::
 ##:::..:: ##:::: ##: ##:::: ##: ##::::::::. ##'##:::
 ##::::::: ##:::: ##: ##:::: ##: ######:::::. ###::::
 ##::::::: ##:::: ##: ##:::: ##: ##...:::::: ## ##:::
 ##::: ##: ##:::: ##: ##:::: ##: ##:::::::: ##:. ##::
. ######::. #######:: ########:: ########: ##:::. ##:
:......::::.......:::........:::........::..:::::..::
```

CodexCli is a tool built using Cobra and the Virus total Client library built in go.

# Usage
```Codex --help```:
```bash
Codex is a virus total cli built by dogpatch626 in Go

Usage:
  Codex [flags]
  Codex [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  info        CodexCli is a in development cli
  scan        Scan a url using the Virus total api

Flags:
  -h, --help   help for Codex

Use "Codex [command] --help" for more information about a command.
```

```Codex scan --vtkey=<VirusTotal Api Key> <Url To Scan>```


# See for more info on how this was made:
https://github.com/VirusTotal/vt-go \
https://github.com/spf13/cobra














> 📝 Note: So far the Cli just returns a barely formated response of the the returned URL object from virus total. It seems that parsing through JSON in GO is quite nasty.