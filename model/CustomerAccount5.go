package model

// Account owned by a customer.
type CustomerAccount5 struct {

	// Identification of the account.
	Identification []*AccountIdentification4Choice `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus3Code `xml:"Sts,omitempty"`

	// Type of the account.
	Type *CashAccountType2Choice `xml:"Tp,omitempty"`

	// Medium of exchange of value.
	Currency *ActiveCurrencyCode `xml:"Ccy"`

	// Monthly average of the payment amounts (that is, payments going out) over a year.
	MonthlyPaymentValue *ImpliedCurrencyAndAmount `xml:"MnthlyPmtVal,omitempty"`

	// Monthly average of the received amounts over a year (that is, payments coming in).
	MonthlyReceivedValue *ImpliedCurrencyAndAmount `xml:"MnthlyRcvdVal,omitempty"`

	// Monthly average of the number of payments (coming in and going out) over a year.
	MonthlyTransactionNumber *Max5NumericText `xml:"MnthlyTxNb,omitempty"`

	// Sum of the end of day balances over a month divided by the number of business days in the month.
	AverageBalance *ImpliedCurrencyAndAmount `xml:"AvrgBal,omitempty"`

	// Specifies the purpose of the account.
	AccountPurpose *Max140Text `xml:"AcctPurp,omitempty"`

	// Specifies the value of the balance under which a notification will be sent to the account owner.
	FloorNotificationAmount *ImpliedCurrencyAndAmount `xml:"FlrNtfctnAmt,omitempty"`

	// Specifies the value of the balance above which a notification will be sent to the account owner.
	CeilingNotificationAmount *ImpliedCurrencyAndAmount `xml:"ClngNtfctnAmt,omitempty"`

	// Specifies how often statements (for audit purposes) will be sent, in which format, to which address.
	StatementFrequencyAndFormat []*StatementFrequencyAndForm1 `xml:"StmtFrqcyAndFrmt,omitempty"`

	// Date when the account will be or was closed.
	ClosingDate *ISODate `xml:"ClsgDt,omitempty"`

	// Restriction on capability or operations allowed.
	Restriction []*Restriction1 `xml:"Rstrctn,omitempty"`
}

func (c *CustomerAccount5) AddIdentification() *AccountIdentification4Choice {
	newValue := new(AccountIdentification4Choice)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *CustomerAccount5) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

func (c *CustomerAccount5) SetStatus(value string) {
	c.Status = (*AccountStatus3Code)(&value)
}

func (c *CustomerAccount5) AddType() *CashAccountType2Choice {
	c.Type = new(CashAccountType2Choice)
	return c.Type
}

func (c *CustomerAccount5) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CustomerAccount5) SetMonthlyPaymentValue(value, currency string) {
	c.MonthlyPaymentValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CustomerAccount5) SetMonthlyReceivedValue(value, currency string) {
	c.MonthlyReceivedValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CustomerAccount5) SetMonthlyTransactionNumber(value string) {
	c.MonthlyTransactionNumber = (*Max5NumericText)(&value)
}

func (c *CustomerAccount5) SetAverageBalance(value, currency string) {
	c.AverageBalance = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CustomerAccount5) SetAccountPurpose(value string) {
	c.AccountPurpose = (*Max140Text)(&value)
}

func (c *CustomerAccount5) SetFloorNotificationAmount(value, currency string) {
	c.FloorNotificationAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CustomerAccount5) SetCeilingNotificationAmount(value, currency string) {
	c.CeilingNotificationAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CustomerAccount5) AddStatementFrequencyAndFormat() *StatementFrequencyAndForm1 {
	newValue := new(StatementFrequencyAndForm1)
	c.StatementFrequencyAndFormat = append(c.StatementFrequencyAndFormat, newValue)
	return newValue
}

func (c *CustomerAccount5) SetClosingDate(value string) {
	c.ClosingDate = (*ISODate)(&value)
}

func (c *CustomerAccount5) AddRestriction() *Restriction1 {
	newValue := new(Restriction1)
	c.Restriction = append(c.Restriction, newValue)
	return newValue
}
