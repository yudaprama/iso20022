package model

// Unique and unambiguous identification of the original message references.
type OriginalMessage2 struct {

	// Original message sender used to identify the message.
	OriginalSender *Party28Choice `xml:"OrgnlSndr,omitempty"`

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, such as pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Original date and time at which the message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Specifies the identification of original package of instructions, entries or records.
	OriginalPackageIdentification *Max35Text `xml:"OrgnlPackgId,omitempty"`

	// Specifies the identification of original entry, instruction or record within the package.
	OriginalRecordIdentification *Max35Text `xml:"OrgnlRcrdId"`
}

func (o *OriginalMessage2) AddOriginalSender() *Party28Choice {
	o.OriginalSender = new(Party28Choice)
	return o.OriginalSender
}

func (o *OriginalMessage2) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage2) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage2) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalMessage2) SetOriginalPackageIdentification(value string) {
	o.OriginalPackageIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage2) SetOriginalRecordIdentification(value string) {
	o.OriginalRecordIdentification = (*Max35Text)(&value)
}
