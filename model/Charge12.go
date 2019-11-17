package model

// Identifies the different types of freight charges associated with goods.
type Charge12 struct {

	// Identifies whether the freight charges associated with the goods are "prepaid" or "collect".
	Type *FreightCharges1Code `xml:"Tp"`

	// Amount of money associated with a service.
	Charges []*ChargesDetails1 `xml:"Chrgs,omitempty"`
}

func (c *Charge12) SetType(value string) {
	c.Type = (*FreightCharges1Code)(&value)
}

func (c *Charge12) AddCharges() *ChargesDetails1 {
	newValue := new(ChargesDetails1)
	c.Charges = append(c.Charges, newValue)
	return newValue
}
