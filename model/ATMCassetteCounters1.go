package model

// ATM cassette counter per unit value or globally.
type ATMCassetteCounters1 struct {

	// Amount of one media unit, if the media type is valued.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal,omitempty"`

	// Currency of the media, if the media type is valued and different from the currency of the requested amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Type of notes.
	ItemType *ATMNoteType2Code `xml:"ItmTp,omitempty"`

	// Counters of media inside the cassette.
	Counter []*ATMCassetteCounters2 `xml:"Cntr,omitempty"`

	// Current number of media present in the cassette.
	CurrentNumber *Number `xml:"CurNb"`

	// Current amount in the cassette.
	CurrentAmount *ImpliedCurrencyAndAmount `xml:"CurAmt,omitempty"`
}

func (a *ATMCassetteCounters1) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMCassetteCounters1) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMCassetteCounters1) SetItemType(value string) {
	a.ItemType = (*ATMNoteType2Code)(&value)
}

func (a *ATMCassetteCounters1) AddCounter() *ATMCassetteCounters2 {
	newValue := new(ATMCassetteCounters2)
	a.Counter = append(a.Counter, newValue)
	return newValue
}

func (a *ATMCassetteCounters1) SetCurrentNumber(value string) {
	a.CurrentNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters1) SetCurrentAmount(value, currency string) {
	a.CurrentAmount = NewImpliedCurrencyAndAmount(value, currency)
}
