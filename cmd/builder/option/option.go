// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package option

import "github.com/spf13/pflag"
import "github.com/Sirupsen/logrus"
import "fmt"

//Config config server
type Config struct {
	EtcdEndPoints        []string
	EtcdTimeout          int
	EtcdPrefix           string
	ClusterName          string
	MysqlConnectionInfo  string
	DBType               string
	PrometheusMetricPath string
	EventLogServers      []string
	KubeConfig           string
	MaxTasks             int
	APIPort              int
	MQAPI                string
	DockerEndpoint       string
	HostIP               string
	CleanUp              bool
}

//Builder  builder server
type Builder struct {
	Config
	LogLevel string
	RunMode  string //default,sync
}

//NewBuilder new server
func NewBuilder() *Builder {
	return &Builder{}
}

//AddFlags config
func (a *Builder) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&a.LogLevel, "log-level", "info", "the entrance log level")
	fs.StringSliceVar(&a.EtcdEndPoints, "etcd-endpoints", []string{"http://127.0.0.1:2379"}, "etcd v3 cluster endpoints.")
	fs.IntVar(&a.EtcdTimeout, "etcd-timeout", 5, "etcd http timeout seconds")
	fs.StringVar(&a.EtcdPrefix, "etcd-prefix", "/store", "the etcd data save key prefix ")
	fs.StringVar(&a.PrometheusMetricPath, "metric", "/metrics", "prometheus metrics path")
	fs.StringVar(&a.DBType, "db-type", "mysql", "db type mysql or etcd")
	fs.StringVar(&a.MysqlConnectionInfo, "mysql", "root:admin@tcp(127.0.0.1:3306)/region", "mysql db connection info")
	fs.StringSliceVar(&a.EventLogServers, "event-servers", []string{"127.0.0.1:6367"}, "event log server address. simple lb")
	fs.StringVar(&a.KubeConfig, "kube-config", "/etc/goodrain/kubernetes/admin.kubeconfig", "kubernetes api server config file")
	fs.IntVar(&a.MaxTasks, "max-tasks", 50, "the max tasks for per node")
	fs.IntVar(&a.APIPort, "api-port", 3228, "the port for api server")
	fs.StringVar(&a.MQAPI, "mq-api", "127.0.0.1:6300", "acp_mq api")
	fs.StringVar(&a.RunMode, "run", "sync", "sync data when worker start")
	fs.StringVar(&a.DockerEndpoint, "dockerd", "127.0.0.1:2376", "dockerd endpoint")
	fs.StringVar(&a.HostIP, "hostIP", "", "Current node Intranet IP")
	fs.BoolVar(&a.CleanUp, "clean-up", false, "Turn on build version cleanup")
}

//SetLog 设置log
func (a *Builder) SetLog() {
	level, err := logrus.ParseLevel(a.LogLevel)
	if err != nil {
		fmt.Println("set log level error." + err.Error())
		return
	}
	logrus.SetLevel(level)
}

//CheckEnv 检测环境变量
func (a *Builder) CheckEnv() error {

	return nil
}
