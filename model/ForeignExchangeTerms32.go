package model

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms32 struct {

	// Currency from which the quoted currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the unit currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects that amount of the quoted currency that can be purchased with one unit of the unit currency , as follows:
	// 1 x CUR1 = nnn x CUR2,
	// where:
	// CUR1 is the unit currency
	// CUR2 is the quoted currency
	// nnn is the exchange rate.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a foreign exchange rate.
	QuotingInstitution *PartyIdentification113 `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms32) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms32) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms32) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms32) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms32) AddQuotingInstitution() *PartyIdentification113 {
	f.QuotingInstitution = new(PartyIdentification113)
	return f.QuotingInstitution
}