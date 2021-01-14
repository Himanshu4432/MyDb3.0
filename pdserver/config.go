package pdserver

import (
	"log"
	"os"

	"github.com/youzan/ZanRedisDB/cluster/pdnode_coord"
)

type ServerConfig struct {
	HTTPAddress        string `flag:"http-address"`
	MetricAddress      string `flag:"metric-address"`
	BroadcastAddr      string `flag:"broadcast-address"`
	BroadcastInterface string `flag:"broadcast-interface"`

	ReverseProxyPort string `flag:"reverse-proxy-port"`
	ProfilePort      string `flag:"profile-port"`

	ClusterID                  string   `flag:"cluster-id"`
	ClusterLeadershipAddresses string   `flag:"cluster-leadership-addresses" cfg:"cluster_leadership_addresses"`
	AutoBalanceAndMigrate      bool     `flag:"auto-balance-and-migrate"`
	BalanceInterval            []string `flag:"balance-interval"`

	LogLevel         int32  `flag:"log-level" cfg:"log_level"`
	LogDir           string `flag:"log-dir" cfg:"log_dir"`
	DataDir          string `flag:"data-dir" cfg:"data_dir"`
	LearnerRole      string `flag:"learner-role" cfg:"learner_role"`
	FilterNamespaces string `flag:"filter-namespaces" cfg:"filter_namespaces"`
	BalanceVer       string `flag:"balance-ver" cfg:"balance_ver"`
}

func NewServerConfig() *ServerConfig {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	return &ServerConfig{
		HTTPAddress:        "0.0.0.0:18001",
		BroadcastAddr:      hostname,
		BroadcastInterface: "eth0",
		ProfilePort:        "7667",

		ClusterLeadershipAddresses: "",
		ClusterID:                  "",
		BalanceVer:                 pdnode_coord.BalanceV2Str,

		LogLevel: 1,
		LogDir:   "",
	}
}
