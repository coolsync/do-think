package conf

type LogTransConf struct {
	KafkaConf `ini:"kafka"`
	EsConf    `ini:"es"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EsConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
	Nums        int    `ini:"nums"`
}
