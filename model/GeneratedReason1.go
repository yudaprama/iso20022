package model

// The status of an instruction, advice or request.
type GeneratedReason1 struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReasons1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (g *GeneratedReason1) AddCode() *GeneratedReasons1Choice {
	g.Code = new(GeneratedReasons1Choice)
	return g.Code
}

func (g *GeneratedReason1) SetAdditionalReasonInformation(value string) {
	g.AdditionalReasonInformation = (*Max210Text)(&value)
}
