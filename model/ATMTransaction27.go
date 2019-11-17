package model

// Transaction for which the exception is sent.
type ATMTransaction27 struct {

	// Identification of the transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId,omitempty"`

	// Identification of the reconciliation period.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Exception occurring outside the service.
	Exception []*FailureReason8Code `xml:"Xcptn"`

	// Explanation of the exception.
	ExceptionDetail []*Max70Text `xml:"XcptnDtl,omitempty"`

	// Balance of the captured card or epurse if available.
	ElectronicPurseBalance *CurrencyAndAmount `xml:"ElctrncPrsBal,omitempty"`
}

func (a *ATMTransaction27) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction27) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction27) AddException(value string) {
	a.Exception = append(a.Exception, (*FailureReason8Code)(&value))
}

func (a *ATMTransaction27) AddExceptionDetail(value string) {
	a.ExceptionDetail = append(a.ExceptionDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction27) SetElectronicPurseBalance(value, currency string) {
	a.ElectronicPurseBalance = NewCurrencyAndAmount(value, currency)
}
