package model

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment3 struct {

	// Specifies the type of adjustment applied to the amount of goods/services by means of a code.
	Type *AdjustmentType2Code `xml:"Tp"`

	// Specifies a type of adjustment not present in the code list.
	OtherAdjustmentType *Max35Text `xml:"OthrAdjstmntTp"`

	// Specifies the monetary amount of the adjustment.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Specifies the percentage rate of the adjustment.
	Rate *PercentageRate `xml:"Rate"`

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`
}

func (a *Adjustment3) SetType(value string) {
	a.Type = (*AdjustmentType2Code)(&value)
}

func (a *Adjustment3) SetOtherAdjustmentType(value string) {
	a.OtherAdjustmentType = (*Max35Text)(&value)
}

func (a *Adjustment3) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *Adjustment3) SetRate(value string) {
	a.Rate = (*PercentageRate)(&value)
}

func (a *Adjustment3) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}
