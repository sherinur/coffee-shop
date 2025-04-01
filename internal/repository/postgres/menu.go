package postgres

import "database/sql"

type Menu struct {
	conn  *sql.DB
	table string
}

const (
	tableMenu = "menu"
)

func NewMenu(conn *sql.DB) *Menu {
	return &Menu{
		conn:  conn,
		table: tableMenu,
	}
}
