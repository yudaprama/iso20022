package model

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation13 struct {

	// Information about the organisation or individual person.
	Party *Party30Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope27 `xml:"ModfdInvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions or purchases and or redemptions or sells  will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of the account may benefit from a fiscal exemption or amnesty, for example, when declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Type of information that must be provided to the account holder.
	Notification []*Notification2 `xml:"Ntfctn,omitempty"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form submitted by the investor or account owner.
	FATCAFormType []*FATCAForm1Choice `xml:"FATCAFormTp,omitempty"`

	// Foreign Account Tax Compliance Act (FATCA) status of the investor or account owner.
	FATCAStatus []*FATCAStatus2 `xml:"FATCASts,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, tax identification number. This may be an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification82 `xml:"OthrId,omitempty"`

	// Tax advantage specific to the account party.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting1 `xml:"TaxRptg,omitempty"`

	// Language in which the organisation or person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Method used for postal mailing.
	MailType *MailType1Choice `xml:"MailTp,omitempty"`

	// Country and residential status of the organisation or individual person.
	CountryAndResidentialStatus *CountryAndResidentialStatusType2 `xml:"CtryAndResdtlSts,omitempty"`

	// Annual wealth of the individual person or share capital value of the legal entity and date on which the annual wealth of the individual person was registered or declared or the date the  stock value of the organisation was registered.
	MonetaryWealth *DateAndAmount1 `xml:"MntryWlth,omitempty"`

	// Amount of total assets minus liabilities of the individual person or the amount of the difference between assets and liabilities plus rights over obligations (net equity) of the organisation and the date on which the equity value was registered.
	EquityValue *DateAndAmount1 `xml:"EqtyVal,omitempty"`

	// Resource or value owned or used by a third-party company and the date on which the working capital amount was registered.
	WorkingCapital *DateAndAmount1 `xml:"WorkgCptl,omitempty"`

	// Account owner's connection with the trading party or broker.
	CompanyLink *CompanyLink1Choice `xml:"CpnyLk,omitempty"`

	// Reference to be specified when a letter (for example, an order confirmation) is sent by an automated mailing system.
	ElectronicMailingServiceReference *Max350Text `xml:"ElctrncMlngSvcRef,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress []*CommunicationAddress6 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress []*CommunicationAddress6 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor or account owner that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if the account party is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Additional information concerning limitations and restrictions on the account party.
	AdditionalInformation []*AccountRestrictions1 `xml:"AddtlInf,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation13) AddParty() *Party30Choice {
	i.Party = new(Party30Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation13) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation13) AddModifiedInvestorProfileValidation() *ModificationScope27 {
	newValue := new(ModificationScope27)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation13) AddNotification() *Notification2 {
	newValue := new(Notification2)
	i.Notification = append(i.Notification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddFATCAFormType() *FATCAForm1Choice {
	newValue := new(FATCAForm1Choice)
	i.FATCAFormType = append(i.FATCAFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddFATCAStatus() *FATCAStatus2 {
	newValue := new(FATCAStatus2)
	i.FATCAStatus = append(i.FATCAStatus, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddOtherIdentification() *GenericIdentification82 {
	newValue := new(GenericIdentification82)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccountOwnershipInformation13) AddTaxReporting() *TaxReporting1 {
	newValue := new(TaxReporting1)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) AddMailType() *MailType1Choice {
	i.MailType = new(MailType1Choice)
	return i.MailType
}

func (i *InvestmentAccountOwnershipInformation13) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType2 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType2)
	return i.CountryAndResidentialStatus
}

func (i *InvestmentAccountOwnershipInformation13) AddMonetaryWealth() *DateAndAmount1 {
	i.MonetaryWealth = new(DateAndAmount1)
	return i.MonetaryWealth
}

func (i *InvestmentAccountOwnershipInformation13) AddEquityValue() *DateAndAmount1 {
	i.EquityValue = new(DateAndAmount1)
	return i.EquityValue
}

func (i *InvestmentAccountOwnershipInformation13) AddWorkingCapital() *DateAndAmount1 {
	i.WorkingCapital = new(DateAndAmount1)
	return i.WorkingCapital
}

func (i *InvestmentAccountOwnershipInformation13) AddCompanyLink() *CompanyLink1Choice {
	i.CompanyLink = new(CompanyLink1Choice)
	return i.CompanyLink
}

func (i *InvestmentAccountOwnershipInformation13) SetElectronicMailingServiceReference(value string) {
	i.ElectronicMailingServiceReference = (*Max350Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) AddPrimaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.PrimaryCommunicationAddress = append(i.PrimaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddSecondaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.SecondaryCommunicationAddress = append(i.SecondaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *InvestmentAccountOwnershipInformation13) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccountOwnershipInformation13) AddAdditionalInformation() *AccountRestrictions1 {
	newValue := new(AccountRestrictions1)
	i.AdditionalInformation = append(i.AdditionalInformation, newValue)
	return newValue
}
