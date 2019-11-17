package model

// Specifies the rate as a percentage or a text.
type RateType4Choice struct {

	// Ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	Percentage *PercentageRate `xml:"Pctg"`

	// Rate type expressed, in an other form.
	Other *Max35Text `xml:"Othr"`
}

func (r *RateType4Choice) SetPercentage(value string) {
	r.Percentage = (*PercentageRate)(&value)
}

func (r *RateType4Choice) SetOther(value string) {
	r.Other = (*Max35Text)(&value)
}
