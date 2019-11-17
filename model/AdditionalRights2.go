package model

// Information about the general meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
type AdditionalRights2 struct {

	// Specific rights granted to the shareholders that can be exercised at shareholders meetings, for example, the right to ask questions, the right to add items to the agenda or table draft resolutions.
	AdditionalRight *AdditionalRightCode1Choice `xml:"AddtlRght"`

	// Address to use over the www (HTTP) service where additional information on specific rights granted to the shareholders can be found.
	AdditionalRightInformationURLAddress *Max256Text `xml:"AddtlRghtInfURLAdr,omitempty"`

	// Additional right granted to determine the date and time by which security holders can propose amendments or new resolutions. This deadline is set by an intermediary.
	AdditionalRightDeadline *DateFormat29Choice `xml:"AddtlRghtDdln,omitempty"`

	// Additional right granted to determine the date and time by which security holders can propose amendments or new resolutions. This deadline is set by the issuer.
	AdditionalRightMarketDeadline *DateFormat29Choice `xml:"AddtlRghtMktDdln,omitempty"`

	// Additional right granted to specify the minimum stake in share capital or cash value or number of security holders required to table resolutions.
	AdditionalRightThreshold *AdditionalRightThreshold1Choice `xml:"AddtlRghtThrshld,omitempty"`
}

func (a *AdditionalRights2) AddAdditionalRight() *AdditionalRightCode1Choice {
	a.AdditionalRight = new(AdditionalRightCode1Choice)
	return a.AdditionalRight
}

func (a *AdditionalRights2) SetAdditionalRightInformationURLAddress(value string) {
	a.AdditionalRightInformationURLAddress = (*Max256Text)(&value)
}

func (a *AdditionalRights2) AddAdditionalRightDeadline() *DateFormat29Choice {
	a.AdditionalRightDeadline = new(DateFormat29Choice)
	return a.AdditionalRightDeadline
}

func (a *AdditionalRights2) AddAdditionalRightMarketDeadline() *DateFormat29Choice {
	a.AdditionalRightMarketDeadline = new(DateFormat29Choice)
	return a.AdditionalRightMarketDeadline
}

func (a *AdditionalRights2) AddAdditionalRightThreshold() *AdditionalRightThreshold1Choice {
	a.AdditionalRightThreshold = new(AdditionalRightThreshold1Choice)
	return a.AdditionalRightThreshold
}
