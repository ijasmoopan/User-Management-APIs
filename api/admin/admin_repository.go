package admin

type Admin2 struct {
	Id        int `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gom:"index"`
	Password  string `json:"password"`
}