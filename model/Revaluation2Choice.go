package model

// Choice between an indicator or a data source scheme to determine the revaluation.
type Revaluation2Choice struct {

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Reevaluation is determined using a data source scheme.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *Revaluation2Choice) SetIndicator(value string) {
	r.Indicator = (*YesNoIndicator)(&value)
}

func (r *Revaluation2Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
