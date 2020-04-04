package main

import (
	"flag"
	"fmt"
	"os"
	"searcher/common"
)

func main() {

	config := common.ClientConfig{}
	config.Addr = flag.String("a","127.0.0.1","server ip address")
	config.Port = flag.Int("p",6789,"server  remote port")
	config.Password = flag.String("P","","auth password")

	flag.Parse()

	var a,b int
	_,err := fmt.Fscanf(os.Stdin,"> %d %d",&a,&b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(a,b)
}

