package entity

type Attachable struct {
	ID              uint    `json:"id"`
	UUID            string  `json:"uuid"`
	Attach          *Attach `json:"attach"`
	AttachId        uint    `json:"attachId"`
	AttachableType  string  `json:"attachable_type"`
	AttachableField string  `json:"attachable_field"`
	AttachableID    string  `json:"attachable_id"`
}
