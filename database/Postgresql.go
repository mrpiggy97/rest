package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mrpiggy97/rest/models"
)

type PostgresqlRepository struct {
	db *sql.DB
}

func closeQuery(query *sql.Rows) {
	err := query.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (repo *PostgresqlRepository) InsertUser(cxt context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(
		cxt,
		"INSERT INTO users(id,email,password) VALUES($1,$2,$3);",
		user.Id,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresqlRepository) GetUserById(cxt context.Context, id string) (*models.User, error) {
	query, queryErr := repo.db.QueryContext(cxt, "SELECT id,email FROM users WHERE id=$1;", id)
	if queryErr != nil {
		return nil, queryErr
	}
	var user *models.User = new(models.User)
	for query.Next() {
		scanningErr := query.Scan(&user.Id, &user.Email)
		if scanningErr != nil {
			return nil, scanningErr
		}
	}

	defer closeQuery(query)
	return user, nil
}

func (repo *PostgresqlRepository) GetUserByEmail(cxt context.Context, email string) (*models.User, error) {
	query, queryErr := repo.db.QueryContext(cxt, "SELECT id, email, password FROM users WHERE email=$1;", email)
	if queryErr != nil {
		return nil, queryErr
	}
	var user *models.User = new(models.User)
	for query.Next() {
		scanningErr := query.Scan(&user.Id, &user.Email, &user.Password)
		if scanningErr != nil {
			return nil, scanningErr
		}
	}

	defer closeQuery(query)
	return user, nil
}

func (repo *PostgresqlRepository) InsertPost(cxt context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(
		cxt,
		"INSERT INTO POSTS(id, post_content, user_id)VALUES($1,$2,$3)",
		post.Id, post.PostContent, post.UserId,
	)
	return err
}

func (repo *PostgresqlRepository) GetPostById(cxt context.Context, id string) (*models.Post, error) {
	query, queryErr := repo.db.QueryContext(cxt, "SELECT * FROM posts WHERE id=$1;", id)
	if queryErr != nil {
		return nil, queryErr
	}
	var post *models.Post = new(models.Post)
	for query.Next() {
		scanningErr := query.Scan(&post.Id, &post.PostContent, &post.UserId, &post.UserId)
		if scanningErr != nil {
			return nil, scanningErr
		}
	}

	defer closeQuery(query)
	return post, nil
}

func (repo *PostgresqlRepository) UpdatePost(cxt context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(
		cxt,
		"UPDATE posts SET post_content=$1 WHERE id=$2 and user_id=$3;",
		post.PostContent,
		post.Id,
		post.UserId,
	)
	return err
}

func (repo *PostgresqlRepository) DeletePost(cxt context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(
		cxt,
		"DELETE FROM posts where id=$1 and user_id=$2;",
		post.Id, post.UserId,
	)
	return err
}

func (repo *PostgresqlRepository) Close() {
	var err error = repo.db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func NewPostgresqlRepository(url string) (*PostgresqlRepository, error) {
	//postgres url should be in the following format
	// postgres://username:password@host:port/dbname?sslmode=disable
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	var repository *PostgresqlRepository = &PostgresqlRepository{
		db: db,
	}
	return repository, nil
}
