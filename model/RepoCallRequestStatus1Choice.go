package model

// Choice of format for the repurchase agreement call acknowledgement.
type RepoCallRequestStatus1Choice struct {

	// Provides the status of the repurchase agreement call request.
	Code *RepoCallRequestStatus1Code `xml:"Cd"`

	// Provides the status of the repurchase agreement call request.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepoCallRequestStatus1Choice) SetCode(value string) {
	r.Code = (*RepoCallRequestStatus1Code)(&value)
}

func (r *RepoCallRequestStatus1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
