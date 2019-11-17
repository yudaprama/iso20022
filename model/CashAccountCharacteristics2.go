package model

// Specifies the characteristics of the cash account.
type CashAccountCharacteristics2 struct {

	// Defines the account level within an account hierarchy.
	AccountLevel *AccountLevel2Code `xml:"AcctLvl"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccount24 `xml:"CshAcct"`

	// Usage: the account servicer is the domicile agent servicing the local account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Defines a parent account to which the cash account is related to.
	ParentAccount *ParentCashAccount2 `xml:"PrntAcct,omitempty"`

	// Defines if and how charges and taxes due are paid to the financial institution.
	CompensationMethod *CompensationMethod1Code `xml:"CompstnMtd"`

	// Defines the account debited for charges and taxes due on the cash account, if different from the cash account.
	DebitAccount *AccountIdentification4Choice `xml:"DbtAcct,omitempty"`

	// Future date on which the account will be automatically debited for charges and taxes due.
	DelayedDebitDate *ISODate `xml:"DelydDbtDt,omitempty"`

	// Free form message advising the customer about the settlement of charges and taxes due.
	SettlementAdvice *Max105Text `xml:"SttlmAdvc,omitempty"`

	// Currency used to specify the account's balance currency.
	AccountBalanceCurrencyCode *ActiveOrHistoricCurrencyCode `xml:"AcctBalCcyCd"`

	// Currency used to specify the account's settlement currency
	SettlementCurrencyCode *ActiveOrHistoricCurrencyCode `xml:"SttlmCcyCd,omitempty"`

	// Currency used to specify the account's taxing host currency.
	HostCurrencyCode *ActiveOrHistoricCurrencyCode `xml:"HstCcyCd,omitempty"`

	// Describes account taxing parameters.
	Tax *AccountTax1 `xml:"Tax,omitempty"`

	// Individual to contact at the financial institution's location regarding problems of a business nature.
	AccountServicerContact *ContactDetails3 `xml:"AcctSvcrCtct"`
}

func (c *CashAccountCharacteristics2) SetAccountLevel(value string) {
	c.AccountLevel = (*AccountLevel2Code)(&value)
}

func (c *CashAccountCharacteristics2) AddCashAccount() *CashAccount24 {
	c.CashAccount = new(CashAccount24)
	return c.CashAccount
}

func (c *CashAccountCharacteristics2) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	c.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return c.AccountServicer
}

func (c *CashAccountCharacteristics2) AddParentAccount() *ParentCashAccount2 {
	c.ParentAccount = new(ParentCashAccount2)
	return c.ParentAccount
}

func (c *CashAccountCharacteristics2) SetCompensationMethod(value string) {
	c.CompensationMethod = (*CompensationMethod1Code)(&value)
}

func (c *CashAccountCharacteristics2) AddDebitAccount() *AccountIdentification4Choice {
	c.DebitAccount = new(AccountIdentification4Choice)
	return c.DebitAccount
}

func (c *CashAccountCharacteristics2) SetDelayedDebitDate(value string) {
	c.DelayedDebitDate = (*ISODate)(&value)
}

func (c *CashAccountCharacteristics2) SetSettlementAdvice(value string) {
	c.SettlementAdvice = (*Max105Text)(&value)
}

func (c *CashAccountCharacteristics2) SetAccountBalanceCurrencyCode(value string) {
	c.AccountBalanceCurrencyCode = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccountCharacteristics2) SetSettlementCurrencyCode(value string) {
	c.SettlementCurrencyCode = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccountCharacteristics2) SetHostCurrencyCode(value string) {
	c.HostCurrencyCode = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccountCharacteristics2) AddTax() *AccountTax1 {
	c.Tax = new(AccountTax1)
	return c.Tax
}

func (c *CashAccountCharacteristics2) AddAccountServicerContact() *ContactDetails3 {
	c.AccountServicerContact = new(ContactDetails3)
	return c.AccountServicerContact
}
