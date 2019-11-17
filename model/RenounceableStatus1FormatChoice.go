package model

// Choice of formats to  express the renounceable status.
type RenounceableStatus1FormatChoice struct {

	// Standard code to specify the renounceable status.
	Code *RenounceableStatus1Code `xml:"Cd"`

	// Proprietary code to  express the renounceable status.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RenounceableStatus1FormatChoice) SetCode(value string) {
	r.Code = (*RenounceableStatus1Code)(&value)
}

func (r *RenounceableStatus1FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
