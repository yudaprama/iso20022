package model

// Choice between a standard code or proprietary code to specify the type of the renounceable entitlement status.
type RenounceableEntitlementStatusTypeFormat3Choice struct {

	// Standard code to specify the renounceable status.
	Code *RenounceableStatus1Code `xml:"Cd"`

	// Proprietary identification of the renounceable status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RenounceableEntitlementStatusTypeFormat3Choice) SetCode(value string) {
	r.Code = (*RenounceableStatus1Code)(&value)
}

func (r *RenounceableEntitlementStatusTypeFormat3Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
