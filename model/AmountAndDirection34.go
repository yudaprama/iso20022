package model

// Amount of money expressed with an optional currency code and debit/credit indicator.
type AmountAndDirection34 struct {

	// Amount of money that results in an increase (positively signed) or decrease (negatively signed), with specification of the currency.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn"`
}

func (a *AmountAndDirection34) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection34) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}
