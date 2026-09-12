package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type aaa struct {
	Id      int
	Login   string
	Pasword string
}

func main() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("ошибка подключения:", err)
		return
	}
	defer db.Close()

	fmt.Println("1 добавить пользователя")

	var t int
	fmt.Scan(&t)

	if t == 1 {
		addUser(db)
	}
	if t == 2 {
		e, err := enter(db)
		if err != nil {
			fmt.Println(err)
			return
		}
		addNotes(db, e)
		t, err := getNotes(db, e)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, i := range t {
			fmt.Println(i.Message)
		}
	}
}

type Notes struct {
	Id      int
	Message string
	UserId  int
}

func getNotes(db *sql.DB, polz aaa) ([]Notes, error) {
	query := `SELECT * FROM notes WHERE user_id=?`

	rows, err := db.Query(query, polz.Id)
	if err != nil {
		fmt.Println("fff", err)
		return []Notes{}, err
	}
	var notes []Notes

	for rows.Next() {
		var note Notes

		err := rows.Scan(
			&note.Id,
			&note.Message,
			&note.UserId,
		)
		if err != nil {
			return []Notes{}, err
		}

		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return []Notes{}, err
	}
	return notes, nil
}

func addNotes(db *sql.DB, e aaa) {
	fmt.Println("введи техт")
	var message string
	fmt.Scan(&message)

	query := `INSERT INTO notes (message,user_id) VALUES (?,?)`

	_, err := db.Exec(query, message, e.Id)
	if err != nil {
		fmt.Println("ошибка при ыыы низнать", err)
		return
	}

}

func addUser(db *sql.DB) {
	fmt.Println("введи логин")
	var login string
	fmt.Scan(&login)

	fmt.Println("введи пароль")
	var pasword string
	fmt.Scan(&pasword)

	query := `INSERT INTO users (login, pasword) VALUES (?,?)`

	_, err := db.Exec(query, login, pasword)
	if err != nil {
		fmt.Println("ошибка при добавлении пользователя", err)
		return
	}

}

func enter(db *sql.DB) (aaa, error) {
	fmt.Println("логин для входа")
	var login string
	fmt.Scan(&login)

	fmt.Println("логин для входа")
	var pasword string
	fmt.Scan(&pasword)

	user := aaa{}

	q := `SELECT *FROM users WHERE login=? AND pasword=?`
	err := db.QueryRow(q, login, pasword).Scan(&user.Id, &user.Login, &user.Pasword)
	if err != nil {
		fmt.Println("ошибка", err)
		return aaa{}, err
	}
	return user, nil
}
