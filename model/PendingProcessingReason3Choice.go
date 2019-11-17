package model

// Choice of format for the pending processing reason.
type PendingProcessingReason3Choice struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending processing status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingProcessingReason3Choice) SetCode(value string) {
	p.Code = (*PendingProcessingReason2Code)(&value)
}

func (p *PendingProcessingReason3Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
