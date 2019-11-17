package model

// The status of an instruction, advice or request.
type GeneratedReason3 struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReasons3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (g *GeneratedReason3) AddCode() *GeneratedReasons3Choice {
	g.Code = new(GeneratedReasons3Choice)
	return g.Code
}

func (g *GeneratedReason3) SetAdditionalReasonInformation(value string) {
	g.AdditionalReasonInformation = (*Max210Text)(&value)
}
