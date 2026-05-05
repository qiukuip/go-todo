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

type QueryCondition struct {
	Field string
	Value any
}

const (
	dbUser     = "root"
	dbPassword = "pass.mysql"
	dbName     = "my_demo"
	dbAddr     = "minipc:3306"
)

const (
	insertTodoSQL         = "insert into todo (content, category, is_complete, deadline) values (?, ?, ?, ?)"
	updateTodoSQL         = "update todo set content = ?, category = ?, is_complete = ?, deadline = ? where todo_id = ?"
	deleteTodoSQL         = "delete from todo where todo_id = ?"
	selectTodoSQL         = "select todo_id, content, category, is_complete, deadline, create_at, update_at from todo"
	selectByIdSQL         = " where todo_id = ?"
	selectTodoCategorySQL = "select distinct category from todo order by category"
)

var db *sql.DB

func getConnection() {
	if db != nil {
		fmt.Println("数据库连接已存在，跳过创建连接")
		return
	}

	cfg := mysql.NewConfig()
	cfg.User = dbUser
	cfg.Passwd = dbPassword
	cfg.DBName = dbName
	cfg.Addr = dbAddr
	cfg.Net = "tcp"
	cfg.ParseTime = true
	cfg.Loc = time.Now().Local().Location()

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

func scanTodo(rows *sql.Rows) (Todo, error) {
	var todo Todo
	err := rows.Scan(&todo.TodoId, &todo.Content, &todo.Category, &todo.IsComplete, &todo.Deadline, &todo.CreateAt, &todo.UpdateAt)
	return todo, err
}

func queryTodos(conditions []QueryCondition) ([]Todo, error) {
	getConnection()

	whereClause := " where"
	var args []any
	for _, cond := range conditions {
		if whereClause != " where" {
			whereClause += " and"
		}
		if cond.Field == "content" {
			whereClause += fmt.Sprintf(" %s like concat('%%', ?, '%%')", cond.Field)
		} else {
			whereClause += fmt.Sprintf(" %s = ?", cond.Field)
		}
		args = append(args, cond.Value)
	}

	sql := selectTodoSQL + whereClause + " order by todo_id"
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		todo, err := scanTodo(rows)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

func AddTodo(todo Todo) (int64, error) {
	getConnection()

	result, err := db.Exec(insertTodoSQL, todo.Content, todo.Category, todo.IsComplete, todo.Deadline)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func DeleteTodo(todoId int64) (int64, error) {
	getConnection()

	result, err := db.Exec(deleteTodoSQL, todoId)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func UpdateTodo(todo Todo) (int64, error) {
	getConnection()

	result, err := db.Exec(updateTodoSQL, todo.Content, todo.Category, todo.IsComplete, todo.Deadline, todo.TodoId)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func SelectTodoCategory() ([]string, error) {
	getConnection()

	var categories []string

	rows, err := db.Query(selectTodoCategorySQL)
	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		var categoryItem string
		if err := rows.Scan(&categoryItem); err != nil {
			return categories, err
		}
		categories = append(categories, categoryItem)
	}

	return categories, nil
}

func SelectTodoByTodoId(todoId int64) (Todo, error) {
	getConnection()

	var todo Todo
	row := db.QueryRow(selectTodoSQL+selectByIdSQL, todoId)
	err := row.Scan(&todo.TodoId, &todo.Content, &todo.Category, &todo.IsComplete, &todo.Deadline, &todo.CreateAt, &todo.UpdateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return todo, fmt.Errorf("todo %d not found", todoId)
		}
		return todo, err
	}

	return todo, nil
}

func SelectTodosByCategory(category string) ([]Todo, error) {
	return queryTodos([]QueryCondition{{Field: "category", Value: category}})
}

func SelectTodosByContent(content string) ([]Todo, error) {
	return queryTodos([]QueryCondition{{Field: "content", Value: content}})
}

func SelectTodosByIsComplete(isComplete string) ([]Todo, error) {
	return queryTodos([]QueryCondition{{Field: "is_complete", Value: isComplete}})
}
