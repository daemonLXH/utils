package main

import (
	"flag"
	"fmt"
	"io"
	"utils/common"
)

func main() {
	var filename string

	flag.StringVar(&filename, "file", "", "-file")

	flag.Parse()

	br := common.GetReader(filename)

	count := 0
	for {
		count += 1
		_, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Println(string(line))
		//if count > 1 {
		//	break
		//}
	}
	fmt.Println(count)

}
