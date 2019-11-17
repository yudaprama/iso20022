package model

// Specifies the type, date and version of the agreement.
type AgreementConditions1 struct {

	// Specifies the type of agreement
	AgreementCode *Max6AlphaText `xml:"AgrmtCd"`

	// Specifies the date of the agreement.
	Date *ISODate `xml:"Dt,omitempty"`

	// Specifies the version of the agreement.
	Version *Exact4NumericText `xml:"Vrsn,omitempty"`
}

func (a *AgreementConditions1) SetAgreementCode(value string) {
	a.AgreementCode = (*Max6AlphaText)(&value)
}

func (a *AgreementConditions1) SetDate(value string) {
	a.Date = (*ISODate)(&value)
}

func (a *AgreementConditions1) SetVersion(value string) {
	a.Version = (*Exact4NumericText)(&value)
}
