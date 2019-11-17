package model

// Specifies the details of the underlying transaction on which the investigation is processed.
type UnderlyingTransaction2Choice struct {

	// Set of elements used to reference the details of the original payment initiation.
	Initiation *UnderlyingPaymentInstruction2 `xml:"Initn"`

	// Set of elements used to reference the details of the original interbank payment transaction.
	Interbank *UnderlyingPaymentTransaction2 `xml:"IntrBk"`

	// Reference details on the underlying statement cash entry.
	StatementEntry *UnderlyingStatementEntry1 `xml:"StmtNtry"`
}

func (u *UnderlyingTransaction2Choice) AddInitiation() *UnderlyingPaymentInstruction2 {
	u.Initiation = new(UnderlyingPaymentInstruction2)
	return u.Initiation
}

func (u *UnderlyingTransaction2Choice) AddInterbank() *UnderlyingPaymentTransaction2 {
	u.Interbank = new(UnderlyingPaymentTransaction2)
	return u.Interbank
}

func (u *UnderlyingTransaction2Choice) AddStatementEntry() *UnderlyingStatementEntry1 {
	u.StatementEntry = new(UnderlyingStatementEntry1)
	return u.StatementEntry
}
