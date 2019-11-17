package model

// Identifies the different types of freight charges associated with goods.
type Charge13 struct {

	// Identifies whether the freight charges associated with the goods are "prepaid" or "collect".
	Type *FreightCharges1Code `xml:"Tp"`

	// Amount of money associated with a service.
	Charges []*ChargesDetails2 `xml:"Chrgs,omitempty"`
}

func (c *Charge13) SetType(value string) {
	c.Type = (*FreightCharges1Code)(&value)
}

func (c *Charge13) AddCharges() *ChargesDetails2 {
	newValue := new(ChargesDetails2)
	c.Charges = append(c.Charges, newValue)
	return newValue
}
