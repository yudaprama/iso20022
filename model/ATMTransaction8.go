package model

// Preferred withdrawal transaction chosen by the the customer.
type ATMTransaction8 struct {

	// Amount to dispense.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt,omitempty"`

	// Currency.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// True if a receipt has to be printed by the ATM for the customer.
	ReceiptFlag *TrueFalseIndicator `xml:"RctFlg,omitempty"`

	// True if a balance has to be printed by the ATM on the customer receipt.
	BalancePrintFlag *TrueFalseIndicator `xml:"BalPrtFlg,omitempty"`

	// Media mix algorithm, the identification of the algorithm is an agreement between the ATM and the ATM manager.
	MixType *Max35Text `xml:"MixTp,omitempty"`

	// Media mix to select.
	Mix []*ATMMediaMix2 `xml:"Mix,omitempty"`
}

func (a *ATMTransaction8) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction8) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransaction8) SetReceiptFlag(value string) {
	a.ReceiptFlag = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction8) SetBalancePrintFlag(value string) {
	a.BalancePrintFlag = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction8) SetMixType(value string) {
	a.MixType = (*Max35Text)(&value)
}

func (a *ATMTransaction8) AddMix() *ATMMediaMix2 {
	newValue := new(ATMMediaMix2)
	a.Mix = append(a.Mix, newValue)
	return newValue
}
