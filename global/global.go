package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"github.com/shinkaym/go-ecommerce-backend-api/pkg/logger"
	"github.com/shinkaym/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
	Mdbc   *sql.DB
	KafkaProducer *kafka.Writer
)
