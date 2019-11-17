package model

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation11 struct {

	// Information about the organisation or individual person.
	Party *Party24Choice `xml:"Pty"`

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

	// Type of Foreign Account Tax Compliance Act (FATCA) form submitted by the investor.
	FATCAFormType []*FATCAForm1Choice `xml:"FATCAFormTp,omitempty"`

	// Foreign Account Tax Compliance Act (FATCA) status of the investor.
	FATCAStatus []*FATCAStatus1 `xml:"FATCASts,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation11) AddParty() *Party24Choice {
	i.Party = new(Party24Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation11) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation11) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) AddModifiedInvestorProfileValidation() *ModificationScope19 {
	newValue := new(ModificationScope19)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation11) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation11) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) AddFATCAFormType() *FATCAForm1Choice {
	newValue := new(FATCAForm1Choice)
	i.FATCAFormType = append(i.FATCAFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation11) AddFATCAStatus() *FATCAStatus1 {
	newValue := new(FATCAStatus1)
	i.FATCAStatus = append(i.FATCAStatus, newValue)
	return newValue
}
