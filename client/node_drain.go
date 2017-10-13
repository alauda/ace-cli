package client

// DrainNode migrates containers off a node and makes it unschedulable.
func (client *Client) DrainNode(ip string, cluster string) error {
	data := nodeActionData{
		Action: "drain",
	}

	err := client.performActionOnNode(ip, cluster, &data)

	return err
}
