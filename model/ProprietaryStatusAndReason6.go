package model

// Provides the proprietary status and reason of an instruction or an instruction cancellation.
type ProprietaryStatusAndReason6 struct {

	// Proprietary identification of the status related to an instruction.
	ProprietaryStatus *GenericIdentification30 `xml:"PrtrySts"`

	// Proprietary identification of the reason related to a proprietary status.
	ProprietaryReason []*ProprietaryReason4 `xml:"PrtryRsn,omitempty"`
}

func (p *ProprietaryStatusAndReason6) AddProprietaryStatus() *GenericIdentification30 {
	p.ProprietaryStatus = new(GenericIdentification30)
	return p.ProprietaryStatus
}

func (p *ProprietaryStatusAndReason6) AddProprietaryReason() *ProprietaryReason4 {
	newValue := new(ProprietaryReason4)
	p.ProprietaryReason = append(p.ProprietaryReason, newValue)
	return newValue
}
