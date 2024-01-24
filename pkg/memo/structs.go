package memo

type CreateMemoRequest struct {
	Amount  int    `json:"amount" db:"amount"`
	Partner string `json:"partner" db:"partner"`
	Memo    string `json:"memo" db:"memo"`
	Date    string `json:"date" db:"date"`
	Period  string `json:"period" db:"period"`
	Type    int    `json:"type" db:"type"`
}

type GetMemoResponse []OneMemo

type OneMemo struct {
	ID      int    `json:"id" db:"id"`
	Amount  int    `json:"amount" db:"amount"`
	Partner string `json:"partner" db:"partner"`
	Memo    string `json:"memo" db:"memo"`
	Date    string `json:"date" db:"date"`
	Period  string `json:"period" db:"period"`
	Type    int    `json:"type" db:"type"`
}
