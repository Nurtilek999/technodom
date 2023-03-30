package pgrepository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"technodom/internal/app/config"
	"technodom/internal/entity"
	"technodom/internal/repository"
	"technodom/internal/util/logger"
)

type PgRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func New(config *config.TomlConfig, logger logger.Logger) repository.Repository {
	db, err := sql.Open("postgres", config.Technodom.DatabaseUrl)
	if err != nil {
		log.Panicln("error on connection: ", err)
	}
	fmt.Println(config.Technodom.DatabaseUrl)

	err = db.Ping()
	if err != nil {
		println(err.Error())
		log.Panicln("error on ping: ", err)

	}

	return &PgRepository{db: db, logger: logger}
}

func (repo *PgRepository) Close() {
	_ = repo.db.Close()
}

func (repo *PgRepository) GetList(limit int, offset int) (*sql.Rows, error) {
	rows, err := repo.db.Query(`select * from urls limit $1 offset $2`, &limit, &offset)
	if err != nil {
		errMsg := fmt.Errorf("error in select query %w", err).Error()
		repo.logger.Error(errMsg, map[string]interface{}{})
		return nil, err
	}
	return rows, nil
}

func (repo *PgRepository) GetByID(id string) *sql.Row {
	row := repo.db.QueryRow(`select * from urls where active_link = $1 or history_link = $1`, &id)
	return row
}

func (repo *PgRepository) Post(url entity.Url) error {
	_, err := repo.db.Exec(`insert into urls(active_link, history_link) values ($1, $2)`, &url.ActiveLink, &url.HistoryLink)
	if err != nil {
		errMsg := fmt.Errorf("error in insert query %w", err).Error()
		repo.logger.Error(errMsg, map[string]interface{}{})
		return err
	}
	return nil
}

func (repo *PgRepository) Update(id, activeLink string) error {
	_, err := repo.db.Exec(`update urls set active_link = $1, history_link = $2 where active_link = $2`, &activeLink, &id)
	if err != nil {
		errMsg := fmt.Errorf("error in update query %w", err).Error()
		repo.logger.Error(errMsg, map[string]interface{}{})
		return err
	}
	return nil
}

func (repo *PgRepository) Delete(id string) error {
	_, err := repo.db.Exec(`delete from urls where active_link = $1 or history_link = $1`, id)
	if err != nil {
		errMsg := fmt.Errorf("error in delete query %w", err).Error()
		repo.logger.Error(errMsg, map[string]interface{}{})
		return err
	}
	return nil
}
