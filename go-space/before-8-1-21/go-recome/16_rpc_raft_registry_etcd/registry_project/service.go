package registry

// 抽象出一些结构体
// ▪ Node：单个节点的结构体，包含 id ip port weight（权重）
// ▪ Service：里面有服务名，还有节点列表，一个服务多台服务器支撑

// Abstract a service
type Service struct{
	Name string `json:"name"`
	Nodes []*Node `json:"nodes"`
}

// Single Node Abstract 
type Node struct {
	// ID int `json:"id"`
	IP string `json:"ip"`
	Port int `json:"port"`
	// Weight int `json:"weight"`
}
