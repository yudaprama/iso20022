package model

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation6 struct {

	// Information about the organisation or individual person.
	Party *Party14Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information to support Know Your Customer processes.
	InvestorProfileValidation []*PartyProfileInformation2 `xml:"InvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Specifies how information is sent to the account holder.
	InformationDistribution *InformationDistribution1Code `xml:"InfDstrbtn,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation6) AddParty() *Party14Choice {
	i.Party = new(Party14Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation6) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation6) AddInvestorProfileValidation() *PartyProfileInformation2 {
	newValue := new(PartyProfileInformation2)
	i.InvestorProfileValidation = append(i.InvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation6) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation6) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation6) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation6) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation6) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation6) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}
