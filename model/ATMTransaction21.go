package model

// Inquiry information for the transaction.
type ATMTransaction21 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Result of the inquiry service.
	TransactionResponse *ResponseType7 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action7 `xml:"Actn,omitempty"`

	// Profile of the customer with the allowed services and restrictions.
	CustomerServiceProfile *ATMCustomerProfile5 `xml:"CstmrSvcPrfl,omitempty"`

	// Dynamic currency conversion result.
	CurrencyConversion *CurrencyConversion10 `xml:"CcyConvs,omitempty"`

	// Account information associated to the customer.
	AccountInformation []*CardAccount12 `xml:"AcctInf,omitempty"`

	// Statement information of an account.
	AccountStatementData []*ATMAccountStatement1 `xml:"AcctStmtData,omitempty"`

	// Exchange rate and calculated amount to be presented to the customer when the dispense currency or the deposit currency (target currency) is different to account currency (source currency).
	CurrencyExchange *CurrencyConversion5 `xml:"CcyXchg,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction21) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction21) AddTransactionResponse() *ResponseType7 {
	a.TransactionResponse = new(ResponseType7)
	return a.TransactionResponse
}

func (a *ATMTransaction21) AddAction() *Action7 {
	newValue := new(Action7)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction21) AddCustomerServiceProfile() *ATMCustomerProfile5 {
	a.CustomerServiceProfile = new(ATMCustomerProfile5)
	return a.CustomerServiceProfile
}

func (a *ATMTransaction21) AddCurrencyConversion() *CurrencyConversion10 {
	a.CurrencyConversion = new(CurrencyConversion10)
	return a.CurrencyConversion
}

func (a *ATMTransaction21) AddAccountInformation() *CardAccount12 {
	newValue := new(CardAccount12)
	a.AccountInformation = append(a.AccountInformation, newValue)
	return newValue
}

func (a *ATMTransaction21) AddAccountStatementData() *ATMAccountStatement1 {
	newValue := new(ATMAccountStatement1)
	a.AccountStatementData = append(a.AccountStatementData, newValue)
	return newValue
}

func (a *ATMTransaction21) AddCurrencyExchange() *CurrencyConversion5 {
	a.CurrencyExchange = new(CurrencyConversion5)
	return a.CurrencyExchange
}

func (a *ATMTransaction21) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction21) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}
