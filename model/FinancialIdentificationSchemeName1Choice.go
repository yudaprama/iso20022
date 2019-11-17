package model

// Sets of elements to identify a name of the organisation identification scheme.
type FinancialIdentificationSchemeName1Choice struct {

	// Name of the identification scheme, in a coded form as published in an external list.
	Code *ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`

	// Name of the identification scheme, in a free text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (f *FinancialIdentificationSchemeName1Choice) SetCode(value string) {
	f.Code = (*ExternalFinancialInstitutionIdentification1Code)(&value)
}

func (f *FinancialIdentificationSchemeName1Choice) SetProprietary(value string) {
	f.Proprietary = (*Max35Text)(&value)
}
