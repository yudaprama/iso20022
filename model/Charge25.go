package model

// Identifies the different types of freight charges associated with goods.
type Charge25 struct {

	// Identifies whether the freight charges associated with the goods are "prepaid" or "collect".
	Type *FreightCharges1Code `xml:"Tp"`

	// Amount of money associated with a service.
	Charges []*ChargesDetails4 `xml:"Chrgs,omitempty"`
}

func (c *Charge25) SetType(value string) {
	c.Type = (*FreightCharges1Code)(&value)
}

func (c *Charge25) AddCharges() *ChargesDetails4 {
	newValue := new(ChargesDetails4)
	c.Charges = append(c.Charges, newValue)
	return newValue
}
