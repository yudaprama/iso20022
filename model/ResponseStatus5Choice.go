package model

// Choice of format for the response status.
type ResponseStatus5Choice struct {

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	Code *ResponseStatus1Code `xml:"Cd"`

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *ResponseStatus5Choice) SetCode(value string) {
	r.Code = (*ResponseStatus1Code)(&value)
}

func (r *ResponseStatus5Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
