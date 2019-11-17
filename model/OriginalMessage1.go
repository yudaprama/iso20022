package model

// Identification of a message previously sent.
type OriginalMessage1 struct {

	// XML schema-instance namespace, for example "tsin.008.001.01"
	MessageDefinitionIdentifier *Max35Text `xml:"MsgDefIdr"`

	// Message sender specified in the original message.
	//
	From *Party9Choice `xml:"Fr"`

	// Message recipient specified in the original message.
	To *Party9Choice `xml:"To"`

	// Message identification specified in the original message.
	BusinessMessageIdentifier *Max35Text `xml:"BizMsgIdr"`

	// Message creation date and time specified in the original message.
	CreationDate *ISONormalisedDateTime `xml:"CreDt"`

	// Indicates whether the message is a copy, a duplicate or a copy of a duplicate of a previously sent ISO 20022 message.
	CopyDuplicate *CopyDuplicate1Code `xml:"CpyDplct,omitempty"`
}

func (o *OriginalMessage1) SetMessageDefinitionIdentifier(value string) {
	o.MessageDefinitionIdentifier = (*Max35Text)(&value)
}

func (o *OriginalMessage1) AddFrom() *Party9Choice {
	o.From = new(Party9Choice)
	return o.From
}

func (o *OriginalMessage1) AddTo() *Party9Choice {
	o.To = new(Party9Choice)
	return o.To
}

func (o *OriginalMessage1) SetBusinessMessageIdentifier(value string) {
	o.BusinessMessageIdentifier = (*Max35Text)(&value)
}

func (o *OriginalMessage1) SetCreationDate(value string) {
	o.CreationDate = (*ISONormalisedDateTime)(&value)
}

func (o *OriginalMessage1) SetCopyDuplicate(value string) {
	o.CopyDuplicate = (*CopyDuplicate1Code)(&value)
}
