package model

// Information about a transfer instruction.
type PEPISATransfer7 struct {

	// Information identifying the primary individual investor, eg, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, eg, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, eg, name, address, social security number and date of birth.
	OtherIndividualInvestor []*IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, eg, name and address.
	PrimaryCorporateInvestor *Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, eg, name and address.
	SecondaryCorporateInvestor *Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, eg, name and address.
	OtherCorporateInvestor []*Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	ClientAccount *Account5 `xml:"ClntAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *Account6 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	NewPlanManager *PartyIdentification2Choice `xml:"NewPlanMgr"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *CashAccount11 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*PEPISATransfer8 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PEPISATransfer7) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PEPISATransfer7) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PEPISATransfer7) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PEPISATransfer7) AddPrimaryCorporateInvestor() *Organisation4 {
	p.PrimaryCorporateInvestor = new(Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PEPISATransfer7) AddSecondaryCorporateInvestor() *Organisation4 {
	p.SecondaryCorporateInvestor = new(Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PEPISATransfer7) AddOtherCorporateInvestor() *Organisation4 {
	newValue := new(Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PEPISATransfer7) AddClientAccount() *Account5 {
	p.ClientAccount = new(Account5)
	return p.ClientAccount
}

func (p *PEPISATransfer7) AddNomineeAccount() *Account6 {
	p.NomineeAccount = new(Account6)
	return p.NomineeAccount
}

func (p *PEPISATransfer7) AddNewPlanManager() *PartyIdentification2Choice {
	p.NewPlanManager = new(PartyIdentification2Choice)
	return p.NewPlanManager
}

func (p *PEPISATransfer7) AddCashAccount() *CashAccount11 {
	p.CashAccount = new(CashAccount11)
	return p.CashAccount
}

func (p *PEPISATransfer7) AddProductTransfer() *PEPISATransfer8 {
	newValue := new(PEPISATransfer8)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PEPISATransfer7) AddExtension() *Extension1 {
	newValue := new(Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
