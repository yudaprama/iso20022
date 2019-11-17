package model

// Choice of format for the pending processing reason.
type PendingProcessingReason11Choice struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending processing status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingProcessingReason11Choice) SetCode(value string) {
	p.Code = (*PendingProcessingReason1Code)(&value)
}

func (p *PendingProcessingReason11Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
