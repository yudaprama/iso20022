package model

// Information to support the Know Your Customer processes.
type PartyProfileInformation2 struct {

	// Indicates whether the certificate type has been obtained and verified.
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd"`

	// Identification of the person who validated the document.
	ValidatingParty *Max140Text `xml:"VldtngPty,omitempty"`

	// Identification of the person who checked the document.
	CheckingParty *Max140Text `xml:"ChckngPty,omitempty"`

	// Identification of the person who is responsible for the document.
	ResponsibleParty *Max140Text `xml:"RspnsblPty,omitempty"`

	// Type of certificate.
	CertificateType *CertificationType1Choice `xml:"CertTp"`

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

	// Specifies an assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Specifies the customer’s money laundering risk.
	RiskLevel *RiskLevel1Choice `xml:"RskLvl,omitempty"`
}

func (p *PartyProfileInformation2) SetCertificationIndicator(value string) {
	p.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (p *PartyProfileInformation2) SetValidatingParty(value string) {
	p.ValidatingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation2) SetCheckingParty(value string) {
	p.CheckingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation2) SetResponsibleParty(value string) {
	p.ResponsibleParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation2) AddCertificateType() *CertificationType1Choice {
	p.CertificateType = new(CertificationType1Choice)
	return p.CertificateType
}

func (p *PartyProfileInformation2) SetCheckingDate(value string) {
	p.CheckingDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation2) SetCheckingFrequency(value string) {
	p.CheckingFrequency = (*EventFrequency1Code)(&value)
}

func (p *PartyProfileInformation2) SetNextRevisionDate(value string) {
	p.NextRevisionDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation2) SetSalaryRange(value string) {
	p.SalaryRange = (*Max35Text)(&value)
}

func (p *PartyProfileInformation2) SetSourceOfWealth(value string) {
	p.SourceOfWealth = (*Max140Text)(&value)
}

func (p *PartyProfileInformation2) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	p.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return p.CustomerConductClassification
}

func (p *PartyProfileInformation2) AddRiskLevel() *RiskLevel1Choice {
	p.RiskLevel = new(RiskLevel1Choice)
	return p.RiskLevel
}
