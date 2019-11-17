package model

// Specifies the details of the underlying transaction on which the investigation is processed.
type UnderlyingTransaction3Choice struct {

	// Set of elements used to reference the details of the original payment initiation.
	Initiation *UnderlyingPaymentInstruction3 `xml:"Initn"`

	// Set of elements used to reference the details of the original interbank payment transaction.
	Interbank *UnderlyingPaymentTransaction2 `xml:"IntrBk"`

	// Reference details on the underlying statement cash entry.
	StatementEntry *UnderlyingStatementEntry1 `xml:"StmtNtry"`
}

func (u *UnderlyingTransaction3Choice) AddInitiation() *UnderlyingPaymentInstruction3 {
	u.Initiation = new(UnderlyingPaymentInstruction3)
	return u.Initiation
}

func (u *UnderlyingTransaction3Choice) AddInterbank() *UnderlyingPaymentTransaction2 {
	u.Interbank = new(UnderlyingPaymentTransaction2)
	return u.Interbank
}

func (u *UnderlyingTransaction3Choice) AddStatementEntry() *UnderlyingStatementEntry1 {
	u.StatementEntry = new(UnderlyingStatementEntry1)
	return u.StatementEntry
}
