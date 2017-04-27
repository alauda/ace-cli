package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

type loginOptions struct {
	server   string
	username string
	password string
}

// NewLoginCmd creates a new login command.
func NewLoginCmd(alauda client.APIClient) *cobra.Command {
	var opts loginOptions

	loginCmd := &cobra.Command{
		Use:   "login [OPTIONS] [SERVER]",
		Short: "Log onto the Alauda platform",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.server = args[0]
			}
			return doLogin(alauda, opts)
		},
	}

	loginCmd.Flags().StringVarP(&opts.username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&opts.password, "password", "p", "", "Password")

	return loginCmd
}

func doLogin(alauda client.APIClient, opts loginOptions) error {
	server, err := getServer(opts)
	if err != nil {
		return err
	}

	account, username, err := getAccountAndUsername(opts)
	if err != nil {
		return err
	}

	password, err := getPassword(opts)
	if err != nil {
		return err
	}

	fmt.Println("Logging into", server, "with user", account, "/", username, "and password", password)

	data := client.LoginOptions{
		Server:   server,
		Account:  account,
		Username: username,
		Password: password,
	}

	result, err := alauda.Login(data)
	if err != nil {
		return err
	}

	fmt.Println("Token:", result.Token)

	return nil
}

func getServer(opts loginOptions) (string, error) {
	if opts.server != "" {
		return opts.server, nil
	}

	return getServerFromConfig()
}

func getServerFromConfig() (string, error) {
	return "", errors.New("no API server specified")
}

func getAccountAndUsername(opts loginOptions) (string, string, error) {
	username, err := getUsername(opts)
	if err != nil {
		return "", "", err
	}

	return parseUsername(username)
}

func getUsername(opts loginOptions) (string, error) {
	if opts.username != "" {
		return opts.username, nil
	}

	return getUsernameFromTerminal()
}

func parseUsername(input string) (string, string, error) {
	result := strings.Split(input, "/")

	switch len(result) {
	case 1:
		return result[0], "", nil
	case 2:
		return result[0], result[1], nil
	}
	return "", "", errors.New("invalid username")
}

func getUsernameFromTerminal() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	input, _, err := reader.ReadLine()
	return string(input), err
}

func getPassword(opts loginOptions) (string, error) {
	if opts.password != "" {
		return opts.password, nil
	}

	return getPasswordFromTerminal()
}

func getPasswordFromTerminal() (string, error) {
	fmt.Print("Password: ")
	input, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	return string(input), err
}
