package model

// Choice between a standard code or proprietary code to specify the type of the renounceable entitlement status.
type RenounceableEntitlementStatusTypeFormat4Choice struct {

	// Standard code to specify the renounceable status.
	Code *RenounceableStatus1Code `xml:"Cd"`

	// Proprietary identification of the renounceable status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RenounceableEntitlementStatusTypeFormat4Choice) SetCode(value string) {
	r.Code = (*RenounceableStatus1Code)(&value)
}

func (r *RenounceableEntitlementStatusTypeFormat4Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
