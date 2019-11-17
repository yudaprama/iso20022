package model

// Detailed description of the differences.
type ValidationResult5 struct {

	// Sequential number assigned to the mismatch.
	SequenceNumber *Number `xml:"SeqNb"`

	// Coded identification of the matching rule that is violated.
	RuleIdentification *Max35Text `xml:"RuleId"`

	// Detailed description of the rule.
	RuleDescription *Max350Text `xml:"RuleDesc"`

	// Description of the element that creates the mismatch.
	MisMatchedElement []*ElementIdentification1 `xml:"MisMtchdElmt,omitempty"`
}

func (v *ValidationResult5) SetSequenceNumber(value string) {
	v.SequenceNumber = (*Number)(&value)
}

func (v *ValidationResult5) SetRuleIdentification(value string) {
	v.RuleIdentification = (*Max35Text)(&value)
}

func (v *ValidationResult5) SetRuleDescription(value string) {
	v.RuleDescription = (*Max350Text)(&value)
}

func (v *ValidationResult5) AddMisMatchedElement() *ElementIdentification1 {
	newValue := new(ElementIdentification1)
	v.MisMatchedElement = append(v.MisMatchedElement, newValue)
	return newValue
}
