package model

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation5 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation1 `xml:"RvslRsnInf,omitempty"`
}

func (o *OriginalGroupInformation5) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation5) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation5) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation5) AddReversalReasonInformation() *ReversalReasonInformation1 {
	newValue := new(ReversalReasonInformation1)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}
