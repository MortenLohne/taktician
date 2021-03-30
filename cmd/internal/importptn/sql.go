package importptn

const createPTNTable = `
CREATE TABLE IF NOT EXISTS ptns (
  id integer primary key,
  ptn string
)
`

/*
CREATE TABLE games (
  id INTEGER PRIMARY KEY,
  date INT,
  size INT,
  player_white VARCHAR(20),
  player_black VARCHAR(20),
  notation TEXT,
  result VARCHAR(10),
  timertime INT DEFAULT 0,
  timerinc INT DEFAULT 0
);
*/

type gameRow struct {
	Id   int `db:"id"`
	Date int `db:"date"`
	Size int `db:"size"`

	PlayerWhite string `db:"player_white"`
	PlayerBlack string `db:"player_black"`
	RatingWhite string `db:"rating_white"`
	RatingBlack string `db:"rating_black"`
	RatingChangeWhite int `db:"rating_change_white"`
	RatingChangeBlack int `db:"rating_change_black"`

	IsRated int `db:"unrated"`
	Tournament string `db:"tournament"`

	Notation string `db:"notation"`
	Result   string `db:"result"`

	Komi int `db:"komi"`
	NumPieces int `db:"pieces"`
	NumCapstones int `db:"capstones"`
	TimerTime int `db:"timertime"`
	TimerInc  int `db:"timerinc"`
}

type ptnRow struct {
	Id  int    `db:"id"`
	PTN string `db:"ptn"`
}

const selectTODO = `
SELECT g.*
FROM games g
`
