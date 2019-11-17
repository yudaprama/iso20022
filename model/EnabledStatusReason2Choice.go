package model

// Choice of formats for an enabled reason code.
type EnabledStatusReason2Choice struct {

	// Reason for the enabled account status expressed as a code.
	Code *EnabledStatusReason1Code `xml:"Cd"`

	// Reason for the enabled account status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (e *EnabledStatusReason2Choice) SetCode(value string) {
	e.Code = (*EnabledStatusReason1Code)(&value)
}

func (e *EnabledStatusReason2Choice) AddProprietary() *GenericIdentification36 {
	e.Proprietary = new(GenericIdentification36)
	return e.Proprietary
}
