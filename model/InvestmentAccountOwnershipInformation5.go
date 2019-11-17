package model

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation5 struct {

	// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
	Organisation *Organisation2 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson10 `xml:"IndvPrsn"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Code `xml:"MnyLndrgChck,omitempty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	ExtendedMoneyLaunderingCheck *Extended350Code `xml:"XtndedMnyLndrgChck,omitempty"`

	// Information to support Know Your Customer processes.
	InvestorProfileValidation []*PartyProfileInformation1 `xml:"InvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation5) AddOrganisation() *Organisation2 {
	i.Organisation = new(Organisation2)
	return i.Organisation
}

func (i *InvestmentAccountOwnershipInformation5) AddIndividualPerson() *IndividualPerson10 {
	i.IndividualPerson = new(IndividualPerson10)
	return i.IndividualPerson
}

func (i *InvestmentAccountOwnershipInformation5) SetMoneyLaunderingCheck(value string) {
	i.MoneyLaunderingCheck = (*MoneyLaunderingCheck1Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) SetExtendedMoneyLaunderingCheck(value string) {
	i.ExtendedMoneyLaunderingCheck = (*Extended350Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) AddInvestorProfileValidation() *PartyProfileInformation1 {
	newValue := new(PartyProfileInformation1)
	i.InvestorProfileValidation = append(i.InvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation5) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}
