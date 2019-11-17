package model

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment5 struct {

	// Specifies whether the adjustment must be substracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`

	// Specifies the monetary amount of the adjustment.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (a *Adjustment5) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}

func (a *Adjustment5) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}
