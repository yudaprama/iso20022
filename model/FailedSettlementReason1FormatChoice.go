package model

// Choice between a standard code or a proprietary code for specifying the reason of a settlement failure.
type FailedSettlementReason1FormatChoice struct {

	// Standard code to specify the reason of a settlement failure.
	Code *FailedSettlementReason1Code `xml:"Cd"`

	// Proprietary code for specifying the reason of a settlement failure.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (f *FailedSettlementReason1FormatChoice) SetCode(value string) {
	f.Code = (*FailedSettlementReason1Code)(&value)
}

func (f *FailedSettlementReason1FormatChoice) AddProprietary() *GenericIdentification13 {
	f.Proprietary = new(GenericIdentification13)
	return f.Proprietary
}
