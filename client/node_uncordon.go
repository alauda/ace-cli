package client

// UncordonNode makes the specific node schedulable again.
func (client *Client) UncordonNode(ip string, cluster string) error {
	data := nodeActionData{
		Action: "uncordon",
	}

	err := client.performActionOnNode(ip, cluster, &data)

	return err
}
