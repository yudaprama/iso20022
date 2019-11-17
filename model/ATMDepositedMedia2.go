package model

// Media item that are deposited.
type ATMDepositedMedia2 struct {

	// Number of deposit media.
	Count *Number `xml:"Cnt,omitempty"`

	// Amount or denomination of one media item, if the media type is valued or entered by the customer.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal,omitempty"`

	// Currency of media items, if valued and different from base currency.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Format of the check code line.
	CodeLineFormat *CheckCodeLine1Code `xml:"CdLineFrmt,omitempty"`

	// Check code line.
	CodeLine *Max70Text `xml:"CdLine,omitempty"`

	// Check amount scanned by the check reader.
	ScannedValue *ImpliedCurrencyAndAmount `xml:"ScnndVal,omitempty"`

	// Percentage of the confidence in the check amount.
	ConfidenceLevel *Percentage `xml:"CnfdncLvl,omitempty"`
}

func (a *ATMDepositedMedia2) SetCount(value string) {
	a.Count = (*Number)(&value)
}

func (a *ATMDepositedMedia2) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMDepositedMedia2) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMDepositedMedia2) SetCodeLineFormat(value string) {
	a.CodeLineFormat = (*CheckCodeLine1Code)(&value)
}

func (a *ATMDepositedMedia2) SetCodeLine(value string) {
	a.CodeLine = (*Max70Text)(&value)
}

func (a *ATMDepositedMedia2) SetScannedValue(value, currency string) {
	a.ScannedValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMDepositedMedia2) SetConfidenceLevel(value string) {
	a.ConfidenceLevel = (*Percentage)(&value)
}
