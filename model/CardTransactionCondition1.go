package model

// Data used to assign specific condition such as liability shift or preferential interchange fees.
type CardTransactionCondition1 struct {

	// Identification of the specific condition.
	Program *Max35Text `xml:"Prgm"`

	// Level of the condition.
	Value *Max35Text `xml:"Val,omitempty"`
}

func (c *CardTransactionCondition1) SetProgram(value string) {
	c.Program = (*Max35Text)(&value)
}

func (c *CardTransactionCondition1) SetValue(value string) {
	c.Value = (*Max35Text)(&value)
}
