package registryimpl

// "service_name":["node1": ip+port, "node2": ip+port]

// 抽象 a service
type Service struct {
	Name  string  `json"name"`
	Nodes []*Node `json:"nodes"`
}

// single node 抽象
type Node struct {
	ID     int    `json:"id"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}
