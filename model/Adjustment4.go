package model

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment4 struct {

	// Specifies the type of adjustment applied to the amount of goods/services by means of a code.
	Type *AdjustmentType2Code `xml:"Tp"`

	// Specifies a type of adjustment not present in the code list.
	OtherAdjustmentType *Max35Text `xml:"OthrAdjstmntTp"`

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`

	// Specifies the monetary amount of the adjustment.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (a *Adjustment4) SetType(value string) {
	a.Type = (*AdjustmentType2Code)(&value)
}

func (a *Adjustment4) SetOtherAdjustmentType(value string) {
	a.OtherAdjustmentType = (*Max35Text)(&value)
}

func (a *Adjustment4) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}

func (a *Adjustment4) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}
