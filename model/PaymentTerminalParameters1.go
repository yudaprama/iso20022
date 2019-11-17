package model

// Manufacturer configuration parameters of the point of interaction (POI).
type PaymentTerminalParameters1 struct {

	// Parameters to synchronise the real time clock of the POI (Point Of Interaction).
	ClockSynchronisation *ClockSynchronisation1 `xml:"ClckSynctn,omitempty"`

	// Time zone line to update in the time zone data base subset stored in the POI (Point Of Interaction). The format of the line is conform to the IANA (Internet Assigned Number Authority) time zone data base.
	TimeZoneLine []*Max70Text `xml:"TmZoneLine,omitempty"`

	// Others manufacturer configuration parameters of the point of interaction.
	OtherParameters *Max10000Binary `xml:"OthrParams,omitempty"`
}

func (p *PaymentTerminalParameters1) AddClockSynchronisation() *ClockSynchronisation1 {
	p.ClockSynchronisation = new(ClockSynchronisation1)
	return p.ClockSynchronisation
}

func (p *PaymentTerminalParameters1) AddTimeZoneLine(value string) {
	p.TimeZoneLine = append(p.TimeZoneLine, (*Max70Text)(&value))
}

func (p *PaymentTerminalParameters1) SetOtherParameters(value string) {
	p.OtherParameters = (*Max10000Binary)(&value)
}
