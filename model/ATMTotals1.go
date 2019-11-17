package model

// Current totals of the ATM.
type ATMTotals1 struct {

	// Type of media inside the cassette.
	MediaType *ATMMediaType1Code `xml:"MdiaTp,omitempty"`

	// Currency of the totals.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Total balance of the ATM including reject cassette, but excluding the retract cassette.
	ATMBalance *ImpliedCurrencyAndAmount `xml:"ATMBal,omitempty"`

	// Available amount for dispense in the ATM.
	ATMCurrent *ImpliedCurrencyAndAmount `xml:"ATMCur,omitempty"`

	// Total number of units for non-valued media, including reject cassette.
	ATMBalanceNumber *Number `xml:"ATMBalNb,omitempty"`

	// Total number of units for non-valued media, excluding reject cassette.
	ATMCurrentNumber *Number `xml:"ATMCurNb,omitempty"`
}

func (a *ATMTotals1) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType1Code)(&value)
}

func (a *ATMTotals1) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTotals1) SetATMBalance(value, currency string) {
	a.ATMBalance = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTotals1) SetATMCurrent(value, currency string) {
	a.ATMCurrent = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTotals1) SetATMBalanceNumber(value string) {
	a.ATMBalanceNumber = (*Number)(&value)
}

func (a *ATMTotals1) SetATMCurrentNumber(value string) {
	a.ATMCurrentNumber = (*Number)(&value)
}
