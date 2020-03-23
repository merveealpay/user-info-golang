package handlers

import (
	"encoding/json"
	"net/http"

	. "../dataloaders"
	. "../models"
)

//main.go icerisinde run func cagirmak icin
func Run() {
	http.HandleFunc("/", Handler)

	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//page nesnesi
	page := Page{ID: 7, Name: "Kullanicilar", Description: "Kullanici Listesi", URI: "/users"}
	//dataloaders
	users := LoadUsers()
	interests := LoadInterests()
	interestsMappings := LoadInterestMappings()

	//interest ve user'ları eslestiriyoruz
	var newUsers []User
	for _, user := range users {
		for _, interestsMapping := range interestsMappings {
			if user.ID == interestsMapping.UserID {
				for _, interest := range interests {
					if interestsMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}

		newUsers = append(newUsers, user)

	}

	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))

}
