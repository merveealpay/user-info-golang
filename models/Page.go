//sayfalar arası da veri göndermek istiyoruz bu yüzden page nesnesine ihtiyac duyuyoruz

package models

type Page struct {
	ID          int
	Name        string
	Description string
	URI         string
}
