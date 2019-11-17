package model

// Deposit limits for the account.
type ATMTransactionAmounts5 struct {

	// True if limits may be displayed on the ATM to the customer.
	DisplayFlag *TrueFalseIndicator `xml:"DispFlg,omitempty"`

	// Maximum amount allowed for deposit on the account.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`
}

func (a *ATMTransactionAmounts5) SetDisplayFlag(value string) {
	a.DisplayFlag = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransactionAmounts5) SetMaximumAmount(value, currency string) {
	a.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}
