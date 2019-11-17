package model

// Provides information about the beneficial owner of the securities.
type PartyIdentification101 struct {

	// Party that is the beneficial owner of the specified quantity of securities.
	OwnerIdentification *PartyIdentification104Choice `xml:"OwnrId"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Holder of the security certifies, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Quantity of securities belonging to the beneficial owner specified.
	OwnedSecuritiesQuantity *FinancialInstrumentQuantity15Choice `xml:"OwndSctiesQty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType11Choice `xml:"CertfctnTp,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *RestrictedFINXMax350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (p *PartyIdentification101) AddOwnerIdentification() *PartyIdentification104Choice {
	p.OwnerIdentification = new(PartyIdentification104Choice)
	return p.OwnerIdentification
}

func (p *PartyIdentification101) AddAlternateIdentification() *AlternatePartyIdentification9 {
	newValue := new(AlternatePartyIdentification9)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

func (p *PartyIdentification101) SetDomicileCountry(value string) {
	p.DomicileCountry = (*CountryCode)(&value)
}

func (p *PartyIdentification101) AddNonDomicileCountry(value string) {
	p.NonDomicileCountry = append(p.NonDomicileCountry, (*CountryCode)(&value))
}

func (p *PartyIdentification101) AddOwnedSecuritiesQuantity() *FinancialInstrumentQuantity15Choice {
	p.OwnedSecuritiesQuantity = new(FinancialInstrumentQuantity15Choice)
	return p.OwnedSecuritiesQuantity
}

func (p *PartyIdentification101) AddCertificationType() *BeneficiaryCertificationType11Choice {
	newValue := new(BeneficiaryCertificationType11Choice)
	p.CertificationType = append(p.CertificationType, newValue)
	return newValue
}

func (p *PartyIdentification101) SetCertificationBreakdown(value string) {
	p.CertificationBreakdown = (*RestrictedFINXMax350Text)(&value)
}
