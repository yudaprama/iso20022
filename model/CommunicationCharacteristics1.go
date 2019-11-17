package model

// Low level communication of the hardware or software component toward another component or an external entity.
type CommunicationCharacteristics1 struct {

	// Type of low level communication.
	CommunicationType *POICommunicationType1Code `xml:"ComTp"`

	// Entity that communicate with the current component, using this communication device.
	RemoteParty *PartyType7Code `xml:"RmotPty"`

	// Communication hardware is activated.
	Active *TrueFalseIndicator `xml:"Actv"`
}

func (c *CommunicationCharacteristics1) SetCommunicationType(value string) {
	c.CommunicationType = (*POICommunicationType1Code)(&value)
}

func (c *CommunicationCharacteristics1) SetRemoteParty(value string) {
	c.RemoteParty = (*PartyType7Code)(&value)
}

func (c *CommunicationCharacteristics1) SetActive(value string) {
	c.Active = (*TrueFalseIndicator)(&value)
}
