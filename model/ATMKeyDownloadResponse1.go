package model

// Information related to the response of an ATM key download from an ATM manager.
type ATMKeyDownloadResponse1 struct {

	// Environment of the key download.
	Environment *ATMEnvironment7 `xml:"Envt"`

	// Context of the ATM for the key download.
	ATMSecurityContext *ATMSecurityContext2 `xml:"ATMSctyCntxt"`

	// Random value from the ATM provided during a previous exchange.
	ATMChallenge *Max140Binary `xml:"ATMChllng,omitempty"`

	// Security parameters of the host downloading the key.
	HostSecurityParameters *SecurityParameters5 `xml:"HstSctyParams"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMKeyDownloadResponse1) AddEnvironment() *ATMEnvironment7 {
	a.Environment = new(ATMEnvironment7)
	return a.Environment
}

func (a *ATMKeyDownloadResponse1) AddATMSecurityContext() *ATMSecurityContext2 {
	a.ATMSecurityContext = new(ATMSecurityContext2)
	return a.ATMSecurityContext
}

func (a *ATMKeyDownloadResponse1) SetATMChallenge(value string) {
	a.ATMChallenge = (*Max140Binary)(&value)
}

func (a *ATMKeyDownloadResponse1) AddHostSecurityParameters() *SecurityParameters5 {
	a.HostSecurityParameters = new(SecurityParameters5)
	return a.HostSecurityParameters
}

func (a *ATMKeyDownloadResponse1) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}
