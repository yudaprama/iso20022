package model

// Choice of format for the response status.
type ResponseStatus7Choice struct {

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	Code *ResponseStatus1Code `xml:"Cd"`

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *ResponseStatus7Choice) SetCode(value string) {
	r.Code = (*ResponseStatus1Code)(&value)
}

func (r *ResponseStatus7Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
