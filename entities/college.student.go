package entities

type CollegeStudent struct {
	Id         int64  `json:"id"`
	FullName   string `json:"full_name"`
	Gender     string `json:"gender"`
	BirthPlace string `json:"birth_place"`
	BirthDay   string `json:"birth_day"`
	Address    string `json:"address"`
}
