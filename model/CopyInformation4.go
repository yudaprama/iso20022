package model

// Information provided when the message is a copy of a previous message.
type CopyInformation4 struct {

	// Indicates whether the message is a copy.
	CopyIndicator *YesNoIndicator `xml:"CpyInd"`

	// Original receiver of the message, if this message is a copy.
	OriginalReceiver *AnyBICIdentifier `xml:"OrgnlRcvr,omitempty"`
}

func (c *CopyInformation4) SetCopyIndicator(value string) {
	c.CopyIndicator = (*YesNoIndicator)(&value)
}

func (c *CopyInformation4) SetOriginalReceiver(value string) {
	c.OriginalReceiver = (*AnyBICIdentifier)(&value)
}
