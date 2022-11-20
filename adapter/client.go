package adapter

type Client struct{}

func (c *Client) InsertSquareUsbInComputer(computer Computer) {
	computer.InsertInSquarePort()
}
