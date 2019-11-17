package model

// Choice of generated status.
type GeneratedStatus3Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the GeneratedStatus.
	Reason []*GeneratedReason1 `xml:"Rsn"`
}

func (g *GeneratedStatus3Choice) SetNoSpecifiedReason(value string) {
	g.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (g *GeneratedStatus3Choice) AddReason() *GeneratedReason1 {
	newValue := new(GeneratedReason1)
	g.Reason = append(g.Reason, newValue)
	return newValue
}
