package model

// Unique identification, as assigned by the original instructing party, to unambiguously identify the message.
type OriginalMessageInformation1 struct {

	// Point to point reference, as assigned by the original initiating party, to unambiguously identify the original mandate request message.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Specifies the message name identifier to which the message refers.
	MessageNameIdentification *Max35Text `xml:"MsgNmId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (o *OriginalMessageInformation1) SetMessageIdentification(value string) {
	o.MessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessageInformation1) SetMessageNameIdentification(value string) {
	o.MessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessageInformation1) SetCreationDateTime(value string) {
	o.CreationDateTime = (*ISODateTime)(&value)
}
