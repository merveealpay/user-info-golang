package models

//hem sayfa hem de kullanıcı listesine erisebiliyor olacagız
type UserViewMOdel struct {
	Page  Page
	Users []User
}
