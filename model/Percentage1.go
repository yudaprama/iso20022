package model

// Rate information.
type Percentage1 struct {

	// Percentage of an amount.
	Rate *PercentageRate `xml:"Rate"`

	// Indication of what the percentage is relative to.
	RelativeTo *ExternalRelativeTo1Code `xml:"RltvTo"`
}

func (p *Percentage1) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}

func (p *Percentage1) SetRelativeTo(value string) {
	p.RelativeTo = (*ExternalRelativeTo1Code)(&value)
}
