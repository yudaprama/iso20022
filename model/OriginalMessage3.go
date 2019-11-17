package model

// Unique and unambiguous identification of the original message references.
type OriginalMessage3 struct {

	// Original message sender used to identify the message.
	OriginalSender *Party28Choice `xml:"OrgnlSndr,omitempty"`

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, such as pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Original date and time at which the message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`
}

func (o *OriginalMessage3) AddOriginalSender() *Party28Choice {
	o.OriginalSender = new(Party28Choice)
	return o.OriginalSender
}

func (o *OriginalMessage3) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage3) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage3) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}
