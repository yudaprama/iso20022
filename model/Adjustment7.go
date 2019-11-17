package model

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment7 struct {

	// Specifies the type of adjustment.
	Type *AdjustmentType1Choice `xml:"Tp"`

	// Specifies the monetary amount or rate of the adjustment.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`
}

func (a *Adjustment7) AddType() *AdjustmentType1Choice {
	a.Type = new(AdjustmentType1Choice)
	return a.Type
}

func (a *Adjustment7) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	a.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return a.AmountOrPercentage
}

func (a *Adjustment7) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}
