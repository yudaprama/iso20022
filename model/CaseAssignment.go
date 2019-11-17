package model

// Represents the assignment of a case to a party. Assignment is a step in the overall process of managing a case.
type CaseAssignment struct {

	// Identification of an assignment within a case.
	Identification *Max35Text `xml:"Id"`

	// Party that assigns the case to another party. This is also the sender of the message.
	Assigner *AnyBICIdentifier `xml:"Assgnr"`

	// Party that the case is assigned to. This is also the receiver of the message.
	Assignee *AnyBICIdentifier `xml:"Assgne"`

	// Date and time at which the assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (c *CaseAssignment) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CaseAssignment) SetAssigner(value string) {
	c.Assigner = (*AnyBICIdentifier)(&value)
}

func (c *CaseAssignment) SetAssignee(value string) {
	c.Assignee = (*AnyBICIdentifier)(&value)
}

func (c *CaseAssignment) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}
