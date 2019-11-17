package model

// Entity hosting the ATM terminal.
type TerminalHosting1 struct {

	// Type of ATM terminal hosting.
	Category *TransactionEnvironment3Code `xml:"Ctgy,omitempty"`

	// Identify the entity hosting the ATM.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (t *TerminalHosting1) SetCategory(value string) {
	t.Category = (*TransactionEnvironment3Code)(&value)
}

func (t *TerminalHosting1) SetIdentification(value string) {
	t.Identification = (*Max35Text)(&value)
}
