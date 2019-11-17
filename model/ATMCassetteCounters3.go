package model

// ATM cassette counter per unit value or globally.
type ATMCassetteCounters3 struct {

	// Amount of one media unit, if the media type is valued.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal,omitempty"`

	// Currency of the media, if the media type is valued and different from the currency of the requested amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Category of media items.
	MediaCategory *ATMMediaType3Code `xml:"MdiaCtgy,omitempty"`

	// Current number of media present in the cassette.
	CurrentNumber *Number `xml:"CurNb"`

	// Current amount in the cassette.
	CurrentAmount *ImpliedCurrencyAndAmount `xml:"CurAmt,omitempty"`

	// Counters of media inside the cassette.
	FlowTotals []*ATMCassetteCounters4 `xml:"FlowTtls,omitempty"`
}

func (a *ATMCassetteCounters3) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMCassetteCounters3) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMCassetteCounters3) SetMediaCategory(value string) {
	a.MediaCategory = (*ATMMediaType3Code)(&value)
}

func (a *ATMCassetteCounters3) SetCurrentNumber(value string) {
	a.CurrentNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters3) SetCurrentAmount(value, currency string) {
	a.CurrentAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMCassetteCounters3) AddFlowTotals() *ATMCassetteCounters4 {
	newValue := new(ATMCassetteCounters4)
	a.FlowTotals = append(a.FlowTotals, newValue)
	return newValue
}
