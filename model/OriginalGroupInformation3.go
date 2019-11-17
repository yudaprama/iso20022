package model

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation3 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Original date and time at which the message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`
}

func (o *OriginalGroupInformation3) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation3) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation3) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}
