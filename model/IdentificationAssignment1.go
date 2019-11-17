package model

// Set of elements that identify the identification assignment.
type IdentificationAssignment1 struct {

	// Point to point reference, as assigned by the assigner, and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The assigner has to make sure that MessageIdentification is unique per assignee for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the identification assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that created the identification assignment.
	Creator *Party7Choice `xml:"Cretr,omitempty"`

	// Party that assigns the identification assignment to another party. This is also the sender of the message.
	Assigner *Party7Choice `xml:"Assgnr"`

	// Party that the identification assignment is assigned to. This is also the receiver of the message.
	Assignee *Party7Choice `xml:"Assgne"`
}

func (i *IdentificationAssignment1) SetMessageIdentification(value string) {
	i.MessageIdentification = (*Max35Text)(&value)
}

func (i *IdentificationAssignment1) SetCreationDateTime(value string) {
	i.CreationDateTime = (*ISODateTime)(&value)
}

func (i *IdentificationAssignment1) AddCreator() *Party7Choice {
	i.Creator = new(Party7Choice)
	return i.Creator
}

func (i *IdentificationAssignment1) AddAssigner() *Party7Choice {
	i.Assigner = new(Party7Choice)
	return i.Assigner
}

func (i *IdentificationAssignment1) AddAssignee() *Party7Choice {
	i.Assignee = new(Party7Choice)
	return i.Assignee
}
