package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
	// TaillogConf `ini:"taillog"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	// Topic   string `ini:"topic"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"collect_log_key"`
	Timeout int    `ini:"timeout"`
}

type TaillogConf struct {
	FileName string `ini:"filename"`
}
