package model

// Choice of generated status.
type GeneratedStatus7Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the generated status.
	Reason []*GeneratedReason5 `xml:"Rsn"`
}

func (g *GeneratedStatus7Choice) SetNoSpecifiedReason(value string) {
	g.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (g *GeneratedStatus7Choice) AddReason() *GeneratedReason5 {
	newValue := new(GeneratedReason5)
	g.Reason = append(g.Reason, newValue)
	return newValue
}
