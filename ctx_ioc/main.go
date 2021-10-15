package main

import (
	"context"
	"fmt"
)

// 定义三种存储方式
type redis struct{}

func (r *redis) Save() {
	fmt.Println("save in redis")
}

type mysql struct{}

func (m *mysql) Save() {
	fmt.Println("save in mysql")
}

type kafka struct{}

func (kafka) Save() {
	fmt.Println("save in kafka")
}

// 定义存储接口
type StorageDriver interface {
	Save()
}

// 使用存储接口
func save(ctx context.Context) {
	val := ctx.Value("StorageDriver")
	driver := val.(StorageDriver)
	driver.Save()
}

// 传入不同的存储方法
func main() {
	ctx := context.Background()

	ctxReids := context.WithValue(ctx, "StorageDriver", &redis{})
	save(ctxReids)

	ctxMysql := context.WithValue(ctx, "StorageDriver", &mysql{})
	save(ctxMysql)

	ctxKafka := context.WithValue(ctx, "StorageDriver", &kafka{})
	save(ctxKafka)
}
