package users

import "errors"

type User struct {
	Uid         string
	PhoneNumber string
}

var Store = make(map[string]User)

func AddUser(user User) (User, error) {
	// 假设这里有一些逻辑来将用户添加到数据库中
	// 并返回新添加的用户对象
	Store[user.Uid] = user
	return user, nil
}

func GetUser(uid string) (u User, err error) {
	u, ok := Store[uid]
	if !ok {
		err = errors.New("404")
	}
	return
}
