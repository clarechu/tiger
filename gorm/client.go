package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"istio.io/pkg/log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Server struct {
	db       *gorm.DB
	consumer *kafka.Consumer
}

func NewServer() *Server {
	return &Server{
		db:       NewDBClient(),
		consumer: NewMQClient(),
	}
}

func main() {
	server := NewServer()
	for {
		msg, err := server.consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			// 操作数据库
			err = server.CreateProduct(msg.String())


		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	server.consumer.Close()
}

func NewMQClient() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
	return c
}

func NewDBClient() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	/*
		// Read
		var product Product
		db.First(&product, 1)                 // 根据整形主键查找
		db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

		// Update - 将 product 的 price 更新为 200
		db.Model(&product).Update("Price", 200)
		// Update - 更新多个字段
		db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
		db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

		// Delete - 删除 product
		db.Delete(&product, 1)*/
	return db
}

func (s *Server) CreateProduct(msg string) error {
	log.Debugf("msg --> %s", msg)
	var product Product
	s.db.First(&product, 1)                 // 根据整形主键查找
	s.db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	return nil
}
