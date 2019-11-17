package model

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation8 struct {

	// Information about the organisation or individual person.
	Party *Party15Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Percentage of ownership or beneficiary ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the account owner signature is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope19 `xml:"ModfdInvstrPrflVldtn,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Specifies how information is sent to the account holder.
	InformationDistribution *InformationDistribution1Code `xml:"InfDstrbtn,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation8) AddParty() *Party15Choice {
	i.Party = new(Party15Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation8) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation8) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation8) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation8) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation8) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation8) AddModifiedInvestorProfileValidation() *ModificationScope19 {
	newValue := new(ModificationScope19)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation8) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation8) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}
