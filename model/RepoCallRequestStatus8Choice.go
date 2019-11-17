package model

// Choice of format for the repurchase agreement call acknowledgement.
type RepoCallRequestStatus8Choice struct {

	// Provides the status of the repurchase agreement call request.
	Code *RepoCallRequestStatus1Code `xml:"Cd"`

	// Provides the status of the repurchase agreement call request.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepoCallRequestStatus8Choice) SetCode(value string) {
	r.Code = (*RepoCallRequestStatus1Code)(&value)
}

func (r *RepoCallRequestStatus8Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
