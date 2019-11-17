package model

// Local time offset to UTC (Coordinated Universal Time).
type LocalDateTime1 struct {

	// Date time of the beginning of the period (inclusive).
	FromDateTime *ISODateTime `xml:"FrDtTm,omitempty"`

	// Date time of the end of the period (exclusive).
	ToDateTime *ISODateTime `xml:"ToDtTm,omitempty"`

	// UTC offset in minutes, of the local time during the period. For instance, 120 for Central European Time, -720 for Central Standard Time (North America).
	UTCOffset *Number `xml:"UTCOffset"`
}

func (l *LocalDateTime1) SetFromDateTime(value string) {
	l.FromDateTime = (*ISODateTime)(&value)
}

func (l *LocalDateTime1) SetToDateTime(value string) {
	l.ToDateTime = (*ISODateTime)(&value)
}

func (l *LocalDateTime1) SetUTCOffset(value string) {
	l.UTCOffset = (*Number)(&value)
}
