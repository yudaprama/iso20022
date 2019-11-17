package model

// Provides the details for one to many netting cut off update requests to be actioned by a central system.
type NettingCutOff1 struct {

	// Identifies the netting party or group.
	NettingIdentification *NettingIdentification1Choice `xml:"NetgId"`

	// Specifies the information regarding the updated netting cut off.
	NewCutOff []*CutOff1 `xml:"NewCutOff"`
}

func (n *NettingCutOff1) AddNettingIdentification() *NettingIdentification1Choice {
	n.NettingIdentification = new(NettingIdentification1Choice)
	return n.NettingIdentification
}

func (n *NettingCutOff1) AddNewCutOff() *CutOff1 {
	newValue := new(CutOff1)
	n.NewCutOff = append(n.NewCutOff, newValue)
	return newValue
}
