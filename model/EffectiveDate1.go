package model

// Date and date parameters.
type EffectiveDate1 struct {

	// Date on which the SSI is effective. If the SSI is effective on a future date, then the date must be provided.
	EffectiveDate *ISODate `xml:"FctvDt"`

	// Specifies how the SSI update effective date is to be applied.
	EffectiveDateParameter *ExternalEffectiveDateParameter1Code `xml:"FctvDtParam,omitempty"`
}

func (e *EffectiveDate1) SetEffectiveDate(value string) {
	e.EffectiveDate = (*ISODate)(&value)
}

func (e *EffectiveDate1) SetEffectiveDateParameter(value string) {
	e.EffectiveDateParameter = (*ExternalEffectiveDateParameter1Code)(&value)
}
