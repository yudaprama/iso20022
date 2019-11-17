package model

// Choice between either a short, long or a proprietary identification format.
type IdentificationFormat4Choice struct {

	// Format expressed as a short identification.
	ShortIdentification *Exact3UpperCaseAlphaNumericText `xml:"ShrtId"`

	// Format expressed as a long identification.
	LongIdentification *RestrictedFINXMax30Text `xml:"LngId"`

	// Format expressed as a proprietary identification.
	ProprietaryIdentification *GenericIdentification86 `xml:"PrtryId"`
}

func (i *IdentificationFormat4Choice) SetShortIdentification(value string) {
	i.ShortIdentification = (*Exact3UpperCaseAlphaNumericText)(&value)
}

func (i *IdentificationFormat4Choice) SetLongIdentification(value string) {
	i.LongIdentification = (*RestrictedFINXMax30Text)(&value)
}

func (i *IdentificationFormat4Choice) AddProprietaryIdentification() *GenericIdentification86 {
	i.ProprietaryIdentification = new(GenericIdentification86)
	return i.ProprietaryIdentification
}
