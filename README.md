# GoFreeRDP

A wrapper or adapter for freerdp version 3.x command line interface in go lang

## Usage

```go
package main

import "github.com/moatasemgamal/gofreerdp"

func main(){
  freeRDP, _ := gofreerdp.Init(gofreerdp.DisplayServer_Xorg)
  rdpConf := &gofreerdp.RDPConfig{
			Addr:     "192.168.100.100", // or with port "192.168.100.100:3389"
			Username: "User",
			Password: "Pass",
		}

  freeRDP.SetConfig(config)

  freeRDP.Run() // build the command and execute it
}
```
