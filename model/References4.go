package model

// Set of elements for the identification of the message and related references.
type References4 struct {

	// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Identifies a process by a unique identifier and the date and time when the first message belonging to the process was created by the sender. The process identification remains the same in all messages belonging to the same process, from the initial request message to the final account report closing the process.
	ProcessIdentification *MessageIdentification1 `xml:"PrcId"`

	// File name of a document logically related to the request.
	AttachedDocumentName []*Max70Text `xml:"AttchdDocNm,omitempty"`
}

func (r *References4) AddMessageIdentification() *MessageIdentification1 {
	r.MessageIdentification = new(MessageIdentification1)
	return r.MessageIdentification
}

func (r *References4) AddProcessIdentification() *MessageIdentification1 {
	r.ProcessIdentification = new(MessageIdentification1)
	return r.ProcessIdentification
}

func (r *References4) AddAttachedDocumentName(value string) {
	r.AttachedDocumentName = append(r.AttachedDocumentName, (*Max70Text)(&value))
}
