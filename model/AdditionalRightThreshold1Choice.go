package model

// Choice of additional right threshold.
type AdditionalRightThreshold1Choice struct {

	// Additional right granted to specify the minimum stake in share capital or cash value or number of security holders required to table resolutions.
	AdditionalRightThreshold *Max35Text `xml:"AddtlRghtThrshld"`

	// Additional right granted to specify the minimum stake in share capital or cash value or number of security holders required to table resolutions. This minimum is expressed as a percentage.
	AdditionalRightThresholdPercentage *PercentageRate `xml:"AddtlRghtThrshldPctg"`
}

func (a *AdditionalRightThreshold1Choice) SetAdditionalRightThreshold(value string) {
	a.AdditionalRightThreshold = (*Max35Text)(&value)
}

func (a *AdditionalRightThreshold1Choice) SetAdditionalRightThresholdPercentage(value string) {
	a.AdditionalRightThresholdPercentage = (*PercentageRate)(&value)
}
