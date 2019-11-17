package model

// Parameters required to request a Fund Processing Passport (FPP).
type FundParameters4 struct {

	// Financial instrument for which the fund processing passport report report is requested.
	FinancialInstrumentDetails []*FinancialInstrument17 `xml:"FinInstrmDtls,omitempty"`

	// Fund management company for which the report is requested.
	FundManagementCompany []*PartyIdentification2Choice `xml:"FndMgmtCpny,omitempty"`

	// Specifies the date on or after which the information required will have been last updated. Only the most recent versions of the data is required.
	DateFrom *ISODate `xml:"DtFr,omitempty"`

	// Country where the fund has legal domicile as reflected in the ISIN classification.
	CountryOfDomicile *CountryCode `xml:"CtryOfDmcl,omitempty"`

	// Countries where the fund is registered for distribution.
	RegisteredDistributionCountry *CountryCode `xml:"RegdDstrbtnCtry,omitempty"`
}

func (f *FundParameters4) AddFinancialInstrumentDetails() *FinancialInstrument17 {
	newValue := new(FinancialInstrument17)
	f.FinancialInstrumentDetails = append(f.FinancialInstrumentDetails, newValue)
	return newValue
}

func (f *FundParameters4) AddFundManagementCompany() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	f.FundManagementCompany = append(f.FundManagementCompany, newValue)
	return newValue
}

func (f *FundParameters4) SetDateFrom(value string) {
	f.DateFrom = (*ISODate)(&value)
}

func (f *FundParameters4) SetCountryOfDomicile(value string) {
	f.CountryOfDomicile = (*CountryCode)(&value)
}

func (f *FundParameters4) SetRegisteredDistributionCountry(value string) {
	f.RegisteredDistributionCountry = (*CountryCode)(&value)
}
