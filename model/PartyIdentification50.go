package model

// Provides information about the beneficial owner of the securities.
type PartyIdentification50 struct {

	// Party that is the beneficial owner of the specified quantity of securities.
	OwnerIdentification *PartyIdentification48Choice `xml:"OwnrId"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Holder of the security certifies, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Quantity of securities belonging to the beneficial owner specified.
	OwnedSecuritiesQuantity *FinancialInstrumentQuantity1Choice `xml:"OwndSctiesQty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType2Choice `xml:"CertfctnTp,omitempty"`

	// Provides details relative to the beneficial owner not included within structured fields of this message.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`
}

func (p *PartyIdentification50) AddOwnerIdentification() *PartyIdentification48Choice {
	p.OwnerIdentification = new(PartyIdentification48Choice)
	return p.OwnerIdentification
}

func (p *PartyIdentification50) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

func (p *PartyIdentification50) SetDomicileCountry(value string) {
	p.DomicileCountry = (*CountryCode)(&value)
}

func (p *PartyIdentification50) AddNonDomicileCountry(value string) {
	p.NonDomicileCountry = append(p.NonDomicileCountry, (*CountryCode)(&value))
}

func (p *PartyIdentification50) AddOwnedSecuritiesQuantity() *FinancialInstrumentQuantity1Choice {
	p.OwnedSecuritiesQuantity = new(FinancialInstrumentQuantity1Choice)
	return p.OwnedSecuritiesQuantity
}

func (p *PartyIdentification50) AddCertificationType() *BeneficiaryCertificationType2Choice {
	newValue := new(BeneficiaryCertificationType2Choice)
	p.CertificationType = append(p.CertificationType, newValue)
	return newValue
}

func (p *PartyIdentification50) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}
