package model

// Provides the proprietary status and reason of an instruction or an instruction cancellation.
type ProprietaryStatusAndReason7 struct {

	// Proprietary identification of the status related to an instruction.
	ProprietaryStatus *GenericIdentification47 `xml:"PrtrySts"`

	// Proprietary identification of the reason related to a proprietary status.
	ProprietaryReason []*ProprietaryReason5 `xml:"PrtryRsn,omitempty"`
}

func (p *ProprietaryStatusAndReason7) AddProprietaryStatus() *GenericIdentification47 {
	p.ProprietaryStatus = new(GenericIdentification47)
	return p.ProprietaryStatus
}

func (p *ProprietaryStatusAndReason7) AddProprietaryReason() *ProprietaryReason5 {
	newValue := new(ProprietaryReason5)
	p.ProprietaryReason = append(p.ProprietaryReason, newValue)
	return newValue
}
