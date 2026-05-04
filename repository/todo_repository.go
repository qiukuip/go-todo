package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Todo struct {
	TodoId     int       `json:"todoId"`
	Content    string    `json:"content"`
	Category   string    `json:"category"`
	IsComplete string    `json:"isComplete"`
	Deadline   time.Time `json:"deadline"`
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
}

const (
	dbUser     = "root"
	dbPassword = "mysql.root"
	dbName     = "my_demo"
	dbAddr     = "minipc:3306"
)

const (
	insertTodoSQL              = "insert into todo (content, category, is_complete, deadline) values (?, ?, ?, ?)"
	updateTodoSQL              = "update todo set content = ?, category = ?, is_complete = ?, deadline = ? where todo_id = ?"
	deleteTodoSQL              = "delete from todo where todo_id = ?"
	selectTodosByCategorySQL   = "select * from todo where category = ? order by todo_id"
	selectTodosByContentSQL    = "select * from todo where content like concat('%', ?, '%') order by todo_id"
	selectTodosByIsCompleteSQL = "select * from todo where is_complete = ? order by todo_id"
)

var (
	db *sql.DB
)

func getConnection() {
	if db != nil {
		fmt.Println("数据库连接已存在，跳过创建连接")
		return
	}

	log.SetPrefix("repository#getConnection: ")

	cfg := mysql.NewConfig()
	cfg.User = dbUser
	cfg.Passwd = dbPassword
	cfg.DBName = dbName
	cfg.Addr = dbAddr
	cfg.Net = "tcp"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("连接数据库失败！")
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("数据库连接成功！")
}

func AddTodo(todo Todo) (int64, error) {
	getConnection()

	result, err := db.Exec(insertTodoSQL, todo.Content, todo.Category, todo.IsComplete, todo.Deadline)
	if err != nil {
		return -1, fmt.Errorf("addTodo: %+v", err)
	}

	todoId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("addTodo: %+v", err)
	}

	return todoId, nil
}

func DeleteTodo(todoId int64) (int64, error) {
	getConnection()

	result, err := db.Exec(deleteTodoSQL, todoId)
	if err != nil {
		return -1, fmt.Errorf("deleteTodo: %+v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("deleteTodo: %+v", err)
	}

	return rowsAffected, nil
}

func UpdateTodo(todo Todo) (int64, error) {
	getConnection()

	result, err := db.Exec(updateTodoSQL, todo.Content, todo.Category, todo.Deadline, todo.TodoId)
	if err != nil {
		return -1, fmt.Errorf("updateTodo: %+v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("updateTodo: %+v", err)
	}

	return rowsAffected, nil
}

func SelectTodosByCategory(category string) ([]Todo, error) {
	getConnection()

	rows, err := db.Query(selectTodosByCategorySQL, category)
	if err != nil {
		return nil, fmt.Errorf("selectTodo %q: %+v", category, err)
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.TodoId, &todo.Content, &todo.Category, &todo.IsComplete, &todo.Deadline, &todo.CreateAt, &todo.UpdateAt); err != nil {
			return nil, fmt.Errorf("selectTodo %q: %+v", category, err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("selectTodo %q: %+v", category, err)
	}

	return todos, nil
}

func SelectTodosByIsComplete(isComplete string) ([]Todo, error) {
	getConnection()

	rows, err := db.Query(selectTodosByIsCompleteSQL, isComplete)
	if err != nil {
		return nil, fmt.Errorf("selectTodo %q: %+v", isComplete, err)
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.TodoId, &todo.Content, &todo.Category, &todo.IsComplete, &todo.Deadline, &todo.CreateAt, &todo.UpdateAt); err != nil {
			return nil, fmt.Errorf("selectTodo %q: %+v", isComplete, err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("selectTodo %q: %+v", isComplete, err)
	}

	return todos, nil
}

func SelectTodosByContent(content string) ([]Todo, error) {
	getConnection()

	rows, err := db.Query(selectTodosByContentSQL, content)
	if err != nil {
		return nil, fmt.Errorf("selectTodo %q: %+v", content, err)
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.TodoId, &todo.Content, &todo.Category, &todo.IsComplete, &todo.Deadline, &todo.CreateAt, &todo.UpdateAt); err != nil {
			return nil, fmt.Errorf("selectTodo %q: %+v", content, err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("selectTodo %q: %+v", content, err)
	}

	return todos, nil
}
