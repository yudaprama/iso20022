package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount39 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s subsequent instructions.
	ReinvestmentDetails []*Reinvestment1 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	ModifiedFundDetails []*ModificationScope13 `xml:"ModfdFndDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus *Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the investment account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`
}

func (i *InvestmentAccount39) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount39) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount39) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount39) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount39) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount39) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount39) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount39) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount39) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount39) AddReinvestmentDetails() *Reinvestment1 {
	newValue := new(Reinvestment1)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount39) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount39) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount39) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount39) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount39) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount39) AddModifiedFundDetails() *ModificationScope13 {
	newValue := new(ModificationScope13)
	i.ModifiedFundDetails = append(i.ModifiedFundDetails, newValue)
	return newValue
}

func (i *InvestmentAccount39) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount39) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount39) AddBlockedStatus() *Blocked1 {
	i.BlockedStatus = new(Blocked1)
	return i.BlockedStatus
}

func (i *InvestmentAccount39) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount39) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount39) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount39) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount39) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}
