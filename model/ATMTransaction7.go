package model

// Inquiry information for the transaction.
type ATMTransaction7 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Result of the inquiry service.
	TransactionResponse *ResponseType3 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action5 `xml:"Actn,omitempty"`

	// Profile of the customer with the allowed services and restrictions.
	CustomerServiceProfile *ATMCustomerProfile3 `xml:"CstmrSvcPrfl,omitempty"`

	// Dynamic currency conversion result.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs,omitempty"`

	// Account information associated to the customer.
	AccountInformation []*CardAccount6 `xml:"AcctInf,omitempty"`

	// Statement information of an account.
	AccountStatementData []*ATMAccountStatement1 `xml:"AcctStmtData,omitempty"`

	// Exchange rate and calculated amount to be presented to the customer when the dispense currency or the deposit currency (target currency) is different to account currency (source currency).
	CurrencyExchange *CurrencyConversion5 `xml:"CcyXchg,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction7) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction7) AddTransactionResponse() *ResponseType3 {
	a.TransactionResponse = new(ResponseType3)
	return a.TransactionResponse
}

func (a *ATMTransaction7) AddAction() *Action5 {
	newValue := new(Action5)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction7) AddCustomerServiceProfile() *ATMCustomerProfile3 {
	a.CustomerServiceProfile = new(ATMCustomerProfile3)
	return a.CustomerServiceProfile
}

func (a *ATMTransaction7) AddCurrencyConversion() *CurrencyConversion3 {
	a.CurrencyConversion = new(CurrencyConversion3)
	return a.CurrencyConversion
}

func (a *ATMTransaction7) AddAccountInformation() *CardAccount6 {
	newValue := new(CardAccount6)
	a.AccountInformation = append(a.AccountInformation, newValue)
	return newValue
}

func (a *ATMTransaction7) AddAccountStatementData() *ATMAccountStatement1 {
	newValue := new(ATMAccountStatement1)
	a.AccountStatementData = append(a.AccountStatementData, newValue)
	return newValue
}

func (a *ATMTransaction7) AddCurrencyExchange() *CurrencyConversion5 {
	a.CurrencyExchange = new(CurrencyConversion5)
	return a.CurrencyExchange
}

func (a *ATMTransaction7) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction7) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}
