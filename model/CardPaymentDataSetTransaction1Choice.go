package model

// Card payment transaction choice between cancellation, authorisation request and authorisation response.
type CardPaymentDataSetTransaction1Choice struct {

	// Completed card payment transaction to be captured.
	Completion *CardPaymentDataSetTransaction2 `xml:"Cmpltn,omitempty"`

	// Cancelled card payment transaction to be captured.
	Cancellation *CardPaymentDataSetTransaction3 `xml:"Cxl,omitempty"`

	// Card payment transaction including an authorisation request.
	AuthorisationRequest *CardPaymentDataSetTransaction4 `xml:"AuthstnReq,omitempty"`

	// Card payment transaction including an authorisation response.
	AuthorisationResponse *CardPaymentDataSetTransaction5 `xml:"AuthstnRspn,omitempty"`
}

func (c *CardPaymentDataSetTransaction1Choice) AddCompletion() *CardPaymentDataSetTransaction2 {
	c.Completion = new(CardPaymentDataSetTransaction2)
	return c.Completion
}

func (c *CardPaymentDataSetTransaction1Choice) AddCancellation() *CardPaymentDataSetTransaction3 {
	c.Cancellation = new(CardPaymentDataSetTransaction3)
	return c.Cancellation
}

func (c *CardPaymentDataSetTransaction1Choice) AddAuthorisationRequest() *CardPaymentDataSetTransaction4 {
	c.AuthorisationRequest = new(CardPaymentDataSetTransaction4)
	return c.AuthorisationRequest
}

func (c *CardPaymentDataSetTransaction1Choice) AddAuthorisationResponse() *CardPaymentDataSetTransaction5 {
	c.AuthorisationResponse = new(CardPaymentDataSetTransaction5)
	return c.AuthorisationResponse
}
