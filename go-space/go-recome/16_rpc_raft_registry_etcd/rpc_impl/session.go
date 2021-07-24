package rpcimpl

import (
	"encoding/binary"
	"io"
	"net"
)

// netword session case ...
type Session struct {
	conn net.Conn
}

// Session out bound api
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// write data to connect
func (s *Session) Write(data []byte) error {
	buf := make([]byte, 4+len(data))

	// record data len to header
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))

	// copy data to header tail
	_ = copy(buf[4:], data) // copy return copy bytes count

	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	// fmt.Println(buf)
	return nil
}

// read data from connect
func (s *Session) Read() ([]byte, error) {
	// from conn read header data
	header := make([]byte, 4)
	_, err := io.ReadFull(s.conn, header) // read 4 bytes to header
	if err != nil {
		return nil, err
	}
	// get header data length
	dataLen := binary.BigEndian.Uint32(header)

	data := make([]byte, dataLen)

	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
