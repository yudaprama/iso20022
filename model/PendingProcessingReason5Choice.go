package model

// Choice of format for the pending processing reason.
type PendingProcessingReason5Choice struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason3Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending processing status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingProcessingReason5Choice) SetCode(value string) {
	p.Code = (*PendingProcessingReason3Code)(&value)
}

func (p *PendingProcessingReason5Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
