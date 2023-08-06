package main

import (
	"reflect"
	"testing"
)

func TestAuthenticationService_Authenticate(t *testing.T) {
	type fields struct {
		AuthMethod  AuthMethod
		UserProfile UserProfile
		Logger      Logger
	}
	type args struct {
		username string
		password string
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
			a := &AuthenticationService{
				AuthMethod:  tt.fields.AuthMethod,
				UserProfile: tt.fields.UserProfile,
				Logger:      tt.fields.Logger,
			}
			got, err := a.Authenticate(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBAuth_Authenticate(t *testing.T) {
	type args struct {
		username string
		password string
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
			d := DBAuth{}
			got, err := d.Authenticate(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLDAPAuth_Authenticate(t *testing.T) {
	type args struct {
		username string
		password string
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
			l := LDAPAuth{}
			got, err := l.Authenticate(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthenticationService(t *testing.T) {
	type args struct {
		am AuthMethod
		up UserProfile
		l  Logger
	}
	tests := []struct {
		name string
		args args
		want *AuthenticationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthenticationService(tt.args.am, tt.args.up, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthenticationService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStdoutLogger_Log(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StdoutLogger{}
			s.Log(tt.args.message)
		})
	}
}

func TestUserProfileData_GetUserProfile(t *testing.T) {
	type args struct {
		username string
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
			u := UserProfileData{}
			got, err := u.GetUserProfile(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
