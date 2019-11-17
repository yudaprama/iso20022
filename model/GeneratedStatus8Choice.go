package model

// Choice of generated status.
type GeneratedStatus8Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the generated status.
	Reason []*GeneratedReason6 `xml:"Rsn"`
}

func (g *GeneratedStatus8Choice) SetNoSpecifiedReason(value string) {
	g.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (g *GeneratedStatus8Choice) AddReason() *GeneratedReason6 {
	newValue := new(GeneratedReason6)
	g.Reason = append(g.Reason, newValue)
	return newValue
}
