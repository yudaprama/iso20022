package model

// Choice of cancellation processing status.
type CancellationProcessingStatus7Choice struct {

	// Provides the status of a cancellation request.
	Code *CancellationProcessingStatus1Code `xml:"Cd"`

	// Provides the status of a cancellation request.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CancellationProcessingStatus7Choice) SetCode(value string) {
	c.Code = (*CancellationProcessingStatus1Code)(&value)
}

func (c *CancellationProcessingStatus7Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
