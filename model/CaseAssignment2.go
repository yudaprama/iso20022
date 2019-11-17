package model

// Set of elements used to represent the assignment of a case to a party.
type CaseAssignment2 struct {

	// Uniquely identifies the case assignment.
	Identification *Max35Text `xml:"Id"`

	// Party who assigns the case.
	// Usage: This is also the sender of the message.
	Assigner *Party7Choice `xml:"Assgnr"`

	// Party to which the case is assigned.
	// Usage: This is also the receiver of the message.
	Assignee *Party7Choice `xml:"Assgne"`

	// Date and time at which the assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (c *CaseAssignment2) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CaseAssignment2) AddAssigner() *Party7Choice {
	c.Assigner = new(Party7Choice)
	return c.Assigner
}

func (c *CaseAssignment2) AddAssignee() *Party7Choice {
	c.Assignee = new(Party7Choice)
	return c.Assignee
}

func (c *CaseAssignment2) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}
