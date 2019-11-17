package model

// Information about an investment fund.
type Fund4 struct {

	// Name of the fund/sub fund.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Identification of the fund/sub fund with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Unique and unambiguous identifier for the fund/sub fund, assigned under a formal or proprietary identification scheme.
	Identification *OtherIdentification4 `xml:"Id,omitempty"`

	// Currency of the fund/sub fund.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	TotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of units of the fund/sub fund.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of units of the fund/sub fund.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Consolidated net cash flow expressed as a percentage of the total NAV for the fund/sub fund.
	PercentageOfFundTotalNAV *PercentageRate `xml:"PctgOfFndTtlNAV,omitempty"`
}

func (f *Fund4) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *Fund4) SetLegalEntityIdentifier(value string) {
	f.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (f *Fund4) AddIdentification() *OtherIdentification4 {
	f.Identification = new(OtherIdentification4)
	return f.Identification
}

func (f *Fund4) SetCurrency(value string) {
	f.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *Fund4) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund4) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund4) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *Fund4) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *Fund4) SetPercentageOfFundTotalNAV(value string) {
	f.PercentageOfFundTotalNAV = (*PercentageRate)(&value)
}
