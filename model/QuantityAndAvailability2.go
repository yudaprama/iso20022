package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type QuantityAndAvailability2 struct {

	// Quantity of securities in the sub-balance.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty"`

	// Indicates whether the quantity of securities on the sub-balance is available.
	AvailabilityIndicator *YesNoIndicator `xml:"AvlbtyInd"`
}

func (q *QuantityAndAvailability2) AddQuantity() *FinancialInstrumentQuantity15Choice {
	q.Quantity = new(FinancialInstrumentQuantity15Choice)
	return q.Quantity
}

func (q *QuantityAndAvailability2) SetAvailabilityIndicator(value string) {
	q.AvailabilityIndicator = (*YesNoIndicator)(&value)
}
