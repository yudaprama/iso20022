package model

// Choice between types of reporting parameter.
type ReportParameter2Choice struct {

	// Party for which the estimated cash flow is being reported.
	Party *PartyIdentification2Choice `xml:"Pty"`

	// Country for which the estimated cash flow is being reported.
	Country *CountryCode `xml:"Ctry"`

	// Currency for which the estimated cash flow is being reported.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy"`

	// Parameter for which the estimated cash flow is being reported.
	UserDefined *DataFormat2Choice `xml:"UsrDfnd"`
}

func (r *ReportParameter2Choice) AddParty() *PartyIdentification2Choice {
	r.Party = new(PartyIdentification2Choice)
	return r.Party
}

func (r *ReportParameter2Choice) SetCountry(value string) {
	r.Country = (*CountryCode)(&value)
}

func (r *ReportParameter2Choice) SetCurrency(value string) {
	r.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *ReportParameter2Choice) AddUserDefined() *DataFormat2Choice {
	r.UserDefined = new(DataFormat2Choice)
	return r.UserDefined
}
