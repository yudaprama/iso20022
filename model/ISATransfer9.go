package model

// Describes the type of product and the assets to be transferred.
type ISATransfer9 struct {

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
	TransferorAccount *Account5 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *Account6 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *PartyIdentification2Choice `xml:"Trfee"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *CashAccount11 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*ISATransfer8 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (i *ISATransfer9) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	i.PrimaryIndividualInvestor = new(IndividualPerson8)
	return i.PrimaryIndividualInvestor
}

func (i *ISATransfer9) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	i.SecondaryIndividualInvestor = new(IndividualPerson8)
	return i.SecondaryIndividualInvestor
}

func (i *ISATransfer9) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	i.OtherIndividualInvestor = append(i.OtherIndividualInvestor, newValue)
	return newValue
}

func (i *ISATransfer9) AddPrimaryCorporateInvestor() *Organisation4 {
	i.PrimaryCorporateInvestor = new(Organisation4)
	return i.PrimaryCorporateInvestor
}

func (i *ISATransfer9) AddSecondaryCorporateInvestor() *Organisation4 {
	i.SecondaryCorporateInvestor = new(Organisation4)
	return i.SecondaryCorporateInvestor
}

func (i *ISATransfer9) AddOtherCorporateInvestor() *Organisation4 {
	newValue := new(Organisation4)
	i.OtherCorporateInvestor = append(i.OtherCorporateInvestor, newValue)
	return newValue
}

func (i *ISATransfer9) AddTransferorAccount() *Account5 {
	i.TransferorAccount = new(Account5)
	return i.TransferorAccount
}

func (i *ISATransfer9) AddNomineeAccount() *Account6 {
	i.NomineeAccount = new(Account6)
	return i.NomineeAccount
}

func (i *ISATransfer9) AddTransferee() *PartyIdentification2Choice {
	i.Transferee = new(PartyIdentification2Choice)
	return i.Transferee
}

func (i *ISATransfer9) AddCashAccount() *CashAccount11 {
	i.CashAccount = new(CashAccount11)
	return i.CashAccount
}

func (i *ISATransfer9) AddProductTransfer() *ISATransfer8 {
	newValue := new(ISATransfer8)
	i.ProductTransfer = append(i.ProductTransfer, newValue)
	return newValue
}

func (i *ISATransfer9) AddExtension() *Extension1 {
	newValue := new(Extension1)
	i.Extension = append(i.Extension, newValue)
	return newValue
}
