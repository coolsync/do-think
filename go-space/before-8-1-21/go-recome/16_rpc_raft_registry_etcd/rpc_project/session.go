package rpcimpl

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// network connect case ...

type Session struct {
	conn net.Conn
}

// 构造method
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// write data to connect
func (s *Session) Write(data []byte) error {
	buf := make([]byte, 4+len(data)) // 拼接 buf

	// header record data length
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))

	// add data
	_ = copy(buf[4:], data)

	// wirte to conn
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	// fmt.Println(string(buf))
	return nil
}

// from connect read data
func (s *Session) Read() ([]byte, error) {
	// read header 中 data len
	header := make([]byte, 4)

	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		fmt.Println("io.ReadFull: ", err)
		return nil, err
	}

	date_len := binary.BigEndian.Uint32(header)

	// 按data len read data []byte
	data := make([]byte, date_len)

	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		fmt.Println("io.ReadFull data: ", err)
		return nil, err
	}

	return data, nil
}
