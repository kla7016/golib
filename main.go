package main

import (
	"fmt"
	"github.com/kla7016/golib/sftp"
	"log"
)

func main() {
	client, err := sftp.NewConn("203.151.254.199", "sftpuser", "SfTp@Passw0rd#123", 22)
	if err != nil {
		log.Fatalln(err)
	}

	listFile, err := client.ListFile("/sftpuser/moderna-data")
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range listFile {
		fmt.Println("name", file.Name, "size", file.Size, "mod_time", file.ModTime)
	}
}
