package main

import (
	"bufio"
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


	f := bufio.NewReader(os.Stdin)
	for {
		_,_=fmt.Fprintf(os.Stdout,"[32m>>>>>[33m")
		res,_,_:= f.ReadLine()
		//if err != nil {
		//	fmt.Println("")
		//	continue
		//}
       fmt.Println(string(res))
	}


}

