package model

// Information about an investment fund.
type Fund3 struct {

	// Name of the fund/sub fund.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Identification of the fund/sub fund with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Unique and unambiguous identifier for the fund/sub fund, assigned under a formal or proprietary identification scheme.
	Identification *OtherIdentification4 `xml:"Id,omitempty"`

	// Currency of the fund/sub fund.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Estimated total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous total value of all the holdings, less the fund's liabilities, of the fund/sub fund
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of units of the fund/sub fund.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous total number of units of the fund/sub fund.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Estimated consolidated net cash flow expressed as a percentage of the previous total NAV for the fund/sub fund.
	EstimatedPercentageOfFundTotalNAV *PercentageRate `xml:"EstmtdPctgOfFndTtlNAV,omitempty"`
}

func (f *Fund3) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *Fund3) SetLegalEntityIdentifier(value string) {
	f.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (f *Fund3) AddIdentification() *OtherIdentification4 {
	f.Identification = new(OtherIdentification4)
	return f.Identification
}

func (f *Fund3) SetCurrency(value string) {
	f.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *Fund3) SetEstimatedTotalNAV(value, currency string) {
	f.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund3) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund3) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.EstimatedTotalUnitsNumber
}

func (f *Fund3) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *Fund3) SetEstimatedPercentageOfFundTotalNAV(value string) {
	f.EstimatedPercentageOfFundTotalNAV = (*PercentageRate)(&value)
}
