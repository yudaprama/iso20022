package model

// Agreement details for the over the counter market.
type Agreement4 struct {

	// Full details of the supporting legal agreement under which the margin call can be issued and/or governed.
	AgreementDetails *Max140Text `xml:"AgrmtDtls"`

	// Common reference to the agreement between the two counterparties.
	AgreementIdentification *Max140Text `xml:"AgrmtId,omitempty"`

	// Date on which the collateral agreement was signed.
	AgreementDate *ISODate `xml:"AgrmtDt"`

	// Denomination currency as specified in the collateral agreement.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Specifies the underlying master agreement.
	AgreementFramework *AgreementFramework1Choice `xml:"AgrmtFrmwk,omitempty"`
}

func (a *Agreement4) SetAgreementDetails(value string) {
	a.AgreementDetails = (*Max140Text)(&value)
}

func (a *Agreement4) SetAgreementIdentification(value string) {
	a.AgreementIdentification = (*Max140Text)(&value)
}

func (a *Agreement4) SetAgreementDate(value string) {
	a.AgreementDate = (*ISODate)(&value)
}

func (a *Agreement4) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *Agreement4) AddAgreementFramework() *AgreementFramework1Choice {
	a.AgreementFramework = new(AgreementFramework1Choice)
	return a.AgreementFramework
}
