package github

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

type Credentials struct {
	username string
	password string
}

func (c *Credentials) Query() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username for 'https://github.com': ")
	username, _ := reader.ReadString('\n')
	c.username = strings.TrimSpace(username)

	fmt.Print(fmt.Sprintf("Password for 'https://%s@github.com': ", c.username))
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("\nPassword Error: %v\n", err)
	}
	password := string(bytePassword)
	c.password = strings.TrimSpace(password)
}
