package dataloaders

//uygulama ayaga kalktıgında bunları yukle, herhangi bir repodan da models kullanabiliriz mesela abc.com/repo/abcproject/models
//util i kısa isim olarak adlandırdık
import (
	"encoding/json"

	. "../models"
	util "../utils"
)

func LoadUsers() []User {
	bytes, _ := util.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterests() []Interest {
	bytes, _ := util.ReadFile("../json/interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}

func LoadInterestMappings() []InterestMapping {
	bytes, _ := util.ReadFile("../json/userInterestMappings.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
