package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"os"
	"xlsxbot/models"
	"xlsxbot/app"
	"errors"
	"strconv"
)

func Connect() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
}

func UserExistsByUserID(user_id int64) (bool, error) {
	var exists bool
	err := app.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", user_id).Scan(&exists) 
	return exists, err
}

func InsertUser(user_id int64) (int, error) {
	var id int 
	query := `INSERT INTO users (user_id, stage) VALUES ($1, $2) RETURNING id`
	err := app.DB.QueryRow(context.Background(), query, user_id, "main").Scan(&id)
	return id, err
}

func GetStageByUserID(chatID int64) (string, error) {
	query := "SELECT stage FROM users where user_id = $1"
	var stage string
	err := app.DB.QueryRow(context.Background(), query, chatID).Scan(&stage)
	return stage, err
}

func SetStageByUserID(chatID int64, value string) (bool, error) {
	exists, err := UserExistsByUserID(chatID)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}
	query := "UPDATE users SET stage = $1 WHERE user_id = $2 RETURNING stage"
	var stage string
	err = app.DB.QueryRow(context.Background(), query, value, chatID).Scan(&stage)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetClassesByUserID(chatID int64) ([]models.Class, error) {
	query := "SELECT * FROM classes WHERE user_id = $1 ORDER BY id"
	rows, err := app.DB.Query(context.Background(), query, chatID)
	if err != nil {
		return nil, err
	}
	classes := make([]models.Class, 0)

	for rows.Next() {
		var class models.Class
		err := rows.Scan(&class.ID, &class.Name, &class.Grade, &class.UserID)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)

	}
	return classes, nil
}

func AddClass(chatID int64, className string, grade int) (int, error) {
	query := "INSERT INTO classes (name, grade, user_id) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := app.DB.QueryRow(context.Background(), query, className, grade, chatID).Scan(&id)
	return id, err
}

func RemoveClass(chatID int64, classID int) (bool, error) {
	query := "DELETE FROM classes WHERE user_id = $1 AND id = $2"
	exec, err := app.DB.Exec(context.Background(), query, chatID, classID)
	if err != nil {
		return false, err
	}
	if exec.RowsAffected() != 0 {
		return true, nil
	}
	return false, nil
}

func GetTemplatesByUserID(chatID int64) ([]models.Template, error) {
	query := "SELECT id, name, class_id, user_id, header, criteria FROM templates WHERE user_id = $1 ORDER BY id"
	rows, err := app.DB.Query(context.Background(), query, chatID)
	if err != nil {
		return nil, err
	}
	templates := make([]models.Template, 0)

	for rows.Next() {
		var template models.Template
		err := rows.Scan(&template.ID, &template.Name, &template.ClassID, &template.UserID, &template.Header, &template.Criteria)
		if err != nil {
			return nil, err
		}
		templates = append(templates, template)

	}
	return templates, nil
}

func GetTemplateByUserID(chatID int64, id int) (models.Template, error) {
	query := "SELECT id, name, header, user_id, criteria, class_id FROM templates WHERE user_id = $1 AND id = $2"
	var template models.Template
	rows, err := app.DB.Query(context.Background(), query, chatID, id)
	if err != nil {
		return template, err
	}
	if rows.Next() {
		err := rows.Scan(&template.ID, &template.Name, &template.Header, &template.UserID, &template.Criteria, &template.ClassID)
		if err != nil {
			return template, err
		}
	} else {
		return template, errors.New("template with ID" + strconv.Itoa(id) + "doesn't exist")
	}
	return template, nil
}

func GetClassByUserID(chatID int64, id int) (models.Class, error) {
	query := "SELECT * FROM classes WHERE user_id = $1 AND id = $2"
	var class models.Class
	rows, err := app.DB.Query(context.Background(), query, chatID, id)
	if err != nil {
		return class, err
	}
	if rows.Next() {
		err := rows.Scan(&class.ID, &class.Name, &class.Grade, &class.UserID)
		if err != nil {
			return class, err
		}
	}
	return class, nil
}

func AddStudentToUser(student models.Student) (int, error) {
	query := "INSERT INTO students (name, class_id, user_id) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := app.DB.QueryRow(context.Background(), query, student.Name, student.ClassID, student.UserID).Scan(&id)
	return id, err
}

func GetStudentsByClassID(classID int, userID int64) ([]models.Student, error) {
	query := "SELECT id, name, points, template_id FROM students WHERE class_id = $1 AND user_id = $2 ORDER BY id"
	rows, err := app.DB.Query(context.Background(), query, classID, userID)
	if err != nil {
		return nil, err
	}
	var students []models.Student 
	for rows.Next() {
		var student models.Student 
		err := rows.Scan(&student.ID, &student.Name, &student.Points, &student.TemplateID)
		if err != nil {
			return nil, err
		}
		student.ClassID = classID
		student.UserID = userID
		students = append(students, student)
	}
	return students, nil
}

func RemoveStudentByClassID(chatID int64, classID int, studentID int) (bool, error) {
	query := "DELETE FROM students WHERE user_id = $1 AND class_id = $2 AND id = $3"
	exec, err := app.DB.Exec(context.Background(), query, chatID, classID, studentID)
	if err != nil {
		return false, err
	}
	if exec.RowsAffected() != 0 {
		return true, nil
	}
	return false, nil
}

func AddTemplateStageFirst(template models.Template) (int, error) {
	query := "INSERT INTO templates (name, user_id, header) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := app.DB.QueryRow(context.Background(), query, template.Name, template.UserID, template.Header).Scan(&id)
	return id, err
}


func AddTemplateStageSecond(template models.Template) (string, error) {
	query := "UPDATE templates SET criteria = $1 WHERE user_id = $2 AND id = $3 RETURNING name"
	var name string
	err := app.DB.QueryRow(context.Background(), query, template.Criteria, template.UserID, template.ID).Scan(&name)
	if err != nil {
		return name, err
	}
	return name, nil
}

func DeleteTemplate(templateID int, userID int64) (bool, error) {
	query := "DELETE FROM templates WHERE user_id = $1 AND id = $2"
	exec, err := app.DB.Exec(context.Background(), query, userID, templateID)
	if err != nil {
		return false, err
	}
	if exec.RowsAffected() != 0 {
		return true, nil
	}
	return false, nil
}

func InsertPointsToStudent(student models.Student) (string, error) {
	query := "UPDATE students SET points = $1 WHERE user_id = $2 AND class_id = $3 and id = $4 RETURNING name"
	var name string
	err := app.DB.QueryRow(context.Background(), query, student.Points, student.UserID, student.ClassID, student.ID).Scan(&name)
	return name, err
}