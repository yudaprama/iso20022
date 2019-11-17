package model

// Specifies the type of adjustment applied to the amount of goods/services by means of a code or free text.
type AdjustmentType1Choice struct {

	// Specifies the type of adjustment applied to the amount of goods/services by means of a code.
	Type *AdjustmentType2Code `xml:"Tp"`

	// Specifies a type of adjustment not present in the code list.
	OtherAdjustmentType *Max35Text `xml:"OthrAdjstmntTp"`
}

func (a *AdjustmentType1Choice) SetType(value string) {
	a.Type = (*AdjustmentType2Code)(&value)
}

func (a *AdjustmentType1Choice) SetOtherAdjustmentType(value string) {
	a.OtherAdjustmentType = (*Max35Text)(&value)
}
