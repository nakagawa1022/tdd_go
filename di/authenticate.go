package main

import "fmt"

// 認証インターフェース(プロバイダーによらず認証するメソッドは同じ)
type AuthMethod interface {
	Authenticate(username, password string) (string, error)
}

// LDAP認証
type LDAPAuth struct{}

// LDAP認証の実装
func (l LDAPAuth) Authenticate(username, password string) (string, error) {
	// LDAP認証のロジック
	return username, nil
}

// DB認証
type DBAuth struct{}

// DB認証の実装
func (d DBAuth) Authenticate(username, password string) (string, error) {
	// DB認証のロジック
	return username, nil
}

// ユーザー情報インターフェース
type UserProfile interface {
	GetUserProfile(username string) (string, error)
}

type UserProfileData struct{}

func (u UserProfileData) GetUserProfile(username string) (string, error) {
	return username, nil
}

// ロガーインターフェース作成
type Logger interface {
	Log(message string)
}

// ロガーの実装
type StdoutLogger struct{}

func (s StdoutLogger) Log(message string) {
	fmt.Println(message)
}

type AuthenticationService struct {
	AuthMethod  AuthMethod
	UserProfile UserProfile
	Logger      Logger
}

func NewAuthenticationService(am AuthMethod, up UserProfile, l Logger) *AuthenticationService {
	return &AuthenticationService{
		AuthMethod:  am,
		UserProfile: up,
		Logger:      l,
	}
}

func (a *AuthenticationService) Authenticate(username, password string) (string, error) {
	user, err := a.AuthMethod.Authenticate(username, password)
	if err != nil {
		return "", err
	}

	userProfile, err := a.UserProfile.GetUserProfile(user)
	if err != nil {
		return "", err
	}

	a.Logger.Log("user profile: " + userProfile)

	return userProfile, nil
}

func main() {
	ldap := LDAPAuth{}
	db := DBAuth{}
	userProfile := UserProfileData{}
	logger := StdoutLogger{}

	ldapAuthService := NewAuthenticationService(ldap, userProfile, logger)
	dbAuthService := NewAuthenticationService(db, userProfile, logger)

	userProfile1, err := ldapAuthService.Authenticate("user1", "pass1")
	if err != nil {
		panic(err)
	}
	fmt.Println(userProfile1)

	userProfile2, err := dbAuthService.Authenticate("user2", "pass2")
	if err != nil {
		panic(err)
	}
	fmt.Println(userProfile2)
}
