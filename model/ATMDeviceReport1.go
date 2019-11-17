package model

// Information related to the report from an ATM device.
type ATMDeviceReport1 struct {

	// Environment of the transaction.
	Environment *ATMEnvironment6 `xml:"Envt"`

	// Global status of the ATM.
	ATMGlobalStatus *ATMStatus1 `xml:"ATMGblSts"`

	// Result of a maintenance command performed by the ATM.
	CommandResult []*ATMCommand2 `xml:"CmdRslt,omitempty"`

	// Maintenance command which has requested the device report.
	CommandContext *ATMCommand3 `xml:"CmdCntxt,omitempty"`

	// Information related to security commands.
	ATMSecurityContext *ATMSecurityContext1 `xml:"ATMSctyCntxt,omitempty"`
}

func (a *ATMDeviceReport1) AddEnvironment() *ATMEnvironment6 {
	a.Environment = new(ATMEnvironment6)
	return a.Environment
}

func (a *ATMDeviceReport1) AddATMGlobalStatus() *ATMStatus1 {
	a.ATMGlobalStatus = new(ATMStatus1)
	return a.ATMGlobalStatus
}

func (a *ATMDeviceReport1) AddCommandResult() *ATMCommand2 {
	newValue := new(ATMCommand2)
	a.CommandResult = append(a.CommandResult, newValue)
	return newValue
}

func (a *ATMDeviceReport1) AddCommandContext() *ATMCommand3 {
	a.CommandContext = new(ATMCommand3)
	return a.CommandContext
}

func (a *ATMDeviceReport1) AddATMSecurityContext() *ATMSecurityContext1 {
	a.ATMSecurityContext = new(ATMSecurityContext1)
	return a.ATMSecurityContext
}
