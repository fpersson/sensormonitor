# How-to.

## Build/install
``` bash
    go install
```

## run
``` bash
    $HOME/go/bin/webserver
```

### run with port settings
``` bash
    PORT=":1111" $HOME/go/bin/webserver
```

### Configfiles
Default path for config files is $XDG_DATA_DIRS/tempsensor

#### Run with config
``` bash
    CONFIG=$HOME/<path>/ $HOME/go/bin/webserver
```

#### Test mode
``` bash
    PORT=":1111" CONFIG="./testdata/" $HOME/go/bin/webserver
```

### html paths (if installed)
``` bash
    /usr/share/demo/templates/
```

#### Test with html
``` bash
    HTML="./" CONFIG="./testdata/" $HOME/go/bin/webserver
```

Vistit localhost:1111 with your browser
