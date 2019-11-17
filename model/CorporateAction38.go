package model

// Provides information about the corporate action event.
type CorporateAction38 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate54 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat15Choice `xml:"EvtStag,omitempty"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat13Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat5Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction38) AddDateDetails() *CorporateActionDate54 {
	c.DateDetails = new(CorporateActionDate54)
	return c.DateDetails
}

func (c *CorporateAction38) AddEventStage() *CorporateActionEventStageFormat15Choice {
	c.EventStage = new(CorporateActionEventStageFormat15Choice)
	return c.EventStage
}

func (c *CorporateAction38) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat13Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat13Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateAction38) AddLotteryType() *LotteryTypeFormat5Choice {
	c.LotteryType = new(LotteryTypeFormat5Choice)
	return c.LotteryType
}
