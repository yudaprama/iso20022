package model

// Set of elements for the identification of the message and related references.
type References6 struct {

	// Identify the type of rejected request.
	RejectedRequestType *UseCases1Code `xml:"RjctdReqTp"`

	// Reason of the message rejection.
	RejectionReason []*Max350Text `xml:"RjctnRsn"`

	// Identification of the rejected request message.
	RejectedRequestIdentification *MessageIdentification1 `xml:"RjctdReqId"`

	// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Identifies a process by a unique identifier and the date and time when the first message belonging to the process was created by the sender. The process identification remains the same in all messages belonging to the same process, from the initial request message to the final account report closing the process.
	ProcessIdentification *MessageIdentification1 `xml:"PrcId"`

	// File name of a document logically related to the request.
	AttachedDocumentName []*Max70Text `xml:"AttchdDocNm,omitempty"`
}

func (r *References6) SetRejectedRequestType(value string) {
	r.RejectedRequestType = (*UseCases1Code)(&value)
}

func (r *References6) AddRejectionReason(value string) {
	r.RejectionReason = append(r.RejectionReason, (*Max350Text)(&value))
}

func (r *References6) AddRejectedRequestIdentification() *MessageIdentification1 {
	r.RejectedRequestIdentification = new(MessageIdentification1)
	return r.RejectedRequestIdentification
}

func (r *References6) AddMessageIdentification() *MessageIdentification1 {
	r.MessageIdentification = new(MessageIdentification1)
	return r.MessageIdentification
}

func (r *References6) AddProcessIdentification() *MessageIdentification1 {
	r.ProcessIdentification = new(MessageIdentification1)
	return r.ProcessIdentification
}

func (r *References6) AddAttachedDocumentName(value string) {
	r.AttachedDocumentName = append(r.AttachedDocumentName, (*Max70Text)(&value))
}
