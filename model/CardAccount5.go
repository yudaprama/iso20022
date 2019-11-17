package model

// Customer account information.
type CardAccount5 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType2Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`
}

func (c *CardAccount5) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount5) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount5) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount5) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount5) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount5) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount5) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}
