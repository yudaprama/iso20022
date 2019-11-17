package model

// Content of the management plan.
type ManagementPlanContent2 struct {

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction2 `xml:"Actn"`
}

func (m *ManagementPlanContent2) AddAction() *TMSAction2 {
	newValue := new(TMSAction2)
	m.Action = append(m.Action, newValue)
	return newValue
}
