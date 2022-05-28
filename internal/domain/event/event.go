package event

type Poster struct {
	ID          int64  `db:"id,omitempty"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Year        int64  `db:"year"`
}
