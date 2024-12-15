package requests

type UserSegmentPair struct {
	UserId  string `json:"user_id" validate:"required,max=255,min=3"`
	Segment string `json:"segment" validate:"required,max=255,min=3"`
}

type SegmentRouteParam struct {
	Segment string `json:"segment" validate:"required,max=255,min=3"`
}
