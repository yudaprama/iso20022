package model

// Choice of cancellation processing status.
type CancellationProcessingStatus8Choice struct {

	// Provides the status of a cancellation request.
	Code *CancellationProcessingStatus1Code `xml:"Cd"`

	// Provides the status of a cancellation request.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CancellationProcessingStatus8Choice) SetCode(value string) {
	c.Code = (*CancellationProcessingStatus1Code)(&value)
}

func (c *CancellationProcessingStatus8Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
