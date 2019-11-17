package model

// Rate is expressed as a percentage or a text.
type RateTypeChoice struct {

	// Percentage charged for the use of an amount of money, usually expressed at an annual rate. The interest rate is the ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	PercentageRate *PercentageRate `xml:"PctgRate"`

	// Rate is expressed as a text.
	TextualRate *Max35Text `xml:"TxtlRate"`
}

func (r *RateTypeChoice) SetPercentageRate(value string) {
	r.PercentageRate = (*PercentageRate)(&value)
}

func (r *RateTypeChoice) SetTextualRate(value string) {
	r.TextualRate = (*Max35Text)(&value)
}
