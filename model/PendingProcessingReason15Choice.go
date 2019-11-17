package model

// Choice of format for the pending processing reason.
type PendingProcessingReason15Choice struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason3Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending processing status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingProcessingReason15Choice) SetCode(value string) {
	p.Code = (*PendingProcessingReason3Code)(&value)
}

func (p *PendingProcessingReason15Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
