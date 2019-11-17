package model

// Choice of quantity of assets to be transferred in percentage rate or units.
type Quantity12Choice struct {

	// Quantity of assets to be transferred.
	Unit *DecimalNumber `xml:"Unit,omitempty"`

	// Percentage rate of assets to be transferred.
	PercentageRate *PercentageRate `xml:"PctgRate,omitempty"`
}

func (q *Quantity12Choice) SetUnit(value string) {
	q.Unit = (*DecimalNumber)(&value)
}

func (q *Quantity12Choice) SetPercentageRate(value string) {
	q.PercentageRate = (*PercentageRate)(&value)
}
