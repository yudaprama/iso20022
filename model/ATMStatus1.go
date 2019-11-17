package model

// Global status of the ATM.
type ATMStatus1 struct {

	// Actual status of the ATM.
	CurrentStatus *ATMStatus1Code `xml:"CurSts"`

	// Present if the status required by the ATM manager is different from the current status.
	DemandedStatus *ATMStatus1Code `xml:"DmnddSts,omitempty"`
}

func (a *ATMStatus1) SetCurrentStatus(value string) {
	a.CurrentStatus = (*ATMStatus1Code)(&value)
}

func (a *ATMStatus1) SetDemandedStatus(value string) {
	a.DemandedStatus = (*ATMStatus1Code)(&value)
}
