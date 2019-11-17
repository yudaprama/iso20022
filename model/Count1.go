package model

// Specifies a sequence number or a total.
type Count1 struct {

	// Sequence or total number.
	Number *Number `xml:"Nb"`
}

func (c *Count1) SetNumber(value string) {
	c.Number = (*Number)(&value)
}
