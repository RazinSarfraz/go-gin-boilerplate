package postres

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
	"user-backend/conf"
	c "user-backend/context"
	"user-backend/models"
	"user-backend/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var gdb *gorm.DB
var storeOnce sync.Once
var store repository.Store

type Store struct {
	db *gorm.DB
}

// SharedStore return global or single instance of firebase connection (bounded in sync once)
func SharedStore() repository.Store {
	storeOnce.Do(func() {
		err := initDb()
		if err != nil {
			panic(err)
		}
		store = NewStore(gdb)
	})
	return store
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	cfg := conf.GetConfig()
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable client_encoding=UTF8",
		cfg.PostDataSource.Addr, cfg.PostDataSource.Port, cfg.PostDataSource.User, cfg.PostDataSource.Password, cfg.PostDataSource.Database)
	var err error
	gormOpt := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}
	gdb, err = gorm.Open(postgres.Open(url), gormOpt)
	for err != nil {
		fmt.Println("Retring connecting to database")
		time.Sleep(5 * time.Second)
		gdb, err = gorm.Open(postgres.Open(url))
		continue
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)

	if cfg.PostDataSource.EnableAutoMigrate {
		var tables = []interface{}{
			&models.User{},
		}
		for _, table := range tables {
			if err := gdb.AutoMigrate(table); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Store) BeginTx(ctx c.Context) (c.Context, error) {
	opts := &sql.TxOptions{}
	tx := s.db.Begin(opts)
	if tx.Error != nil {
		return c.Context{}, tx.Error
	}
	newCtx := ctx.WithTx(tx)
	return newCtx, nil
}

func (s *Store) Rollback(ctx c.Context) error {
	return ctx.Tx().Rollback().Error
}

func (s *Store) CommitTx(ctx c.Context) error {
	return ctx.Tx().Commit().Error
}
