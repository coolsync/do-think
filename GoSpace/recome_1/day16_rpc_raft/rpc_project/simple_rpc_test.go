package rpcimpl

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age  int
}

// panic: reflect.MakeFunc: value of type rpcimpl.User is not assignable to type *rpcimpl.User [recovered]
func queryUser(uid int) (User, error) {

	user_map := make(map[int]User)

	user_map[0] = User{Name: "bob", Age: 30}
	user_map[1] = User{Name: "mark", Age: 32}
	user_map[2] = User{Name: "paul", Age: 20}

	if user, ok := user_map[uid]; ok {
		return user, nil
	}

	return User{}, fmt.Errorf("user %d is not exists ... ", uid)
}

func TestRpc(t *testing.T) {
	gob.Register(User{})
	addr := "localhost:8080"

	// start server
	ser := NewServer(addr)

	ser.Register("queryUser", queryUser)

	go ser.Run()

	// cli
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("net dial err:%v\n", err)
	}
	defer conn.Close()

	cli := NewClient(conn)
	// set func variable
	var queryuser func(int) (User, error)
	cli.callRPC("queryUser", &queryuser)

	// query user data
	user, err := queryuser(4)
	if err != nil {
		t.Fatalf("queryUser err:%v\n", err)
	}

	fmt.Println("user:", user)

}
