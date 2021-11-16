package Models

type User struct {
	Id    int    `json:"id"`
	Phone string `json:"phone" gorm:"size:15;unique"`
	Otp   string `json:"otp" gorm:"size:6"`
	Stock int    `json:"stock"`
	Setor int    `json:"stock"`

	Orders []Order `gorm:"ForeignKey:UserId"`
}
