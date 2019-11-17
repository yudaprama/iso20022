package model

// Account owned by a customer.
type CustomerAccountModification1 struct {

	// Identification of the account.
	Identification []*AccountIdentification4Choice `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *NameModification1 `xml:"Nm,omitempty"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatusModification1 `xml:"Sts,omitempty"`

	// Type of the account.
	Type *TypeModification1 `xml:"Tp,omitempty"`

	// Medium of exchange of value.
	Currency *ActiveCurrencyCode `xml:"Ccy"`

	// Monthly average of the payment amounts (that is, payments going out) over a year.
	MonthlyPaymentValue *AmountModification1 `xml:"MnthlyPmtVal,omitempty"`

	// Monthly average of the received amounts over a year (that is, payments coming in).
	MonthlyReceivedValue *AmountModification1 `xml:"MnthlyRcvdVal,omitempty"`

	// Monthly average of the number of payments (coming in and going out) over a year.
	MonthlyTransactionNumber *NumberModification1 `xml:"MnthlyTxNb,omitempty"`

	// Sum of the end of day balances over a month divided by the number of business days in the month.
	AverageBalance *AmountModification1 `xml:"AvrgBal,omitempty"`

	// Specifies the purpose of the account.
	AccountPurpose *PurposeModification1 `xml:"AcctPurp,omitempty"`

	// Specifies the value of the balance under which a notification will be sent to the account owner.
	FloorNotificationAmount *AmountModification1 `xml:"FlrNtfctnAmt,omitempty"`

	// Specifies the value of the balance above which a notification will be sent to the account owner.
	CeilingNotificationAmount *AmountModification1 `xml:"ClngNtfctnAmt,omitempty"`

	// Specifies how often statements (for audit purposes) will be sent, in which format, to which address.
	StatementFrequencyAndFormat []*StatementFrequencyAndFormModification1 `xml:"StmtFrqcyAndFrmt,omitempty"`

	// Date when the account will be or was closed.
	ClosingDate *DateModification1 `xml:"ClsgDt,omitempty"`

	// Restriction on capability or operations allowed.
	Restriction []*RestrictionModification1 `xml:"Rstrctn,omitempty"`
}

func (c *CustomerAccountModification1) AddIdentification() *AccountIdentification4Choice {
	newValue := new(AccountIdentification4Choice)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *CustomerAccountModification1) AddName() *NameModification1 {
	c.Name = new(NameModification1)
	return c.Name
}

func (c *CustomerAccountModification1) AddStatus() *AccountStatusModification1 {
	c.Status = new(AccountStatusModification1)
	return c.Status
}

func (c *CustomerAccountModification1) AddType() *TypeModification1 {
	c.Type = new(TypeModification1)
	return c.Type
}

func (c *CustomerAccountModification1) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CustomerAccountModification1) AddMonthlyPaymentValue() *AmountModification1 {
	c.MonthlyPaymentValue = new(AmountModification1)
	return c.MonthlyPaymentValue
}

func (c *CustomerAccountModification1) AddMonthlyReceivedValue() *AmountModification1 {
	c.MonthlyReceivedValue = new(AmountModification1)
	return c.MonthlyReceivedValue
}

func (c *CustomerAccountModification1) AddMonthlyTransactionNumber() *NumberModification1 {
	c.MonthlyTransactionNumber = new(NumberModification1)
	return c.MonthlyTransactionNumber
}

func (c *CustomerAccountModification1) AddAverageBalance() *AmountModification1 {
	c.AverageBalance = new(AmountModification1)
	return c.AverageBalance
}

func (c *CustomerAccountModification1) AddAccountPurpose() *PurposeModification1 {
	c.AccountPurpose = new(PurposeModification1)
	return c.AccountPurpose
}

func (c *CustomerAccountModification1) AddFloorNotificationAmount() *AmountModification1 {
	c.FloorNotificationAmount = new(AmountModification1)
	return c.FloorNotificationAmount
}

func (c *CustomerAccountModification1) AddCeilingNotificationAmount() *AmountModification1 {
	c.CeilingNotificationAmount = new(AmountModification1)
	return c.CeilingNotificationAmount
}

func (c *CustomerAccountModification1) AddStatementFrequencyAndFormat() *StatementFrequencyAndFormModification1 {
	newValue := new(StatementFrequencyAndFormModification1)
	c.StatementFrequencyAndFormat = append(c.StatementFrequencyAndFormat, newValue)
	return newValue
}

func (c *CustomerAccountModification1) AddClosingDate() *DateModification1 {
	c.ClosingDate = new(DateModification1)
	return c.ClosingDate
}

func (c *CustomerAccountModification1) AddRestriction() *RestrictionModification1 {
	newValue := new(RestrictionModification1)
	c.Restriction = append(c.Restriction, newValue)
	return newValue
}
