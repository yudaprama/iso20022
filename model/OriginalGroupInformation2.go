package model

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation2 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Detailed information on the return reason.
	ReturnReasonInformation []*ReturnReasonInformation1 `xml:"RtrRsnInf,omitempty"`
}

func (o *OriginalGroupInformation2) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation2) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation2) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation2) AddReturnReasonInformation() *ReturnReasonInformation1 {
	newValue := new(ReturnReasonInformation1)
	o.ReturnReasonInformation = append(o.ReturnReasonInformation, newValue)
	return newValue
}
