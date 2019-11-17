package model

// Specifies the details of the underlying transaction on which the investigation is processed.
type UnderlyingTransaction1Choice struct {

	// Set of elements used to reference the details of the original payment initiation.
	Initiation *UnderlyingPaymentInstruction1 `xml:"Initn"`

	// Set of elements used to reference the details of the original interbank payment transaction.
	Interbank *UnderlyingPaymentTransaction1 `xml:"IntrBk"`

	// Reference details on the underlying statement cash entry.
	StatementEntry *UnderlyingStatementEntry1 `xml:"StmtNtry"`
}

func (u *UnderlyingTransaction1Choice) AddInitiation() *UnderlyingPaymentInstruction1 {
	u.Initiation = new(UnderlyingPaymentInstruction1)
	return u.Initiation
}

func (u *UnderlyingTransaction1Choice) AddInterbank() *UnderlyingPaymentTransaction1 {
	u.Interbank = new(UnderlyingPaymentTransaction1)
	return u.Interbank
}

func (u *UnderlyingTransaction1Choice) AddStatementEntry() *UnderlyingStatementEntry1 {
	u.StatementEntry = new(UnderlyingStatementEntry1)
	return u.StatementEntry
}
