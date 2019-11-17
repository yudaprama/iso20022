package model

// Unique identification, as assigned by the original requestor, to unambiguously identify the business query message.
type OriginalBusinessQuery1 struct {

	// Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Specifies the query message name identifier to which the message refers.
	MessageNameIdentification *Max35Text `xml:"MsgNmId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (o *OriginalBusinessQuery1) SetMessageIdentification(value string) {
	o.MessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalBusinessQuery1) SetMessageNameIdentification(value string) {
	o.MessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalBusinessQuery1) SetCreationDateTime(value string) {
	o.CreationDateTime = (*ISODateTime)(&value)
}
