package internal


type Filter struct {
	Key string `json:"key"`

	// If value is empty, just return everything but sorted with the key
	Value string `json:"value,omitempty"`
}

type Wallet struct {
	Name        string  `json:"name" form:"name"`
	Type        string  `json:"type" form:"type"`
	Balance     float64 `json:"balance"`
	Slug        string  `json:"slug" form:"slug"`
	Description string  ` json:"description" form:"description"`
	UserId      uint
}
