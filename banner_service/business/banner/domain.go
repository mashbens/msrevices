package banner

type Result struct {
	TotalAllData int
	CurrentPage  int
	PerPage      int
	TotalPage    int
	Domain       []Domain
}
type Domain struct {
	Id          int
	Name        string
	File        string
	Status      int
	Position    int
	File_mobile string
}
