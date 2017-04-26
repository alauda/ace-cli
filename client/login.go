package client

// LoginOptions defines the input options for the login API.
type LoginOptions struct {
	Account  string
	Username string
	Password string
}

// LoginSuccess contains the token after a successful login.
type LoginSuccess struct {
	Token string
}

// Login authenticates against the Alauda server.
func (client *Client) Login(opts LoginOptions) (LoginSuccess, error) {
	return LoginSuccess{
		Token: "fake-token",
	}, nil
}
