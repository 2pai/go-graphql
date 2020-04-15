package links

import (
	database "github.com/2pai/go-graphql/internal/pkg/db/mysql"
	"github.com/2pai/go-graphql/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {

	statement, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := statement.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}
func GetAll() []Link {
	statements, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer statements.Close()
	rows, err := statements.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		links    []Link
		username string
		Id       string
	)

	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &Id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       Id,
			Username: username,
		}

		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return links
}
