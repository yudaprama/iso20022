package model

// Set of elements for the identification of the message and related references.
type References3 struct {

	// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Identification of the request message that has to be completed.
	RequestToBeCompletedIdentification *MessageIdentification1 `xml:"ReqToBeCmpltdId"`

	// Identifies a process by a unique identifier and the date and time when the first message belonging to the process was created by the sender. The process identification remains the same in all messages belonging to the same process, from the initial request message to the final account report closing the process.
	ProcessIdentification *MessageIdentification1 `xml:"PrcId"`

	// Reason of the request.
	RequestReason []*Max35Text `xml:"ReqRsn"`

	// File name of a document logically related to the request.
	AttachedDocumentName []*Max70Text `xml:"AttchdDocNm,omitempty"`
}

func (r *References3) AddMessageIdentification() *MessageIdentification1 {
	r.MessageIdentification = new(MessageIdentification1)
	return r.MessageIdentification
}

func (r *References3) AddRequestToBeCompletedIdentification() *MessageIdentification1 {
	r.RequestToBeCompletedIdentification = new(MessageIdentification1)
	return r.RequestToBeCompletedIdentification
}

func (r *References3) AddProcessIdentification() *MessageIdentification1 {
	r.ProcessIdentification = new(MessageIdentification1)
	return r.ProcessIdentification
}

func (r *References3) AddRequestReason(value string) {
	r.RequestReason = append(r.RequestReason, (*Max35Text)(&value))
}

func (r *References3) AddAttachedDocumentName(value string) {
	r.AttachedDocumentName = append(r.AttachedDocumentName, (*Max70Text)(&value))
}
