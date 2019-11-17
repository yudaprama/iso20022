package model

// Physical packaging of goods for transport.
type Consignment3 struct {

	// Total quantity of packaging units, eg number of boxes, containers, pallets, etc
	TotalQuantity *Quantity10 `xml:"TtlQty,omitempty"`

	// Total volume of goods shipped, eg number of cubic meters.
	TotalVolume *Quantity10 `xml:"TtlVol,omitempty"`

	// Total weight of goods shipped, eg number of kg, tons.
	TotalWeight *Quantity10 `xml:"TtlWght,omitempty"`
}

func (c *Consignment3) AddTotalQuantity() *Quantity10 {
	c.TotalQuantity = new(Quantity10)
	return c.TotalQuantity
}

func (c *Consignment3) AddTotalVolume() *Quantity10 {
	c.TotalVolume = new(Quantity10)
	return c.TotalVolume
}

func (c *Consignment3) AddTotalWeight() *Quantity10 {
	c.TotalWeight = new(Quantity10)
	return c.TotalWeight
}
