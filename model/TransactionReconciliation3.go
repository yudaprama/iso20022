package model

// Reconciliation transaction between an acceptor and an acquirer.
type TransactionReconciliation3 struct {

	// Indicates if the transaction requires a closure of the reconciliation period.
	ClosePeriod *TrueFalseIndicator `xml:"ClsPrd,omitempty"`

	// Unique identification of a reconciliation transaction.
	ReconciliationTransactionIdentification *TransactionIdentifier1 `xml:"RcncltnTxId"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Transaction totals during the reconciliation period for a certain type of transaction.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls,omitempty"`

	// Additional information related to the reconciliation transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (t *TransactionReconciliation3) SetClosePeriod(value string) {
	t.ClosePeriod = (*TrueFalseIndicator)(&value)
}

func (t *TransactionReconciliation3) AddReconciliationTransactionIdentification() *TransactionIdentifier1 {
	t.ReconciliationTransactionIdentification = new(TransactionIdentifier1)
	return t.ReconciliationTransactionIdentification
}

func (t *TransactionReconciliation3) SetReconciliationIdentification(value string) {
	t.ReconciliationIdentification = (*Max35Text)(&value)
}

func (t *TransactionReconciliation3) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	t.TransactionTotals = append(t.TransactionTotals, newValue)
	return newValue
}

func (t *TransactionReconciliation3) SetAdditionalTransactionData(value string) {
	t.AdditionalTransactionData = (*Max70Text)(&value)
}
