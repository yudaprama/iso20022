package model

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation21 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Set of elements used to provide detailed information on the return reason.
	ReturnReasonInformation []*ReturnReasonInformation9 `xml:"RtrRsnInf,omitempty"`
}

func (o *OriginalGroupInformation21) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation21) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation21) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation21) AddReturnReasonInformation() *ReturnReasonInformation9 {
	newValue := new(ReturnReasonInformation9)
	o.ReturnReasonInformation = append(o.ReturnReasonInformation, newValue)
	return newValue
}
