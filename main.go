package main

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"fmt"
)

type DBPaper struct {
	ID sql.NullInt64 `db:"id"`

	Title sql.NullString `db:"title"`
	Year  sql.NullInt64  `db:"year"`

	Authors []string
}

type AuthorsAndPapers struct {
	PaperID  int64 `db:"paper_id"`
	AuthorID int64 `db:"author_id"`
}

type Author struct {
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Middle    sql.NullString `db:"middle"`
}

func main() {
	db, err := sqlx.Connect("mysql", "springuser:123@(localhost:3306)/miniscopus?charset=utf8")
	if err != nil {
		panic(err)
	}

	papers := []DBPaper{}
	db.Select(&papers, "SELECT id, title, year FROM papers LIMIT 10")

	papers = fetchAuthors(db, papers)

	fmt.Println(papers)
}

func fetchAuthors(db *sqlx.DB, papers []DBPaper) []DBPaper {
	for j := 0; j < len(papers); j++ {
		ids := []AuthorsAndPapers{}

		db.Select(&ids, "SELECT * FROM authors_linked_to_papers WHERE paper_id=?", papers[j].ID.Int64)

		if len(ids) == 0 {
			continue
		}

		authors := []Author{}
		where := "id=" + strconv.FormatInt(ids[0].AuthorID, 10)

		for i := 1; i < len(ids); i++ {
			where = where + " OR " + "id=" + strconv.FormatInt(ids[i].AuthorID, 10)
		}

		db.Select(&authors, "SELECT first_name, last_name, middle FROM authors WHERE " + where)

		for _, author := range authors {
			papers[j].Authors = append(papers[j].Authors, author.FirstName.String+" "+author.Middle.String+" "+author.LastName.String)
		}
	}

	return papers
}
