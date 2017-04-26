package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

type loginOptions struct {
	server   string
	username string
	password string
}

var (
	opts loginOptions

	loginCmd = &cobra.Command{
		Use:   "login [OPTIONS] [SERVER]",
		Short: "Log onto the Alauda platform",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.server = args[0]
			}
			return doLogin(opts)
		},
	}
)

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&opts.username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&opts.password, "password", "p", "", "Password")
}

func doLogin(opts loginOptions) error {
	server, err := getServer(opts)
	if err != nil {
		return err
	}

	username, err := getUsername(opts)
	if err != nil {
		return err
	}

	password, err := getPassword(opts)
	if err != nil {
		return err
	}

	// TODO: Add SDK integration here.
	fmt.Println("Logging into", server, "with username", username, "and password", password)

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

func getUsername(opts loginOptions) (string, error) {
	if opts.username != "" {
		return opts.username, nil
	}

	return getUsernameFromTerminal()
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
