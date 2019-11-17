package model

// Contains the new, current and previous cut offs for a netting cut off held at the central system.
type CutOffData1 struct {

	// Describes the participant receiving the net report.
	ParticipantIdentification *PartyIdentification73Choice `xml:"PtcptId"`

	// Specifies the information regarding the updated netting cut off.
	NettingCutOffDetails []*NettingCutOff1 `xml:"NetgCutOffDtls"`
}

func (c *CutOffData1) AddParticipantIdentification() *PartyIdentification73Choice {
	c.ParticipantIdentification = new(PartyIdentification73Choice)
	return c.ParticipantIdentification
}

func (c *CutOffData1) AddNettingCutOffDetails() *NettingCutOff1 {
	newValue := new(NettingCutOff1)
	c.NettingCutOffDetails = append(c.NettingCutOffDetails, newValue)
	return newValue
}
