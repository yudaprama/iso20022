package model

// Environment of the ATM.
type ATMEnvironment9 struct {

	// Acquirer of the ATM transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine7 `xml:"ATM"`
}

func (a *ATMEnvironment9) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment9) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment9) AddATM() *AutomatedTellerMachine7 {
	a.ATM = new(AutomatedTellerMachine7)
	return a.ATM
}
