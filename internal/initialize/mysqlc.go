package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/shinkaym/go-ecommerce-backend-api/global"
	"github.com/shinkaym/go-ecommerce-backend-api/internal/model"
	"go.uber.org/zap"
	"gorm.io/gen"
)

func checkErrorPanicc(err error, errString string) {
	if err != nil {
		global.Logger.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlc() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL successfully")
	global.Mdbc = db

	// set Pool
	SetPool()
	// genTableDAO()
	// migrateTables()
}

func SetPoolc() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s\n", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

// mysql to model
func genTableDAOc() {
	// Initiate the tables
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	g.GenerateModel("go_crm_user")
	// g.GenerateAllTable()
	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})
	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})
	// Generate the code
	g.Execute()
}

func migrateTablesc() {
	err := global.Mdb.AutoMigrate(
		// &po.User{},
		// &po.Role{},
		&model.GoCrmUserV2{},
	)
	if err != nil {
		fmt.Println("Migration tables error: ", err)
	}
}
