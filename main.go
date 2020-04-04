package main

import (
	"flag"
	"fmt"
	"searcher/common"
)

func main() {

	config := common.ClientConfig{}
	config.Addr = *(flag.String("a","127.0.0.1","server ip address"))
	config.Port = *(flag.Int("p",6789,"server  remote port"))
	config.Password = *(flag.String("P","127.0.0.1","server ip address"))

    fmt.Println(config.Password,config.Port,config.Addr)

}

