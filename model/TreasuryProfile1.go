package model

// Treasury trading profile.
type TreasuryProfile1 struct {

	// Execution date of treasury bond trade.
	Date *ISODate `xml:"Dt"`

	// Type of party that performs trading operations, for example, investor or custodian.
	TraderType *PartyRole5Choice `xml:"TradrTp"`

	// Tax rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (t *TreasuryProfile1) SetDate(value string) {
	t.Date = (*ISODate)(&value)
}

func (t *TreasuryProfile1) AddTraderType() *PartyRole5Choice {
	t.TraderType = new(PartyRole5Choice)
	return t.TraderType
}

func (t *TreasuryProfile1) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}
