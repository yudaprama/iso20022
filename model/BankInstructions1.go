package model

// Details of the instructions from the bank.
type BankInstructions1 struct {

	// Instructions from the bank.
	Text []*Max2000Text `xml:"Txt,omitempty"`

	// Last date for a response to the bank instructions.
	LastDateForResponse *ISODate `xml:"LastDtForRspn,omitempty"`
}

func (b *BankInstructions1) AddText(value string) {
	b.Text = append(b.Text, (*Max2000Text)(&value))
}

func (b *BankInstructions1) SetLastDateForResponse(value string) {
	b.LastDateForResponse = (*ISODate)(&value)
}
