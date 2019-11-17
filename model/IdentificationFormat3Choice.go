package model

// Choice between either a short, long or a proprietary identification format.
type IdentificationFormat3Choice struct {

	// Format expressed as a short identification.
	ShortIdentification *Exact3UpperCaseAlphaNumericText `xml:"ShrtId"`

	// Format expressed as a long identification.
	LongIdentification *Max30Text `xml:"LngId"`

	// Format expressed as a proprietary identification.
	ProprietaryIdentification *GenericIdentification36 `xml:"PrtryId"`
}

func (i *IdentificationFormat3Choice) SetShortIdentification(value string) {
	i.ShortIdentification = (*Exact3UpperCaseAlphaNumericText)(&value)
}

func (i *IdentificationFormat3Choice) SetLongIdentification(value string) {
	i.LongIdentification = (*Max30Text)(&value)
}

func (i *IdentificationFormat3Choice) AddProprietaryIdentification() *GenericIdentification36 {
	i.ProprietaryIdentification = new(GenericIdentification36)
	return i.ProprietaryIdentification
}
