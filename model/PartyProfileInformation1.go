package model

// Information to support the Know Your Customer processes.
type PartyProfileInformation1 struct {

	// Indicates whether the certificate type has been obtained and verified.
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd"`

	// Identification of the person who validated the document.
	ValidatingParty *Max140Text `xml:"VldtngPty,omitempty"`

	// Identification of the person who checked the document.
	CheckingParty *Max140Text `xml:"ChckngPty,omitempty"`

	// Identification of the person who is responsible for the document.
	ResponsibleParty *Max140Text `xml:"RspnsblPty,omitempty"`

	// Identifies the type of certificate.
	CertificateType *CertificateType1Code `xml:"CertTp"`

	// Identifies the type of certificate.
	ExtendedCertificateType *Extended350Code `xml:"XtndedCertTp"`

	// Date at which the certification check has been performed.
	CheckingDate *ISODate `xml:"ChckngDt,omitempty"`

	// Specifies how frequently the check is performed.
	CheckingFrequency *EventFrequency1Code `xml:"ChckngFrqcy,omitempty"`

	// Specifies the date at which the next certification check will be performed.
	NextRevisionDate *ISODate `xml:"NxtRvsnDt,omitempty"`

	// Limits between which a person's salary is estimated.
	SalaryRange *Max35Text `xml:"SlryRg,omitempty"`

	// Indicates the main source of revenue.
	SourceOfWealth *Max140Text `xml:"SrcOfWlth,omitempty"`
}

func (p *PartyProfileInformation1) SetCertificationIndicator(value string) {
	p.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (p *PartyProfileInformation1) SetValidatingParty(value string) {
	p.ValidatingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation1) SetCheckingParty(value string) {
	p.CheckingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation1) SetResponsibleParty(value string) {
	p.ResponsibleParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation1) SetCertificateType(value string) {
	p.CertificateType = (*CertificateType1Code)(&value)
}

func (p *PartyProfileInformation1) SetExtendedCertificateType(value string) {
	p.ExtendedCertificateType = (*Extended350Code)(&value)
}

func (p *PartyProfileInformation1) SetCheckingDate(value string) {
	p.CheckingDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation1) SetCheckingFrequency(value string) {
	p.CheckingFrequency = (*EventFrequency1Code)(&value)
}

func (p *PartyProfileInformation1) SetNextRevisionDate(value string) {
	p.NextRevisionDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation1) SetSalaryRange(value string) {
	p.SalaryRange = (*Max35Text)(&value)
}

func (p *PartyProfileInformation1) SetSourceOfWealth(value string) {
	p.SourceOfWealth = (*Max140Text)(&value)
}
