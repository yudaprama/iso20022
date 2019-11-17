package model

// Information related to the request to an ATM to contact the ATM manager.
type HostToATMRequest1 struct {

	// Environment of the ATM.
	Environment *ATMEnvironment9 `xml:"Envt"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Message that have to be sent by the ATM.
	ExpectedMessageFunction *MessageFunction8Code `xml:"XpctdMsgFctn"`
}

func (h *HostToATMRequest1) AddEnvironment() *ATMEnvironment9 {
	h.Environment = new(ATMEnvironment9)
	return h.Environment
}

func (h *HostToATMRequest1) AddCommandIdentification() *ATMCommandIdentification1 {
	h.CommandIdentification = new(ATMCommandIdentification1)
	return h.CommandIdentification
}

func (h *HostToATMRequest1) SetExpectedMessageFunction(value string) {
	h.ExpectedMessageFunction = (*MessageFunction8Code)(&value)
}
