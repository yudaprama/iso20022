package model

// Information related to the acknowledgement from an ATM to contact the ATM manager.
type HostToATMAcknowledgement1 struct {

	// Environment of the ATM.
	Environment *ATMEnvironment9 `xml:"Envt"`
}

func (h *HostToATMAcknowledgement1) AddEnvironment() *ATMEnvironment9 {
	h.Environment = new(ATMEnvironment9)
	return h.Environment
}
