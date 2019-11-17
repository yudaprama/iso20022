package model

// Specifies the type of change to statement frequency and form.
type StatementFrequencyAndFormModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Specifies the statement frequency, format, delivery address.
	StatementFrequencyAndForm *StatementFrequencyAndForm1 `xml:"StmtFrqcyAndForm"`
}

func (s *StatementFrequencyAndFormModification1) SetModificationCode(value string) {
	s.ModificationCode = (*Modification1Code)(&value)
}

func (s *StatementFrequencyAndFormModification1) AddStatementFrequencyAndForm() *StatementFrequencyAndForm1 {
	s.StatementFrequencyAndForm = new(StatementFrequencyAndForm1)
	return s.StatementFrequencyAndForm
}
