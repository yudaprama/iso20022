package model

// Account involved in the card transaction.
type CardAccount2 struct {

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType2Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification30Choice `xml:"AcctIdr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	//  Balance of the account.
	Balance *AmountAndDirection41 `xml:"Bal,omitempty"`

	//  Date of the balance.
	BalanceDate *ISODate `xml:"BalDt,omitempty"`
}

func (c *CardAccount2) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount2) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount2) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount2) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount2) AddAccountIdentifier() *AccountIdentification30Choice {
	c.AccountIdentifier = new(AccountIdentification30Choice)
	return c.AccountIdentifier
}

func (c *CardAccount2) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount2) AddBalance() *AmountAndDirection41 {
	c.Balance = new(AmountAndDirection41)
	return c.Balance
}

func (c *CardAccount2) SetBalanceDate(value string) {
	c.BalanceDate = (*ISODate)(&value)
}
