![example workflow](https://github.com/fpersson/sensormonitor/actions/workflows/go.yml/badge.svg)

# How-to.

## OpenSuSE
This is the easy way....

``` bash
    sudo zypper addrepo -f https://download.opensuse.org/repositories/home:fpersson/openSUSE_Tumbleweed/home:fpersson.repo
    sudo zypper refresh
    sudo zypper --gpg-auto-import-keys ref
    sudo zypper install demo
    sudo systemctl enable webserver.service
    sudo systemctl start webserver.service
```

## Build/install
``` bash
    go install
```

## Running 
### run with default settings
* PORT: 8080
* COFIG: $XDG_DATA_DIRS
* HTMLDIR: ./

``` bash
    $HOME/go/bin/sensormonitor
```

### Test mode
``` bash
    PORT=":1111" HTML="./" CONFIG="./testdata/" $HOME/go/bin/sensormonitor
```

Vistit localhost:1111 with your browser

### html paths (if installed from repo)
``` bash
    /usr/share/demo/templates/
```
