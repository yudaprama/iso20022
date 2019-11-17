package model

// Choice of quantity of assets to be transferred in percentage rate or units.
type Quantity14Choice struct {

	// Total quantity of securities to be settled.
	Unit *Unit4 `xml:"Unit"`

	// Percentage rate of assets to be settled.
	PercentageRate *PercentageRate `xml:"PctgRate"`
}

func (q *Quantity14Choice) AddUnit() *Unit4 {
	q.Unit = new(Unit4)
	return q.Unit
}

func (q *Quantity14Choice) SetPercentageRate(value string) {
	q.PercentageRate = (*PercentageRate)(&value)
}
