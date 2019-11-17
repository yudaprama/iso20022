package model

// Content of the management plan.
type ManagementPlanContent1 struct {

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction1 `xml:"Actn"`
}

func (m *ManagementPlanContent1) AddAction() *TMSAction1 {
	newValue := new(TMSAction1)
	m.Action = append(m.Action, newValue)
	return newValue
}
