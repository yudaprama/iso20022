package model

// Account to or from which a cash entry is made.
type CashAccount33 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName5 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification70Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *FinancialInstitutionIdentification7Choice `xml:"AcctSvcr,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	AccountServicerBranch *BranchData `xml:"AcctSvcrBrnch,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, tax identification number. This may be an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	AccountOwnerOtherIdentification []*GenericIdentification82 `xml:"AcctOwnrOthrId,omitempty"`

	// Type of account.
	InvestmentAccountType *AccountType2Choice `xml:"InvstmtAcctTp,omitempty"`

	// Specifies if the account is for credits or debits.
	CreditDebit *CreditDebit3Code `xml:"CdtDbt,omitempty"`

	// Type of transaction for which the cash account is specified.
	SettlementInstructionReason *SettlementInstructionReason1Choice `xml:"SttlmInstrRsn,omitempty"`

	// Purpose of the cash account.
	CashAccountPurpose *CashAccountType3Choice `xml:"CshAcctPurp,omitempty"`

	// Specifies whether the account is the primary or secondary account if there are two accounts for the same purpose.
	CashAccountDesignation *AccountDesignation1Choice `xml:"CshAcctDsgnt,omitempty"`

	// Percentage of the dividend payment not to be reinvested, that is, to be paid in cash.
	DividendPercentage *PercentageBoundedRate `xml:"DvddPctg,omitempty"`
}

func (c *CashAccount33) SetSettlementCurrency(value string) {
	c.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CashAccount33) AddIdentification() *AccountIdentificationAndName5 {
	c.Identification = new(AccountIdentificationAndName5)
	return c.Identification
}

func (c *CashAccount33) AddAccountOwner() *PartyIdentification70Choice {
	c.AccountOwner = new(PartyIdentification70Choice)
	return c.AccountOwner
}

func (c *CashAccount33) AddAccountServicer() *FinancialInstitutionIdentification7Choice {
	c.AccountServicer = new(FinancialInstitutionIdentification7Choice)
	return c.AccountServicer
}

func (c *CashAccount33) AddAccountServicerBranch() *BranchData {
	c.AccountServicerBranch = new(BranchData)
	return c.AccountServicerBranch
}

func (c *CashAccount33) AddAccountOwnerOtherIdentification() *GenericIdentification82 {
	newValue := new(GenericIdentification82)
	c.AccountOwnerOtherIdentification = append(c.AccountOwnerOtherIdentification, newValue)
	return newValue
}

func (c *CashAccount33) AddInvestmentAccountType() *AccountType2Choice {
	c.InvestmentAccountType = new(AccountType2Choice)
	return c.InvestmentAccountType
}

func (c *CashAccount33) SetCreditDebit(value string) {
	c.CreditDebit = (*CreditDebit3Code)(&value)
}

func (c *CashAccount33) AddSettlementInstructionReason() *SettlementInstructionReason1Choice {
	c.SettlementInstructionReason = new(SettlementInstructionReason1Choice)
	return c.SettlementInstructionReason
}

func (c *CashAccount33) AddCashAccountPurpose() *CashAccountType3Choice {
	c.CashAccountPurpose = new(CashAccountType3Choice)
	return c.CashAccountPurpose
}

func (c *CashAccount33) AddCashAccountDesignation() *AccountDesignation1Choice {
	c.CashAccountDesignation = new(AccountDesignation1Choice)
	return c.CashAccountDesignation
}

func (c *CashAccount33) SetDividendPercentage(value string) {
	c.DividendPercentage = (*PercentageBoundedRate)(&value)
}
