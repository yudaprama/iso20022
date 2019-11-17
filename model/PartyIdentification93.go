package model

// Provides information about the beneficial owner of the securities.
type PartyIdentification93 struct {

	// Party that is the beneficial owner of the specified quantity of securities.
	OwnerIdentification *PartyIdentification71Choice `xml:"OwnrId"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Holder of the security certifies, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Quantity of securities belonging to the beneficial owner specified.
	OwnedSecuritiesQuantity *FinancialInstrumentQuantity1Choice `xml:"OwndSctiesQty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType10Choice `xml:"CertfctnTp,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *Max350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (p *PartyIdentification93) AddOwnerIdentification() *PartyIdentification71Choice {
	p.OwnerIdentification = new(PartyIdentification71Choice)
	return p.OwnerIdentification
}

func (p *PartyIdentification93) AddAlternateIdentification() *AlternatePartyIdentification7 {
	newValue := new(AlternatePartyIdentification7)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

func (p *PartyIdentification93) SetDomicileCountry(value string) {
	p.DomicileCountry = (*CountryCode)(&value)
}

func (p *PartyIdentification93) AddNonDomicileCountry(value string) {
	p.NonDomicileCountry = append(p.NonDomicileCountry, (*CountryCode)(&value))
}

func (p *PartyIdentification93) AddOwnedSecuritiesQuantity() *FinancialInstrumentQuantity1Choice {
	p.OwnedSecuritiesQuantity = new(FinancialInstrumentQuantity1Choice)
	return p.OwnedSecuritiesQuantity
}

func (p *PartyIdentification93) AddCertificationType() *BeneficiaryCertificationType10Choice {
	newValue := new(BeneficiaryCertificationType10Choice)
	p.CertificationType = append(p.CertificationType, newValue)
	return newValue
}

func (p *PartyIdentification93) SetCertificationBreakdown(value string) {
	p.CertificationBreakdown = (*Max350Text)(&value)
}
