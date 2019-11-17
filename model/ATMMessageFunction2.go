package model

// Identifies the type of process related to an ATM message.
type ATMMessageFunction2 struct {

	// Type of requested function.
	Function *MessageFunction11Code `xml:"Fctn"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the ATM manager host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`
}

func (a *ATMMessageFunction2) SetFunction(value string) {
	a.Function = (*MessageFunction11Code)(&value)
}

func (a *ATMMessageFunction2) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMMessageFunction2) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}
