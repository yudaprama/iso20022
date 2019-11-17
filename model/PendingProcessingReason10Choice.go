package model

// Choice of format for the pending processing reason.
type PendingProcessingReason10Choice struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending processing status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingProcessingReason10Choice) SetCode(value string) {
	p.Code = (*PendingProcessingReason2Code)(&value)
}

func (p *PendingProcessingReason10Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
