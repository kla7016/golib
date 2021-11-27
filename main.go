package main

import (
	"github.com/kla7016/golib/sftp"
	"log"
)

func main() {
	client, err := sftp.NewConn("<host>", "<username>", "<password>", 22)
	if err != nil {
		log.Fatalln(err)
	}

	listFile, err := client.ListFile("/folder")
	if err != nil {
		log.Fatalln(err)
	}
	_ = listFile
}
