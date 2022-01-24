package ethrelay

func (c Client) Account() string {
	return c.account.Hex()
}