package main

import (
	"reflect"
	"testing"
)

func TestMySQLDatastore_GetData(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MySQLDatastore{}
			got, err := m.GetData(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrderService(t *testing.T) {
	type args struct {
		ds Datastore
	}
	tests := []struct {
		name string
		args args
		want *OrderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderService(tt.args.ds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserService(t *testing.T) {
	type args struct {
		ds Datastore
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.ds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_GetOrderDetails(t *testing.T) {
	type fields struct {
		Datastore Datastore
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderService{
				Datastore: tt.fields.Datastore,
			}
			got, err := o.GetOrderDetails(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetOrderDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserDetails(t *testing.T) {
	type fields struct {
		Datastore Datastore
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				Datastore: tt.fields.Datastore,
			}
			got, err := u.GetUserDetails(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}
