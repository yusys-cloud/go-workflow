// Author: yangzq80@gmail.com
// Date: 2022/8/6
//
package admin

type Response struct {
	Code string
	Data interface{}
}

func NewResponse(data interface{}) *Response {
	return &Response{"200", data}
}

type Flow struct {
	Id   string
	Name string
	Node *Node
}
type Node struct {
	Id    string
	Name  string
	Users []*User
	child *Node
}
type User struct {
	Id   string
	Name string
}
