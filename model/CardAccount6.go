package model

// Customer account information.
type CardAccount6 struct {

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

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

	// Balance of the account.
	Balance *AmountAndDirection43 `xml:"Bal,omitempty"`

	// Indicates if the balance must be displayed to the customer on the ATM.
	BalanceDisplayFlag *TrueFalseIndicator `xml:"BalDispFlg,omitempty"`

	// Indicates if this is the default account.
	DefaultAccountIndicator *TrueFalseIndicator `xml:"DfltAcctInd,omitempty"`

	// Services allowed on the account.
	AllowedServices []*ATMServiceType4Code `xml:"AllwdSvcs,omitempty"`

	// Withdrawal limits for the account.
	WithdrawalLimits *ATMTransactionAmounts4 `xml:"WdrwlLmts,omitempty"`

	// Deposit limits for the account.
	DepositLimits *ATMTransactionAmounts5 `xml:"DpstLmts,omitempty"`
}

func (c *CardAccount6) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount6) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount6) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount6) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount6) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount6) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount6) AddBalance() *AmountAndDirection43 {
	c.Balance = new(AmountAndDirection43)
	return c.Balance
}

func (c *CardAccount6) SetBalanceDisplayFlag(value string) {
	c.BalanceDisplayFlag = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount6) SetDefaultAccountIndicator(value string) {
	c.DefaultAccountIndicator = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount6) AddAllowedServices(value string) {
	c.AllowedServices = append(c.AllowedServices, (*ATMServiceType4Code)(&value))
}

func (c *CardAccount6) AddWithdrawalLimits() *ATMTransactionAmounts4 {
	c.WithdrawalLimits = new(ATMTransactionAmounts4)
	return c.WithdrawalLimits
}

func (c *CardAccount6) AddDepositLimits() *ATMTransactionAmounts5 {
	c.DepositLimits = new(ATMTransactionAmounts5)
	return c.DepositLimits
}
