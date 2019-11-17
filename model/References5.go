package model

// Set of elements for the identification of the message and related references.
type References5 struct {

	// Identifies the type of acknowledged request.
	RequestType *UseCases1Code `xml:"ReqTp"`

	// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Identifies a process by a unique identifier and the date and time when the first message belonging to the process was created by the sender. The process identification remains the same in all messages belonging to the same process, from the initial request message to the final account report closing the process.
	ProcessIdentification *MessageIdentification1 `xml:"PrcId"`

	// Reference to the message that is acknowledged.
	AcknowledgedMessageIdentification []*MessageIdentification1 `xml:"AckdMsgId,omitempty"`

	// Status of the request.
	Status *Max35Text `xml:"Sts,omitempty"`

	// File name of a document logically related to the request.
	AttachedDocumentName []*Max70Text `xml:"AttchdDocNm,omitempty"`
}

func (r *References5) SetRequestType(value string) {
	r.RequestType = (*UseCases1Code)(&value)
}

func (r *References5) AddMessageIdentification() *MessageIdentification1 {
	r.MessageIdentification = new(MessageIdentification1)
	return r.MessageIdentification
}

func (r *References5) AddProcessIdentification() *MessageIdentification1 {
	r.ProcessIdentification = new(MessageIdentification1)
	return r.ProcessIdentification
}

func (r *References5) AddAcknowledgedMessageIdentification() *MessageIdentification1 {
	newValue := new(MessageIdentification1)
	r.AcknowledgedMessageIdentification = append(r.AcknowledgedMessageIdentification, newValue)
	return newValue
}

func (r *References5) SetStatus(value string) {
	r.Status = (*Max35Text)(&value)
}

func (r *References5) AddAttachedDocumentName(value string) {
	r.AttachedDocumentName = append(r.AttachedDocumentName, (*Max70Text)(&value))
}
