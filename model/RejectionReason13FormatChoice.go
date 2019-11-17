package model

// Choice of formats to  express the reason of a rejection of a movement.
type RejectionReason13FormatChoice struct {

	// Standard code to specify the reason of a rejection of a movement.
	Code *RejectionReason13Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of a movement.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason13FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason13Code)(&value)
}

func (r *RejectionReason13FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
