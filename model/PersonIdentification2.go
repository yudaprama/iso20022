package model

// Unique and unambiguous way to identify a person.
type PersonIdentification2 struct {

	// Number assigned by a license authority to a driver's license.
	DriversLicenseNumber *Max35Text `xml:"DrvrsLicNb"`

	// Number assigned by a social security agency.
	SocialSecurityNumber *Max35Text `xml:"SclSctyNb"`

	// Number assigned by a government agency to identify foreign nationals.
	AlienRegistrationNumber *Max35Text `xml:"AlnRegnNb"`

	// Number assigned by a passport authority to a passport.
	PassportNumber *Max35Text `xml:"PsptNb"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb"`

	// Number assigned by a national authority to an identity card.
	IdentityCardNumber *Max35Text `xml:"IdntyCardNb"`

	// Number assigned to an employer by a registration authority.
	EmployerIdentificationNumber *Max35Text `xml:"MplyrIdNb"`

	// Identifier issued to a person for which no specific identifier has been defined.
	OtherIdentification *GenericIdentification4 `xml:"OthrId"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (p *PersonIdentification2) SetDriversLicenseNumber(value string) {
	p.DriversLicenseNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetSocialSecurityNumber(value string) {
	p.SocialSecurityNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetAlienRegistrationNumber(value string) {
	p.AlienRegistrationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetPassportNumber(value string) {
	p.PassportNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetTaxIdentificationNumber(value string) {
	p.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetIdentityCardNumber(value string) {
	p.IdentityCardNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetEmployerIdentificationNumber(value string) {
	p.EmployerIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) AddOtherIdentification() *GenericIdentification4 {
	p.OtherIdentification = new(GenericIdentification4)
	return p.OtherIdentification
}

func (p *PersonIdentification2) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}
