package model

// Information related to the request of a key download from an ATM.
type ATMKeyDownloadRequest1 struct {

	// Environment of the key download.
	Environment *ATMEnvironment8 `xml:"Envt"`

	// Result of a maintenance command performed by the ATM.
	CommandResult []*ATMCommand2 `xml:"CmdRslt,omitempty"`

	// Security command in progress inside which the key download is requested.
	CommandContext *ATMCommand3 `xml:"CmdCntxt,omitempty"`

	// Context of the ATM for the key download.
	ATMSecurityContext *ATMSecurityContext2 `xml:"ATMSctyCntxt"`

	// Security parameters of the ATM for the initiated key download.
	ATMSecurityParameters *SecurityParameters4 `xml:"ATMSctyParams"`

	// Random value from the host provided during a previous exchange.
	HostChallenge *Max140Binary `xml:"HstChllng,omitempty"`
}

func (a *ATMKeyDownloadRequest1) AddEnvironment() *ATMEnvironment8 {
	a.Environment = new(ATMEnvironment8)
	return a.Environment
}

func (a *ATMKeyDownloadRequest1) AddCommandResult() *ATMCommand2 {
	newValue := new(ATMCommand2)
	a.CommandResult = append(a.CommandResult, newValue)
	return newValue
}

func (a *ATMKeyDownloadRequest1) AddCommandContext() *ATMCommand3 {
	a.CommandContext = new(ATMCommand3)
	return a.CommandContext
}

func (a *ATMKeyDownloadRequest1) AddATMSecurityContext() *ATMSecurityContext2 {
	a.ATMSecurityContext = new(ATMSecurityContext2)
	return a.ATMSecurityContext
}

func (a *ATMKeyDownloadRequest1) AddATMSecurityParameters() *SecurityParameters4 {
	a.ATMSecurityParameters = new(SecurityParameters4)
	return a.ATMSecurityParameters
}

func (a *ATMKeyDownloadRequest1) SetHostChallenge(value string) {
	a.HostChallenge = (*Max140Binary)(&value)
}
