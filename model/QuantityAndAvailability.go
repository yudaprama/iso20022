package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type QuantityAndAvailability struct {

	// Quantity of securities in the sub-balance.
	Quantity *FinancialInstrumentQuantityChoice `xml:"Qty"`

	// Indicates whether the quantity of securities on the sub-balance is available.
	AvailabilityIndicator *YesNoIndicator `xml:"AvlbtyInd"`
}

func (q *QuantityAndAvailability) AddQuantity() *FinancialInstrumentQuantityChoice {
	q.Quantity = new(FinancialInstrumentQuantityChoice)
	return q.Quantity
}

func (q *QuantityAndAvailability) SetAvailabilityIndicator(value string) {
	q.AvailabilityIndicator = (*YesNoIndicator)(&value)
}
