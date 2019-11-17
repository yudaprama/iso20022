package model

// Represents the assignment of a case to a party.
type CaseAssignment3 struct {

	// Uniquely identifies the case assignment.
	Identification *Max35Text `xml:"Id"`

	// Party who assigns the case.
	// Usage: This is also the sender of the message.
	Assigner *Party12Choice `xml:"Assgnr"`

	// Party to which the case is assigned.
	// Usage: This is also the receiver of the message.
	Assignee *Party12Choice `xml:"Assgne"`

	// Date and time at which the assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (c *CaseAssignment3) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CaseAssignment3) AddAssigner() *Party12Choice {
	c.Assigner = new(Party12Choice)
	return c.Assigner
}

func (c *CaseAssignment3) AddAssignee() *Party12Choice {
	c.Assignee = new(Party12Choice)
	return c.Assignee
}

func (c *CaseAssignment3) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}
