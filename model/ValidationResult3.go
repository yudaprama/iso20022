package model

// Describes the error that is the cause of the rejection.
type ValidationResult3 struct {

	// Sequential number assigned to the error.
	SequenceNumber *Number `xml:"SeqNb"`

	// Coded identification of the rule that is violated by the rejected message.
	RuleIdentification *Max35Text `xml:"RuleId"`

	// Detailed description of the rule.
	RuleDescription *Max350Text `xml:"RuleDesc"`

	// Description of the elements that violated the rule.
	Element []*ElementIdentification3 `xml:"Elmt,omitempty"`
}

func (v *ValidationResult3) SetSequenceNumber(value string) {
	v.SequenceNumber = (*Number)(&value)
}

func (v *ValidationResult3) SetRuleIdentification(value string) {
	v.RuleIdentification = (*Max35Text)(&value)
}

func (v *ValidationResult3) SetRuleDescription(value string) {
	v.RuleDescription = (*Max350Text)(&value)
}

func (v *ValidationResult3) AddElement() *ElementIdentification3 {
	newValue := new(ElementIdentification3)
	v.Element = append(v.Element, newValue)
	return newValue
}
