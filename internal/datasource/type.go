package datasource

type User struct {
	ID        int      `json:"id" db:"id"`
	Name      string   `json:"name" db:"name"`
	Email     string   `json:"email" db:"email"`
	Pastebins []string `json:"pastebins,omitempty" db:"pastebins"`
}

type Pastebin struct {
	ID      int    `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
	UserID  int    `json:"user_id" db:"user_id"`
}
