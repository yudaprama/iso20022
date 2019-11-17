package model

// Choice of format for the revaluation.
type RevaluationIndicator1Choice struct {

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Revaluation information provided as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RevaluationIndicator1Choice) SetIndicator(value string) {
	r.Indicator = (*YesNoIndicator)(&value)
}

func (r *RevaluationIndicator1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
