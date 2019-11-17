package model

// Status of an account.
type OtherAccountStatus1 struct {

	// Status of the account.
	Status *GenericIdentification36 `xml:"Sts"`

	// Reason for the status of the account.
	Reason *GenericIdentification36 `xml:"Rsn,omitempty"`
}

func (o *OtherAccountStatus1) AddStatus() *GenericIdentification36 {
	o.Status = new(GenericIdentification36)
	return o.Status
}

func (o *OtherAccountStatus1) AddReason() *GenericIdentification36 {
	o.Reason = new(GenericIdentification36)
	return o.Reason
}
