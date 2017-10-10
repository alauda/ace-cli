package client

// CordonNode makes the specific node unschedulable.
func (client *Client) CordonNode(ip string, cluster string) error {
	data := nodeActionData{
		Action: "cordon",
	}

	err := client.performActionOnNode(ip, cluster, &data)

	return err
}
