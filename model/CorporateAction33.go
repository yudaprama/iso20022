package model

// Provides information about the corporate action event.
type CorporateAction33 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate49 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat14Choice `xml:"EvtStag,omitempty"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat11Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat4Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction33) AddDateDetails() *CorporateActionDate49 {
	c.DateDetails = new(CorporateActionDate49)
	return c.DateDetails
}

func (c *CorporateAction33) AddEventStage() *CorporateActionEventStageFormat14Choice {
	c.EventStage = new(CorporateActionEventStageFormat14Choice)
	return c.EventStage
}

func (c *CorporateAction33) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat11Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat11Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateAction33) AddLotteryType() *LotteryTypeFormat4Choice {
	c.LotteryType = new(LotteryTypeFormat4Choice)
	return c.LotteryType
}
