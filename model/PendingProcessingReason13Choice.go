package model

// Choice of format for the pending processing reason.
type PendingProcessingReason13Choice struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending processing status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingProcessingReason13Choice) SetCode(value string) {
	p.Code = (*PendingProcessingReason1Code)(&value)
}

func (p *PendingProcessingReason13Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
