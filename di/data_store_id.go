package main

import "fmt"

// Datastoreはデータを取得するためのインタフェイス
type Datastore interface {
	GetData(id string) (string, error)
}

// MySQLの実装
type MySQLDatastore struct{}

func (m MySQLDatastore) GetData(id string) (string, error) {
	// MySQLからデータを取得するロジック
	return "data from MySQL: " + id, nil
}

// ユーザーサービス
type UserService struct {
	Datastore Datastore // 依存関係
}

// ユーザーサービスのコンストラクタ
func NewUserService(ds Datastore) *UserService {
	return &UserService{Datastore: ds}
}

func (u *UserService) GetUserDetails(id string) (string, error) {
	// データストアからユーザー詳細を取得
	return u.Datastore.GetData(id)
}

// 注文サービス
type OrderService struct {
	Datastore Datastore // 依存関係
}

// 注文サービスのコンストラクタ
func NewOrderService(ds Datastore) *OrderService {
	return &OrderService{Datastore: ds}
}

func (o *OrderService) GetOrderDetails(id string) (string, error) {
	// データストアから注文詳細を取得
	return o.Datastore.GetData(id)
}

func hoge() {
	datastore := MySQLDatastore{}              // データストアの具体的なインスタンス
	userService := NewUserService(datastore)   // 依存関係を注入
	orderService := NewOrderService(datastore) // 依存関係を注入

	userData, err := userService.GetUserDetails("123") // ユーザー詳細を取得
	if err != nil {
		panic(err)
	}
	fmt.Println(userData) // Output: data from MySQL: 123

	orderData, err := orderService.GetOrderDetails("456") // 注文詳細を取得
	if err != nil {
		panic(err)
	}
	fmt.Println(orderData) // Output: data from MySQL: 456
}
