package model

// Choice of formats for the transaction type.
type TransactionType1CodeChoice struct {

	// Transaction type in a structured format.
	Structured *TransactionType2Code `xml:"Strd"`

	// Transaction type in free text form.
	Unstructured *Max35Text `xml:"Ustrd"`
}

func (t *TransactionType1CodeChoice) SetStructured(value string) {
	t.Structured = (*TransactionType2Code)(&value)
}

func (t *TransactionType1CodeChoice) SetUnstructured(value string) {
	t.Unstructured = (*Max35Text)(&value)
}
