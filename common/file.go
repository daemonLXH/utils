package common

import (
	"bufio"
	"log"
	"os"
)

func GetReader(filename string) (br *bufio.Reader) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("打开文件失败, %v \n", err)
	}

	defer func() {
		err := file.Close()
		log.Println(err)
	}()

	br = bufio.NewReader(file)
	return
}
