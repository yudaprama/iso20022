package model

// Detailed description of an error.
type ErrorReporting1 struct {

	// Type of error.
	Type *Max35Text `xml:"Tp"`

	// Detailed description of the error.
	Description *Max500Text `xml:"Desc"`
}

func (e *ErrorReporting1) SetType(value string) {
	e.Type = (*Max35Text)(&value)
}

func (e *ErrorReporting1) SetDescription(value string) {
	e.Description = (*Max500Text)(&value)
}
