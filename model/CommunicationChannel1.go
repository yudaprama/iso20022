package model

// Communication channel information.
type CommunicationChannel1 struct {

	// Method by which the original undertaking or proposed amendment is to be made available.
	//
	Method *ExternalChannel1Code `xml:"Mtd"`

	// Type of party to whom the original undertaking or proposed amendment is intended to be delivered.
	DeliverToPartyType *PartyType1Choice `xml:"DlvrToPtyTp"`

	// Name of party to whom the original undertaking or proposed amendment is intended to be delivered.
	DeliverToName *Max140Text `xml:"DlvrToNm,omitempty"`

	// Address of party to whom the original undertaking or proposed amendment is intended to be delivered.
	DeliverToAddress *PostalAddress6 `xml:"DlvrToAdr,omitempty"`
}

func (c *CommunicationChannel1) SetMethod(value string) {
	c.Method = (*ExternalChannel1Code)(&value)
}

func (c *CommunicationChannel1) AddDeliverToPartyType() *PartyType1Choice {
	c.DeliverToPartyType = new(PartyType1Choice)
	return c.DeliverToPartyType
}

func (c *CommunicationChannel1) SetDeliverToName(value string) {
	c.DeliverToName = (*Max140Text)(&value)
}

func (c *CommunicationChannel1) AddDeliverToAddress() *PostalAddress6 {
	c.DeliverToAddress = new(PostalAddress6)
	return c.DeliverToAddress
}
