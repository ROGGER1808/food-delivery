package common

type SimpleCategory struct {
	SQLModel `json:",inline"`
	Name     string `json:"name"`
}

func (SimpleCategory) TableName() string {
	return "categories"
}
