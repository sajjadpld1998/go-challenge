package repositories

import (
	"fmt"
	"math"
	"skeleton/config"
	"skeleton/error_handler"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ReadConnection  *gorm.DB
	WriteConnection *gorm.DB
)

type Pagination struct {
	// count of elements in per page
	Limit int `json:"limit,omitempty;query:limit" example:"20"`
	// current page number
	Page int `json:"page,omitempty;query:page" example:"3"`
	// sort order by
	Sort string `json:"sort,omitempty;query:sort" example:"created_at desc"`
	// count of total elements exists
	TotalRows int64 `json:"total_rows" example:"1000"`
	// count of total pages
	TotalPages int `json:"total_pages" example:"12"`
	// elements
	Data interface{} `json:"data"`
}

const (
	DefaultPaginationPage  = 1
	DefaultPaginationLimit = 10
	MaximumPaginationLimit = 30
)

func InitDBConnection() {
	readConnectionInit()
	writeConnectionInit()
}

func connectionToPostgres(dbHost, dbUser, dbPassword, dbName, sslmode string, dbPort int) *gorm.DB {
	PgDataLoad := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, sslmode)

	db, err := gorm.Open(postgres.Open(PgDataLoad), &gorm.Config{})
	if err != nil {
		error_handler.ThrowServerError(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		error_handler.ThrowServerError(err)
	}
	sqlDB.SetMaxIdleConns(5)  //100//2
	sqlDB.SetMaxOpenConns(10) //300//5
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	return db
}

func readConnectionInit() {
	dbHost := config.GetConfig().Database.LandrockerRead.Host
	dbPort := config.GetConfig().Database.LandrockerRead.Port
	dbName := config.GetConfig().Database.LandrockerRead.Database
	dbUser := config.GetConfig().Database.LandrockerRead.User
	dbPassword := config.GetConfig().Database.LandrockerRead.Password
	sslmode := config.GetConfig().Database.LandrockerRead.Ssl

	db := connectionToPostgres(dbHost, dbUser, dbPassword, dbName, sslmode, dbPort)
	ReadConnection = db
}

func writeConnectionInit() {
	dbHost := config.GetConfig().Database.LandrockerWrite.Host
	dbPort := config.GetConfig().Database.LandrockerWrite.Port
	dbName := config.GetConfig().Database.LandrockerWrite.Database
	dbUser := config.GetConfig().Database.LandrockerWrite.User
	dbPassword := config.GetConfig().Database.LandrockerWrite.Password
	sslmode := config.GetConfig().Database.LandrockerWrite.Ssl

	db := connectionToPostgres(dbHost, dbUser, dbPassword, dbName, sslmode, dbPort)
	WriteConnection = db
}

func (p *Pagination) getOffset() int {
	return (p.getPage() - 1) * p.getLimit()
}

func (p *Pagination) getLimit() int {
	if p.Limit == 0 {
		p.Limit = DefaultPaginationLimit
	}
	return p.Limit
}

func (p *Pagination) getPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) getSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

func (pagination *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	totalPages := int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.getOffset()).Limit(pagination.getLimit()).Order(pagination.getSort())
	}
}

func CheckExistsCustom(table, column, value string) bool {
	var exists int64
	err := ReadConnection.Table(table).Where(column+"::text = ? AND deleted_at IS NULL", value).Count(&exists).Error

	if err != nil {
		error_handler.ThrowServerError(err)
	}

	return exists > 0
}
