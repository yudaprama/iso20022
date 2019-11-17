package model

// Unique identification, as assigned by the first instructing agent, to unambiguously identify the group of transactions.
type UnderlyingGroupInformation1 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Original channel used for the delivery of the message, to allow the receiver of the request to locate the payment with greater ease.
	OriginalMessageDeliveryChannel *Max35Text `xml:"OrgnlMsgDlvryChanl,omitempty"`
}

func (u *UnderlyingGroupInformation1) SetOriginalMessageIdentification(value string) {
	u.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingGroupInformation1) SetOriginalMessageNameIdentification(value string) {
	u.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingGroupInformation1) SetOriginalCreationDateTime(value string) {
	u.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (u *UnderlyingGroupInformation1) SetOriginalMessageDeliveryChannel(value string) {
	u.OriginalMessageDeliveryChannel = (*Max35Text)(&value)
}
