package model

// Choice of cancellation processing status.
type CancellationProcessingStatus1Choice struct {

	// Provides the status of a cancellation request.
	Code *CancellationProcessingStatus1Code `xml:"Cd"`

	// Provides the status of a cancellation request.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CancellationProcessingStatus1Choice) SetCode(value string) {
	c.Code = (*CancellationProcessingStatus1Code)(&value)
}

func (c *CancellationProcessingStatus1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
