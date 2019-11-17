package model

// Choice of format for the revaluation.
type RevaluationIndicator4Choice struct {

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Revaluation information provided as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RevaluationIndicator4Choice) SetIndicator(value string) {
	r.Indicator = (*YesNoIndicator)(&value)
}

func (r *RevaluationIndicator4Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
