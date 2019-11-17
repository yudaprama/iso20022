package model

// Provides information on the original group, to which the message refers.
type OriginalGroupHeader2 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Provides detailed information on the return reason.
	ReturnReasonInformation []*PaymentReturnReason1 `xml:"RtrRsnInf,omitempty"`
}

func (o *OriginalGroupHeader2) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader2) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader2) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader2) AddReturnReasonInformation() *PaymentReturnReason1 {
	newValue := new(PaymentReturnReason1)
	o.ReturnReasonInformation = append(o.ReturnReasonInformation, newValue)
	return newValue
}
