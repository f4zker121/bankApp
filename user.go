package bankapp

type User struct {
	Id      int     `json:"-" db:"id"`
	Balance float64 `json:"balance"`
}
