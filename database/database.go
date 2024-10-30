package database

import (
	"database/sql"
	"errors"

	"github.com/assaidy/url-shortener/models"
	_ "github.com/mattn/go-sqlite3"
)

type DBService struct {
	db *sql.DB
}

var (
	dbName   = "database/urls.db"
	instance *DBService
)

func NewDBService() *DBService {
	if instance != nil {
		return instance
	}
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	instance = &DBService{db: db}
	return instance
}

func (dbs *DBService) InsertURL(inout *models.URL) error {
	query := `
    insert into urls (original_url, short_code, created_at, updated_at)
    values(?, ?, ?, ?);
    `
	res, err := dbs.db.Exec(
		query,
		inout.OriginalURL,
		inout.ShortCode,
		inout.CreatedAt,
		inout.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	inout.Id = int(id)

	return nil
}

func (dbs *DBService) GetURL(sc string) (*models.URL, error) {
	query := `
    SELECT
        id,
        original_url,
        created_at,
        updated_at
    FROM urls
    WHERE short_code = ?;
    `
	url := models.URL{ShortCode: sc}
	if err := dbs.db.QueryRow(query, sc).Scan(
		&url.Id,
		&url.OriginalURL,
		&url.CreatedAt,
		&url.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if err := dbs.incrementAccessCount(sc); err != nil {
		return nil, err
	}
	return &url, nil
}

func (dbs *DBService) GetURLWithAccessCount(sc string) (*models.URL, error) {
	query := `
    SELECT
        id,
        original_url,
        created_at,
        updated_at,
        access_count
    FROM urls
    WHERE short_code = ?;
    `
	url := models.URL{ShortCode: sc}
	if err := dbs.db.QueryRow(query, sc).Scan(
		&url.Id,
		&url.OriginalURL,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.AccessCount,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &url, nil
}

func (dbs *DBService) incrementAccessCount(sc string) error {
	query := `
    UPDATE urls
    SET
        access_count = access_count + 1
    WHERE short_code = ?;
    `
	if _, err := dbs.db.Exec(query, sc); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) UpdateURL(url *models.URL) error {
	query := `
    UPDATE urls
    SET
        original_url = ?,
        updated_at = ?
    WHERE short_code = ?;
    `
	if _, err := dbs.db.Exec(
		query,
		url.OriginalURL,
		url.UpdatedAt,
		url.ShortCode,
	); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) DeleteURL(sc string) error {
	query := `DELETE FROM urls WHERE short_code = ?;`
	if _, err := dbs.db.Exec(query, sc); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) CheckIfURLExists(sc string) (bool, error) {
	query := `SELECT 1 FROM urls WHERE short_code = ? LIMIT 1;`
	if err := dbs.db.QueryRow(query, sc).Scan(new(int)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
