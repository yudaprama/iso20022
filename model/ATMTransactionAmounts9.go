package model

// Limit of deposited media for the customer.
type ATMTransactionAmounts9 struct {

	// Type of media.
	MediaType *ATMMediaType2Code `xml:"MdiaTp"`

	// Currency of the media.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Minimum number of media.
	MinimumNumber *Number `xml:"MinNb,omitempty"`

	// Maximum number of media.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// True if limits may be displayed to the customer on the ATM.
	DisplayFlag *TrueFalseIndicator `xml:"DispFlg,omitempty"`
}

func (a *ATMTransactionAmounts9) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType2Code)(&value)
}

func (a *ATMTransactionAmounts9) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts9) SetMinimumNumber(value string) {
	a.MinimumNumber = (*Number)(&value)
}

func (a *ATMTransactionAmounts9) SetMaximumNumber(value string) {
	a.MaximumNumber = (*Number)(&value)
}

func (a *ATMTransactionAmounts9) SetDisplayFlag(value string) {
	a.DisplayFlag = (*TrueFalseIndicator)(&value)
}
