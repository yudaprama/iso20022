package model

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation22 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Set of elements used to provide detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation6 `xml:"RvslRsnInf,omitempty"`
}

func (o *OriginalGroupInformation22) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation22) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation22) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation22) AddReversalReasonInformation() *ReversalReasonInformation6 {
	newValue := new(ReversalReasonInformation6)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}
