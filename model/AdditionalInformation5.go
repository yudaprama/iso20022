package model

// Contains additional information related to the message.
type AdditionalInformation5 struct {

	// Contains additional information related to the message.
	Information []*Max256Text `xml:"Inf"`
}

func (a *AdditionalInformation5) AddInformation(value string) {
	a.Information = append(a.Information, (*Max256Text)(&value))
}
