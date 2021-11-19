package sftp

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

func TestListFile(t *testing.T) {
	godotenv.Load(".env")
	inputIp := os.Getenv("SFTP_IP")
	inputPort, _ := strconv.Atoi(os.Getenv("SFTP_PORT"))
	inputUsername := os.Getenv("SFTP_USERNAME")
	inputPassword := os.Getenv("SFTP_PASSWORD")
	dir := "."

	client, err := NewConn(inputIp, inputUsername, inputPassword, inputPort)
	if err != nil {
		t.Errorf("Input ip: %v port: %v username: %v password: %v ErrorConnect: %v", inputIp, inputPort, inputUsername, inputPassword, err)
	}

	_, err = client.ListFile(dir)
	if err != nil {
		t.Errorf("Input ip: %v port: %v username: %v password: %v ErrorListFile: %v", inputIp, inputPort, inputUsername, inputPassword, err)
	}
}

func TestDownloadFile(t *testing.T) {
	godotenv.Load(".env")
	inputIp := os.Getenv("SFTP_IP")
	inputPort, _ := strconv.Atoi(os.Getenv("SFTP_PORT"))
	inputUsername := os.Getenv("SFTP_USERNAME")
	inputPassword := os.Getenv("SFTP_PASSWORD")
	dir := os.Getenv("DIR_FILE_READ")

	sc, err := NewConn(inputIp, inputUsername, inputPassword, inputPort)
	if err != nil {
		t.Errorf("Input ip: %v port: %v username: %v password: %v ErrorConnect: %v", inputIp, inputPort, inputUsername, inputPassword, err)
	}

	srcFile, err := sc.client.OpenFile(dir, (os.O_RDONLY))
	if err != nil {
		t.Errorf("Unable to open remote file: %v", err)
		return
	}
	defer srcFile.Close()
}

