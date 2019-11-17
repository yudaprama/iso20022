package model

// Choice of formats for the reason of a pending status.
type PendingProcessingReason8Choice struct {

	// Reason for the pending status expressed as an ISO 20022 code.
	Code *ExternalPendingProcessingReason1Code `xml:"Cd"`

	// Reason for the pending status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (p *PendingProcessingReason8Choice) SetCode(value string) {
	p.Code = (*ExternalPendingProcessingReason1Code)(&value)
}

func (p *PendingProcessingReason8Choice) AddProprietary() *GenericIdentification36 {
	p.Proprietary = new(GenericIdentification36)
	return p.Proprietary
}
