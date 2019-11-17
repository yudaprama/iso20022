package model

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque4 struct {

	// Party to which a cheque is made payable.
	PayeeIdentification *NameAndAddress5 `xml:"PyeeId"`
}

func (c *Cheque4) AddPayeeIdentification() *NameAndAddress5 {
	c.PayeeIdentification = new(NameAndAddress5)
	return c.PayeeIdentification
}
