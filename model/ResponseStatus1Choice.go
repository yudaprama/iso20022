package model

// Choice of format for the response status.
type ResponseStatus1Choice struct {

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	Code *ResponseStatus1Code `xml:"Cd"`

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *ResponseStatus1Choice) SetCode(value string) {
	r.Code = (*ResponseStatus1Code)(&value)
}

func (r *ResponseStatus1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
