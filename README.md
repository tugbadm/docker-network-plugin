# docker-network-plugin
Docker [go-plugin-helper](https://github.com/docker/go-plugins-helpers/) network implementation

[![Build Status](https://travis-ci.org/tugbadartici/docker-network-plugin.svg?branch=master)](https://travis-ci.org/tugbadartici/docker-network-plugin)

To get this repo
```
go get github.com/tugbadartici/docker-network-plugin
```

________
Docker network plugin uses a network driver file which is located on __"/etc/docker/plugins/"__. If this dir does not exist create it.

1- Create a __json__ file on this location & give an empty port 

```json
{
	"Name": "mynetwork",
	"Addr": "tcp://127.0.0.1:8010"
}
```

2- Example usage of __mydriver__ package with __golang__ is like this. You need to pass the __port__ as parameter to _Handle_ function.

```go
package main

import (
	"github.com/tugbadartici/docker-network-plugin"
	"github.com/docker/go-plugins-helpers/network"
)

func main() {
	d := mydriver.NewDriver()
	h := network.NewHandler(d)
	
	// About 3rd parameter:
	// Windows default daemon: WindowsDefaultDaemonRootDir() function. 
	// On Unix, this parameter is ignored.
	h.ServeTCP("test", ":8010", "", nil)  
}
```


##### more details please visit: https://medium.com/@tugba