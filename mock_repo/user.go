//go:generate mockgen -source=user.go -destination=./mock.go -package=mock_repo

package mock_repo

type User interface {
	Update(user *User) error
}
