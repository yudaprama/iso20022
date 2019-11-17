package model

// Physical packaging of goods for transport.
type Consignment1 struct {

	// Total quantity of packaging units, eg number of boxes, containers, pallets, etc
	TotalQuantity *Quantity3 `xml:"TtlQty,omitempty"`

	// Total volume of goods shipped, eg number of cubic meters.
	TotalVolume *Quantity3 `xml:"TtlVol,omitempty"`

	// Total weight of goods shipped, eg number of kg, tons.
	TotalWeight *Quantity3 `xml:"TtlWght,omitempty"`
}

func (c *Consignment1) AddTotalQuantity() *Quantity3 {
	c.TotalQuantity = new(Quantity3)
	return c.TotalQuantity
}

func (c *Consignment1) AddTotalVolume() *Quantity3 {
	c.TotalVolume = new(Quantity3)
	return c.TotalVolume
}

func (c *Consignment1) AddTotalWeight() *Quantity3 {
	c.TotalWeight = new(Quantity3)
	return c.TotalWeight
}
