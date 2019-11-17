package model

// Choice between a code or a proprietary code for  the underlying master agreement.
type AgreementFramework1Choice struct {

	// Code to specify the type of collateral agreement.
	AgreementFramework *AgreementFramework1Code `xml:"AgrmtFrmwk"`

	// Proprietary identification to specify the type of collateral agreement.
	ProprietaryIdentification *GenericIdentification30 `xml:"PrtryId"`
}

func (a *AgreementFramework1Choice) SetAgreementFramework(value string) {
	a.AgreementFramework = (*AgreementFramework1Code)(&value)
}

func (a *AgreementFramework1Choice) AddProprietaryIdentification() *GenericIdentification30 {
	a.ProprietaryIdentification = new(GenericIdentification30)
	return a.ProprietaryIdentification
}
