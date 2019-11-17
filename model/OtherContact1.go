package model

// Communication device number or electronic address used for communication.
type OtherContact1 struct {

	// Method used to contact the financial institution’s contact for the specific tax region.
	ChannelType *Max4Text `xml:"ChanlTp"`

	// Communication value such as phone number or email address.
	Identification *Max128Text `xml:"Id,omitempty"`
}

func (o *OtherContact1) SetChannelType(value string) {
	o.ChannelType = (*Max4Text)(&value)
}

func (o *OtherContact1) SetIdentification(value string) {
	o.Identification = (*Max128Text)(&value)
}
