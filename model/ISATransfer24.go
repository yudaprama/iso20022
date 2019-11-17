package model

// Describes the type of product and the assets to be transferred.
type ISATransfer24 struct {

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, for example, name, address, social security number and date of birth.
	OtherIndividualInvestor []*IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, for example, name and address.
	PrimaryCorporateInvestor *Organisation21 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, for example, name and address.
	SecondaryCorporateInvestor *Organisation21 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, for example, name and address.
	OtherCorporateInvestor []*Organisation21 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *Account19 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *Account19 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *PartyIdentification70Choice `xml:"Trfee"`

	// Identification of a related party or intermediary.
	IntermediaryInformation []*Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *CashAccount34 `xml:"CshAcct,omitempty"`

	// Details of the transfer to be cancelled.
	ProductTransferAndReference *ISATransfer25 `xml:"PdctTrfAndRef"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (i *ISATransfer24) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	i.PrimaryIndividualInvestor = new(IndividualPerson8)
	return i.PrimaryIndividualInvestor
}

func (i *ISATransfer24) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	i.SecondaryIndividualInvestor = new(IndividualPerson8)
	return i.SecondaryIndividualInvestor
}

func (i *ISATransfer24) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	i.OtherIndividualInvestor = append(i.OtherIndividualInvestor, newValue)
	return newValue
}

func (i *ISATransfer24) AddPrimaryCorporateInvestor() *Organisation21 {
	i.PrimaryCorporateInvestor = new(Organisation21)
	return i.PrimaryCorporateInvestor
}

func (i *ISATransfer24) AddSecondaryCorporateInvestor() *Organisation21 {
	i.SecondaryCorporateInvestor = new(Organisation21)
	return i.SecondaryCorporateInvestor
}

func (i *ISATransfer24) AddOtherCorporateInvestor() *Organisation21 {
	newValue := new(Organisation21)
	i.OtherCorporateInvestor = append(i.OtherCorporateInvestor, newValue)
	return newValue
}

func (i *ISATransfer24) AddTransferorAccount() *Account19 {
	i.TransferorAccount = new(Account19)
	return i.TransferorAccount
}

func (i *ISATransfer24) AddNomineeAccount() *Account19 {
	i.NomineeAccount = new(Account19)
	return i.NomineeAccount
}

func (i *ISATransfer24) AddTransferee() *PartyIdentification70Choice {
	i.Transferee = new(PartyIdentification70Choice)
	return i.Transferee
}

func (i *ISATransfer24) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *ISATransfer24) AddCashAccount() *CashAccount34 {
	i.CashAccount = new(CashAccount34)
	return i.CashAccount
}

func (i *ISATransfer24) AddProductTransferAndReference() *ISATransfer25 {
	i.ProductTransferAndReference = new(ISATransfer25)
	return i.ProductTransferAndReference
}

func (i *ISATransfer24) AddExtension() *Extension1 {
	newValue := new(Extension1)
	i.Extension = append(i.Extension, newValue)
	return newValue
}
