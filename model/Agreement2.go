package model

// Agreement details for the over the counter market.
type Agreement2 struct {

	// Full details of the supporting legal agreement under which the margin call can be issued and/or governed.
	AgreementDetails *Max140Text `xml:"AgrmtDtls"`

	// Common reference to the agreement between the two counterparties.
	AgreementIdentification *Max140Text `xml:"AgrmtId,omitempty"`

	// Date on which the collateral agreement was signed.
	AgreementDate *ISODate `xml:"AgrmtDt"`

	// Denomination currency as specified in the collateral agreement.
	BaseCurrency *CurrencyCode `xml:"BaseCcy"`

	// Specifies the underlying master agreement.
	AgreementFramework *AgreementFramework1Choice `xml:"AgrmtFrmwk,omitempty"`
}

func (a *Agreement2) SetAgreementDetails(value string) {
	a.AgreementDetails = (*Max140Text)(&value)
}

func (a *Agreement2) SetAgreementIdentification(value string) {
	a.AgreementIdentification = (*Max140Text)(&value)
}

func (a *Agreement2) SetAgreementDate(value string) {
	a.AgreementDate = (*ISODate)(&value)
}

func (a *Agreement2) SetBaseCurrency(value string) {
	a.BaseCurrency = (*CurrencyCode)(&value)
}

func (a *Agreement2) AddAgreementFramework() *AgreementFramework1Choice {
	a.AgreementFramework = new(AgreementFramework1Choice)
	return a.AgreementFramework
}
