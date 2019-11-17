package model

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment6 struct {

	// Specifies the type of adjustment.
	Type *AdjustmentType1Choice `xml:"Tp"`

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`

	// Specifies the monetary amount of the adjustment.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (a *Adjustment6) AddType() *AdjustmentType1Choice {
	a.Type = new(AdjustmentType1Choice)
	return a.Type
}

func (a *Adjustment6) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}

func (a *Adjustment6) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}
