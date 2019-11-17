package model

// Specifies the number of occurrences of a particular event and the maximum number of times this event may occur.
type Limit1 struct {

	// Number of occurrences of a particular event.
	Current *Max3NumericText `xml:"Cur"`

	// Specifies the maximum number of times an event may occur.
	Limit *Max3NumericText `xml:"Lmt"`
}

func (l *Limit1) SetCurrent(value string) {
	l.Current = (*Max3NumericText)(&value)
}

func (l *Limit1) SetLimit(value string) {
	l.Limit = (*Max3NumericText)(&value)
}
