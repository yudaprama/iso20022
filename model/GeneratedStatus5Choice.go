package model

// Choice of generated status.
type GeneratedStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the generated status.
	Reason []*GeneratedReason3 `xml:"Rsn"`
}

func (g *GeneratedStatus5Choice) SetNoSpecifiedReason(value string) {
	g.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (g *GeneratedStatus5Choice) AddReason() *GeneratedReason3 {
	newValue := new(GeneratedReason3)
	g.Reason = append(g.Reason, newValue)
	return newValue
}
