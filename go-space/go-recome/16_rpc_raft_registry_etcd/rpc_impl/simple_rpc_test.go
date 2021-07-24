package rpcimpl

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

// def User
type User struct {
	Name string
	Age  int
}

func queryUsr(uid int) (User, error) {
	user_map := make(map[int]User)
	user_map[0] = User{Name: "bob", Age: 30}
	user_map[1] = User{Name: "mark", Age: 32}
	user_map[3] = User{Name: "paul", Age: 20}

	_, ok := user_map[uid]
	if !ok {
		return User{}, fmt.Errorf("user %d info is not found ...", uid)
		// return User{}, errors.New("user info is not found ...")
	}
	return user_map[uid], nil
}

func TestRPC(t *testing.T) {
	gob.Register(User{})

	addr := "localhost:8082"

	// server
	srv := NewServer(addr)
	srv.RegisterName("queryUsr", queryUsr)

	go srv.Run()

	// client

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	cli := NewClient(conn)

	var query_user func(int) (User, error)
	cli.callRPC("queryUsr", &query_user)
	// reflect: wrong return count from function created by MakeFunc [recovered]

	user, err := query_user(4) // err:gob: type not registered for interface: errors.errorString
	if err != nil {
		// t.Fatal(err)
		// t.Errorf()
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}
