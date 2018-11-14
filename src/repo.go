package src

type Repository struct {
	repo []*Crawler
}

func NewRepo() *Repository {
	return &Repository{}
}

func (r *Repository) Add(c *Crawler) bool {
	return true
}

func (r *Repository) AddAll(r []MySqlRow) bool {
	// Loop thru array
	// 	Parse the row
	// 	Instantiate a crawler type
	// 	Insert into the repo
	return true
}

func (r *Repository) Delete(ID int) bool {
	return true
}

func (r *Repository) GetAll() []*Crawler {
	return r.repo
}
