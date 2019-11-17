package model

// Choice of format for the transfer reason.
type TransferReason1 struct {

	// Transfer reason expressed as an ISO 20022 code.
	Code *TransferReason1Code `xml:"Cd"`

	// Transfer reason expressed as a proprietary code.
	Proprietary *GenericIdentification27 `xml:"Prtry"`
}

func (t *TransferReason1) SetCode(value string) {
	t.Code = (*TransferReason1Code)(&value)
}

func (t *TransferReason1) AddProprietary() *GenericIdentification27 {
	t.Proprietary = new(GenericIdentification27)
	return t.Proprietary
}
