package model

// Information related to the response of an ATM key download from an ATM manager.
type ATMKeyDownloadResponse2 struct {

	// Environment of the key download.
	Environment *ATMEnvironment7 `xml:"Envt"`

	// Context of the ATM for the key download.
	ATMSecurityContext *ATMSecurityContext2 `xml:"ATMSctyCntxt"`

	// Random value from the ATM provided during a previous exchange.
	ATMChallenge *Max140Binary `xml:"ATMChllng,omitempty"`

	// Security parameters of the host downloading the key.
	HostSecurityParameters *SecurityParameters5 `xml:"HstSctyParams"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMKeyDownloadResponse2) AddEnvironment() *ATMEnvironment7 {
	a.Environment = new(ATMEnvironment7)
	return a.Environment
}

func (a *ATMKeyDownloadResponse2) AddATMSecurityContext() *ATMSecurityContext2 {
	a.ATMSecurityContext = new(ATMSecurityContext2)
	return a.ATMSecurityContext
}

func (a *ATMKeyDownloadResponse2) SetATMChallenge(value string) {
	a.ATMChallenge = (*Max140Binary)(&value)
}

func (a *ATMKeyDownloadResponse2) AddHostSecurityParameters() *SecurityParameters5 {
	a.HostSecurityParameters = new(SecurityParameters5)
	return a.HostSecurityParameters
}

func (a *ATMKeyDownloadResponse2) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}
