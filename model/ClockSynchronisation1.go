package model

// Parameters to synchronise a real time clock.
type ClockSynchronisation1 struct {

	// Name of the time zone where is located the POI (Point Of Interaction), as definined by the IANA (Internet Assigned Number Authority) time zone data base.
	POITimeZone *Max70Text `xml:"POITmZone"`

	// Parameters to contact a time server.
	SynchronisationServer []*NetworkParameters2 `xml:"SynctnSvr,omitempty"`
}

func (c *ClockSynchronisation1) SetPOITimeZone(value string) {
	c.POITimeZone = (*Max70Text)(&value)
}

func (c *ClockSynchronisation1) AddSynchronisationServer() *NetworkParameters2 {
	newValue := new(NetworkParameters2)
	c.SynchronisationServer = append(c.SynchronisationServer, newValue)
	return newValue
}
