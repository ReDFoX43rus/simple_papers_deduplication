package main

import (
	"github.com/jmoiron/sqlx"
)

func MergePapers(db *sqlx.DB, matched []MatchResult) {
	for _, papers := range matched {
		paper1 := papers.Paper1
		paper2 := papers.Paper2

		tx, err := db.Begin()
		if err != nil {
			continue
		}

		tx.Exec("UPDATE authors_linked_to_papers SET paper_id=? WHERE paper_id=?", paper2.ID.Value, paper1.ID.Value)
		// In case when paper1 is full duplicate (with foreign keys) of paper2
		tx.Exec("DELETE authors_linked_to_papers WHERE paper_id=?", paper1.ID)
		tx.Exec("DELETE FROM papers WHERE id=?", paper1.ID)

		tx.Commit()
	}
}