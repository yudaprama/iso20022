package model

// Choice of format for the pending processing reason.
type PendingProcessingReason14Choice struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending processing status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingProcessingReason14Choice) SetCode(value string) {
	p.Code = (*PendingProcessingReason2Code)(&value)
}

func (p *PendingProcessingReason14Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
