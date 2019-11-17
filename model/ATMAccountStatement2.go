package model

// Statement information of an account.
type ATMAccountStatement2 struct {

	// Date of the transaction.
	TransactionDate *ISODate `xml:"TxDt,omitempty"`

	// Value date of the transaction.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Short text to display or print for the statement.
	ShortText *Max70Text `xml:"ShrtTxt,omitempty"`

	// True if credit transaction.
	CreditTransaction *TrueFalseIndicator `xml:"CdtTx,omitempty"`

	// Amount of the transaction.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Alternative text of the statement to print or display.
	LongText *Max256Text `xml:"LngTxt,omitempty"`
}

func (a *ATMAccountStatement2) SetTransactionDate(value string) {
	a.TransactionDate = (*ISODate)(&value)
}

func (a *ATMAccountStatement2) SetValueDate(value string) {
	a.ValueDate = (*ISODate)(&value)
}

func (a *ATMAccountStatement2) SetShortText(value string) {
	a.ShortText = (*Max70Text)(&value)
}

func (a *ATMAccountStatement2) SetCreditTransaction(value string) {
	a.CreditTransaction = (*TrueFalseIndicator)(&value)
}

func (a *ATMAccountStatement2) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMAccountStatement2) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMAccountStatement2) SetLongText(value string) {
	a.LongText = (*Max256Text)(&value)
}
