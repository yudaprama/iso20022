package model

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation4 struct {

	// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
	Organisation *Organisation3 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson11 `xml:"IndvPrsn"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Code `xml:"MnyLndrgChck,omitempty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	ExtendedMoneyLaunderingCheck *Extended350Code `xml:"XtndedMnyLndrgChck,omitempty"`

	// Percentage of ownership or beneficiary ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the account owner signature is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope11 `xml:"ModfdInvstrPrflVldtn,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation4) AddOrganisation() *Organisation3 {
	i.Organisation = new(Organisation3)
	return i.Organisation
}

func (i *InvestmentAccountOwnershipInformation4) AddIndividualPerson() *IndividualPerson11 {
	i.IndividualPerson = new(IndividualPerson11)
	return i.IndividualPerson
}

func (i *InvestmentAccountOwnershipInformation4) SetMoneyLaunderingCheck(value string) {
	i.MoneyLaunderingCheck = (*MoneyLaunderingCheck1Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetExtendedMoneyLaunderingCheck(value string) {
	i.ExtendedMoneyLaunderingCheck = (*Extended350Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) AddModifiedInvestorProfileValidation() *ModificationScope11 {
	newValue := new(ModificationScope11)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}
