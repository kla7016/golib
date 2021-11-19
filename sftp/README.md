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
for _, file := range listFile {
    fmt.Println("name", file.Name)
    fmt.Println("name", file.Size)
    fmt.Println("name", file.ModTime)
}
```


## Example OpenFile
```
client, err := sftp.NewConn("<host>", "<username>", "<password>", 22)
if err != nil {
    log.Fatalln(err)
}

// <remote_dir> ex:  '/folder/a.xlsx'
srcFile, err := sc.client.OpenFile(<remote_dir>, (os.O_RDONLY))
if err != nil {
    log.Fatalln(err)
}
defer srcFile.Close()
```