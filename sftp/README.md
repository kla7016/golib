# FTP Client Lib

## Install
```
go get -u github.com/kla7016/golib/sftp
```
## Example ListFile
```
        client, err := sftp.NewConn("<host>", "<username>", "<password>", 22)
	if err != nil {
		log.Fatalln(err)
	}

        // listFile, err := client.ListFile("/var/lib")
	listFile, err := client.ListFile(".")
	if err != nil {
		log.Fatalln(err)
	}
```


## Example OpenFile
```
        client, err := sftp.NewConn("<host>", "<username>", "<password>", 22)
	if err != nil {
		log.Fatalln(err)
	}

	srcFile, err := sc.client.OpenFile(<remote_dir>, (os.O_RDONLY))
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()
```