package models

//hem sayfa hem de kullanıcı listesine erisebiliyor olacagız
type UserViewModel struct {
	Page  Page
	Users []User
}
