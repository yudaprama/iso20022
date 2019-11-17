package model

// Way to identify a customer account or a relationship to its account affected for the transaction.
type CardAccount1 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType2Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification30Choice `xml:"AcctIdr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`
}

func (c *CardAccount1) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount1) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount1) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount1) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount1) AddAccountIdentifier() *AccountIdentification30Choice {
	c.AccountIdentifier = new(AccountIdentification30Choice)
	return c.AccountIdentifier
}

func (c *CardAccount1) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}
