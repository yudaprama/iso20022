package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount27 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus2Code `xml:"Sts"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	Type *FundCashAccount3Code `xml:"Tp,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Ownership status of the account, eg, joint owners.
	OwnershipType *AccountOwnershipType3Code `xml:"OwnrshTp"`

	// Ownership status of the account, eg, joint owners.
	ExtendedOwnershipType *Extended350Code `xml:"XtndedOwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemptionReason *TaxExemptReason1Code `xml:"TaxXmptnRsn,omitempty"`

	// Tax advantage specific to the account.
	ExtendedTaxExemptionReason *Extended350Code `xml:"XtndedTaxXmptnRsn,omitempty"`

	// Regularity at which a statement is issued.
	StatementFrequency *EventFrequency1Code `xml:"StmtFrqcy,omitempty"`

	// Regularity at which a statement is issued.
	ExtendedStatementFrequency *Extended350Code `xml:"XtndedStmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod1Code `xml:"TaxWhldgMtd,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Detailed information about the investment fund associated to the account.
	FundsDetails []*FinancialInstrument10 `xml:"FndsDtls,omitempty"`

	// Part of the investment account to or from which cash entries are made.
	CashAccount []*CashAccount12 `xml:"CshAcct,omitempty"`

	// Part of the investment account to or from which securities entries are made.
	SecuritiesAccount []*SecuritiesAccount4 `xml:"SctiesAcct,omitempty"`
}

func (i *InvestmentAccount27) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount27) SetStatus(value string) {
	i.Status = (*AccountStatus2Code)(&value)
}

func (i *InvestmentAccount27) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount27) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount27) SetType(value string) {
	i.Type = (*FundCashAccount3Code)(&value)
}

func (i *InvestmentAccount27) SetExtendedType(value string) {
	i.ExtendedType = (*Extended350Code)(&value)
}

func (i *InvestmentAccount27) SetOwnershipType(value string) {
	i.OwnershipType = (*AccountOwnershipType3Code)(&value)
}

func (i *InvestmentAccount27) SetExtendedOwnershipType(value string) {
	i.ExtendedOwnershipType = (*Extended350Code)(&value)
}

func (i *InvestmentAccount27) SetTaxExemptionReason(value string) {
	i.TaxExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (i *InvestmentAccount27) SetExtendedTaxExemptionReason(value string) {
	i.ExtendedTaxExemptionReason = (*Extended350Code)(&value)
}

func (i *InvestmentAccount27) SetStatementFrequency(value string) {
	i.StatementFrequency = (*EventFrequency1Code)(&value)
}

func (i *InvestmentAccount27) SetExtendedStatementFrequency(value string) {
	i.ExtendedStatementFrequency = (*Extended350Code)(&value)
}

func (i *InvestmentAccount27) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount27) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount27) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount27) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod1Code)(&value)
}

func (i *InvestmentAccount27) SetLetterIntentReference(value string) {
	i.LetterIntentReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount27) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount27) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount27) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount27) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount27) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount27) AddFundsDetails() *FinancialInstrument10 {
	newValue := new(FinancialInstrument10)
	i.FundsDetails = append(i.FundsDetails, newValue)
	return newValue
}

func (i *InvestmentAccount27) AddCashAccount() *CashAccount12 {
	newValue := new(CashAccount12)
	i.CashAccount = append(i.CashAccount, newValue)
	return newValue
}

func (i *InvestmentAccount27) AddSecuritiesAccount() *SecuritiesAccount4 {
	newValue := new(SecuritiesAccount4)
	i.SecuritiesAccount = append(i.SecuritiesAccount, newValue)
	return newValue
}
