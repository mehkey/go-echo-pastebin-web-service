package datasource

type DB interface {
	//GET Services
	GetAllPastebins() ([]Pastebin, error)
	GetAllUsers() ([]User, error)
	//GetPastebinByID(int) (*Pastebin, error)
	GetUserByID(int) (*User, error)
	GetPastebinsForUser(int) ([]Pastebin, error)

	//POST Services
	CreateNewUser(*User) (int, error)
	//AddUserPastebin(int, *Pastebin) (int, error)
}
