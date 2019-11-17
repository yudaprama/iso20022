package model

// Choice between a standard code or proprietary code to specify the type of the renounceable entitlement status.
type RenounceableEntitlementStatusTypeFormat1Choice struct {

	// Standard code to specify the renounceable status.
	Code *RenounceableStatus1Code `xml:"Cd"`

	// Proprietary identification of the renounceable status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RenounceableEntitlementStatusTypeFormat1Choice) SetCode(value string) {
	r.Code = (*RenounceableStatus1Code)(&value)
}

func (r *RenounceableEntitlementStatusTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
