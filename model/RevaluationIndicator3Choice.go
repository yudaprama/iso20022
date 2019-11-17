package model

// Choice of format for the revaluation.
type RevaluationIndicator3Choice struct {

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Revaluation information provided as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RevaluationIndicator3Choice) SetIndicator(value string) {
	r.Indicator = (*YesNoIndicator)(&value)
}

func (r *RevaluationIndicator3Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
