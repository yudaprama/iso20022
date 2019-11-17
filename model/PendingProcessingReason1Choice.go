package model

// Choice of format for the pending processing reason.
type PendingProcessingReason1Choice struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a pending processing status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingProcessingReason1Choice) SetCode(value string) {
	p.Code = (*PendingProcessingReason1Code)(&value)
}

func (p *PendingProcessingReason1Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
