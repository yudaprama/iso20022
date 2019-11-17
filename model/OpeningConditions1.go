package model

// Specifies the conditions for the NDF opening.
type OpeningConditions1 struct {

	// Specifies the settlement currency of the non deliverable trade.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Specifies the valuation date for a non deliverable trade.
	ValuationDate *ISODate `xml:"ValtnDt"`

	// Specifies the rate source associated with the non deliverable trade.
	SettlementRateSource []*SettlementRateSource1 `xml:"SttlmRateSrc"`
}

func (o *OpeningConditions1) SetSettlementCurrency(value string) {
	o.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (o *OpeningConditions1) SetValuationDate(value string) {
	o.ValuationDate = (*ISODate)(&value)
}

func (o *OpeningConditions1) AddSettlementRateSource() *SettlementRateSource1 {
	newValue := new(SettlementRateSource1)
	o.SettlementRateSource = append(o.SettlementRateSource, newValue)
	return newValue
}
