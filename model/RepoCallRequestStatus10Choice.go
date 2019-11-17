package model

// Choice of format for the repurchase agreement call acknowledgement.
type RepoCallRequestStatus10Choice struct {

	// Provides the status of the repurchase agreement call request.
	Code *RepoCallRequestStatus1Code `xml:"Cd"`

	// Provides the status of the repurchase agreement call request.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RepoCallRequestStatus10Choice) SetCode(value string) {
	r.Code = (*RepoCallRequestStatus1Code)(&value)
}

func (r *RepoCallRequestStatus10Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
