package models

//	"sort"
//	"strconv"
//	"strings"

type Article struct {
	Id         int
	Title      string
	Content    string
	Views      int
	Favourites int
}

func (m *Article) SetId(Id int) {
	m.Id = Id
}
func (m *Article) GetId() int {
	return m.Id
}
func (m *Article) SetTitle(Title string) {
	m.Title = Title
}
func (m *Article) GetTitle() string {
	return m.Title
}
func (m *Article) SetContent(Content string) {
	m.Content = Content
}
func (m *Article) GetContent() string {
	return m.Content
}
func (m *Article) SetViews(Views int) {
	m.Views = Views
}
func (m *Article) GetViews() int {
	return m.Views
}
func (m *Article) SetFavourites(Favourites int) {
	m.Favourites = Favourites
}
func (m *Article) GetFavourites() int {
	return m.Favourites
}
func ToStringDictionary(m *Article) map[string]interface{} {
	ArtMap := make(map[string]interface{}, 0)
	ArtMap["Id"] = m.Id
	ArtMap["Title"] = m.Title
	ArtMap["Content"] = m.Content
	ArtMap["Views"] = m.Views
	ArtMap["Favourites"] = m.Favourites
	return ArtMap
}
