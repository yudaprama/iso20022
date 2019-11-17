package model

// Manufacturer configuration parameters of the point of interaction (POI).
type PaymentTerminalParameters3 struct {

	// Type of action for the configuration parameters.
	ActionType *TerminalManagementAction3Code `xml:"ActnTp"`

	// Identification of the vendor for the MTM, if the POI manages various subsets of terminal parameters.
	VendorIdentification *Max35Text `xml:"VndrId,omitempty"`

	// Version of the terminal parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Parameters to synchronise the real time clock of the POI (Point Of Interaction).
	ClockSynchronisation *ClockSynchronisation1 `xml:"ClckSynctn,omitempty"`

	// Time zone line to update in the time zone data base subset stored in the POI (Point Of Interaction). The format of the line is conform to the IANA (Internet Assigned Number Authority) time zone data base.
	TimeZoneLine []*Max70Text `xml:"TmZoneLine,omitempty"`

	// Local time offset to UTC (Coordinated Universal Time).
	LocalDateTime []*LocalDateTime1 `xml:"LclDtTm,omitempty"`

	// Others manufacturer configuration parameters of the point of interaction.
	OtherParameters *Max10000Binary `xml:"OthrParams,omitempty"`
}

func (p *PaymentTerminalParameters3) SetActionType(value string) {
	p.ActionType = (*TerminalManagementAction3Code)(&value)
}

func (p *PaymentTerminalParameters3) SetVendorIdentification(value string) {
	p.VendorIdentification = (*Max35Text)(&value)
}

func (p *PaymentTerminalParameters3) SetVersion(value string) {
	p.Version = (*Max256Text)(&value)
}

func (p *PaymentTerminalParameters3) AddClockSynchronisation() *ClockSynchronisation1 {
	p.ClockSynchronisation = new(ClockSynchronisation1)
	return p.ClockSynchronisation
}

func (p *PaymentTerminalParameters3) AddTimeZoneLine(value string) {
	p.TimeZoneLine = append(p.TimeZoneLine, (*Max70Text)(&value))
}

func (p *PaymentTerminalParameters3) AddLocalDateTime() *LocalDateTime1 {
	newValue := new(LocalDateTime1)
	p.LocalDateTime = append(p.LocalDateTime, newValue)
	return newValue
}

func (p *PaymentTerminalParameters3) SetOtherParameters(value string) {
	p.OtherParameters = (*Max10000Binary)(&value)
}
