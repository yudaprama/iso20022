package model

// Provides the proprietary status and reason of an instruction or an instruction cancellation.
type ProprietaryStatusAndReason1 struct {

	// Proprietary identification of the status related to an instruction.
	ProprietaryStatus *GenericIdentification20 `xml:"PrtrySts"`

	// Proprietary identification of the reason related to a proprietary status.
	ProprietaryReason []*ProprietaryReason1 `xml:"PrtryRsn,omitempty"`
}

func (p *ProprietaryStatusAndReason1) AddProprietaryStatus() *GenericIdentification20 {
	p.ProprietaryStatus = new(GenericIdentification20)
	return p.ProprietaryStatus
}

func (p *ProprietaryStatusAndReason1) AddProprietaryReason() *ProprietaryReason1 {
	newValue := new(ProprietaryReason1)
	p.ProprietaryReason = append(p.ProprietaryReason, newValue)
	return newValue
}
