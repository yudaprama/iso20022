package model

// Specifies the type of change to an amount.
type AmountModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Amount.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`
}

func (a *AmountModification1) SetModificationCode(value string) {
	a.ModificationCode = (*Modification1Code)(&value)
}

func (a *AmountModification1) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}
