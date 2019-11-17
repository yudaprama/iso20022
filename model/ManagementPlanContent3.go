package model

// Content of the management plan.
type ManagementPlanContent3 struct {

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction3 `xml:"Actn"`
}

func (m *ManagementPlanContent3) AddAction() *TMSAction3 {
	newValue := new(TMSAction3)
	m.Action = append(m.Action, newValue)
	return newValue
}
