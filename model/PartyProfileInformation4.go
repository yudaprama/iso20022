package model

// Information to support the Know Your Customer (KYC) processes.
type PartyProfileInformation4 struct {

	// Indicates whether the certificate type has been obtained and verified.
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Identification of the person who validated the document.
	ValidatingParty *Max140Text `xml:"VldtngPty,omitempty"`

	// Identification of the person who checked the document.
	CheckingParty *Max140Text `xml:"ChckngPty,omitempty"`

	// Identification of the person who is responsible for the document.
	ResponsibleParty *Max140Text `xml:"RspnsblPty,omitempty"`

	// Type of certificate.
	CertificateType *CertificationType1Choice `xml:"CertTp,omitempty"`

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

	// Specifies the type of due diligence checks carried out on the investor. For definitions of ordinary, simple and enhanced know your customer checks, local market regulations should be consulted.
	KnowYourCustomerCheckType *KYCCheckType1Choice `xml:"KnowYourCstmrChckTp,omitempty"`

	// Specifies whether a customer has been checked in a Know Your Customer (KYC) database.
	KnowYourCustomerDatabaseCheck *DataBaseCheck1 `xml:"KnowYourCstmrDBChck,omitempty"`
}

func (p *PartyProfileInformation4) SetCertificationIndicator(value string) {
	p.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (p *PartyProfileInformation4) SetValidatingParty(value string) {
	p.ValidatingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation4) SetCheckingParty(value string) {
	p.CheckingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation4) SetResponsibleParty(value string) {
	p.ResponsibleParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation4) AddCertificateType() *CertificationType1Choice {
	p.CertificateType = new(CertificationType1Choice)
	return p.CertificateType
}

func (p *PartyProfileInformation4) SetCheckingDate(value string) {
	p.CheckingDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation4) SetCheckingFrequency(value string) {
	p.CheckingFrequency = (*EventFrequency1Code)(&value)
}

func (p *PartyProfileInformation4) SetNextRevisionDate(value string) {
	p.NextRevisionDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation4) SetSalaryRange(value string) {
	p.SalaryRange = (*Max35Text)(&value)
}

func (p *PartyProfileInformation4) SetSourceOfWealth(value string) {
	p.SourceOfWealth = (*Max140Text)(&value)
}

func (p *PartyProfileInformation4) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	p.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return p.CustomerConductClassification
}

func (p *PartyProfileInformation4) AddRiskLevel() *RiskLevel1Choice {
	p.RiskLevel = new(RiskLevel1Choice)
	return p.RiskLevel
}

func (p *PartyProfileInformation4) AddKnowYourCustomerCheckType() *KYCCheckType1Choice {
	p.KnowYourCustomerCheckType = new(KYCCheckType1Choice)
	return p.KnowYourCustomerCheckType
}

func (p *PartyProfileInformation4) AddKnowYourCustomerDatabaseCheck() *DataBaseCheck1 {
	p.KnowYourCustomerDatabaseCheck = new(DataBaseCheck1)
	return p.KnowYourCustomerDatabaseCheck
}
