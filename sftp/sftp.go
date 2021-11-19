package sftp

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
	"time"
)

type sftpClient struct {
	host, user, password string
	port                 int
	client               *sftp.Client
}

func NewConn(host, user, password string, port int) (client *sftpClient, err error) {

	var auths []ssh.AuthMethod

	// Try to use $SSH_AUTH_SOCK which contains the path of the unix file socket that the sshd agent uses
	// for communication with other processes.
	if aconn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(aconn).Signers))
	}

	// Use password authentication if provided
	if password != "" {
		auths = append(auths, ssh.Password(password))
	}

	if port <= 0 || port > 65535 {
		return nil, fmt.Errorf("%s", "Invalid Parameters")
	}

	// Initialize client configuration
	config := ssh.ClientConfig{
		User: user,
		Auth: auths,
		// Uncomment to ignore host key check
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 10,
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	// Connect to server
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return nil, err
	}

	//defer conn.Close()

	// Create new SFTP client
	sc, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	client = &sftpClient{
		host:     host,
		user:     user,
		password: password,
		port:     port,
		client:   sc,
	}
	return client, err
}

type listFileStruct struct {
	Name    string
	ModTime time.Time
	Size    int64
}

func (sc *sftpClient) ListFile(remoteDir string) ([]listFileStruct, error) {
	result := make([]listFileStruct, 0)
	files, err := sc.client.ReadDir(remoteDir)
	if err != nil {
		return result, err
	}
	for _, f := range files {
		var name string
		var size int64
		var modTime time.Time
		name = f.Name()
		modTime = f.ModTime()
		size = f.Size()

		if f.IsDir() {
			name = name + "/"
			modTime = time.Time{}
			size = 0
		}
		// Output each file name and size in bytes
		result = append(result, listFileStruct{
			Name:    name,
			ModTime: modTime,
			Size:    size,
		})
	}
	return result, nil
}
