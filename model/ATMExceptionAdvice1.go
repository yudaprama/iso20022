package model

// Information related to exceptions occurring on the ATM.
type ATMExceptionAdvice1 struct {

	// Environment of the exception.
	Environment *ATMEnvironment16 `xml:"Envt,omitempty"`

	// Context of the exception.
	Context *ATMContext20 `xml:"Cntxt,omitempty"`

	// Transaction for which the exception is sent.
	Transaction *ATMTransaction27 `xml:"Tx"`
}

func (a *ATMExceptionAdvice1) AddEnvironment() *ATMEnvironment16 {
	a.Environment = new(ATMEnvironment16)
	return a.Environment
}

func (a *ATMExceptionAdvice1) AddContext() *ATMContext20 {
	a.Context = new(ATMContext20)
	return a.Context
}

func (a *ATMExceptionAdvice1) AddTransaction() *ATMTransaction27 {
	a.Transaction = new(ATMTransaction27)
	return a.Transaction
}
