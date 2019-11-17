package model

// Identifies the different types of freight charges associated with goods.
type Charge24 struct {

	// Identifies whether the freight charges associated with the goods are "prepaid" or "collect".
	Type *FreightCharges1Code `xml:"Tp"`

	// Amount of money associated with a service.
	Charges []*ChargesDetails3 `xml:"Chrgs,omitempty"`
}

func (c *Charge24) SetType(value string) {
	c.Type = (*FreightCharges1Code)(&value)
}

func (c *Charge24) AddCharges() *ChargesDetails3 {
	newValue := new(ChargesDetails3)
	c.Charges = append(c.Charges, newValue)
	return newValue
}
