package prometheus

import "github.com/ClareChu/tiger/prometheus/client"

type IstioMetrics struct {
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	Step      string `json:"step"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
type RequestCountData struct {
	ConnectionSecurityPolicy     string `json:"connection_security_policy"`
	DestinationApp               string `json:"destination_app"`
	DestinationCanonicalService  string `json:"destination_canonical_service"`
	DestinationPrincipal         string `json:"destination_principal"`
	DestinationService           string `json:"destination_service"`
	DestinationServiceName       string `json:"destination_service_name"`
	DestinationServiceNamespace  string `json:"destination_service_namespace"`
	DestinationVersion           string `json:"destination_version"`
	DestinationWorkload          string `json:"destination_workload"`
	DestinationWorkloadNamespace string `json:"destination_workload_namespace"`
	Instance                     string `json:"instance"`
	Job                          string `json:"job"`
	Namespace                    string `json:"namespace"`
	PodName                      string `json:"pod_name"`
	Reporter                     string `json:"reporter"`
	RequestProtocol              string `json:"request_protocol"`
	ResponseCode                 string `json:"response_code"`
	ResponseFlags                string `json:"response_flags"`
	SourceApp                    string `json:"source_app"`
	SourceCanonicalRevision      string `json:"source_canonical_revision"`
	SourceCanonicalService       string `json:"source_canonical_service"`
	SourcePrincipal              string `json:"source_principal"`
	SourceVersion                string `json:"source_version"`
	SourceWorkload               string `json:"source_workload"`
	SourceWorkloadNamespace      string `json:"source_workload_namespace"`
}

//RequestCount (istio_requests_total) COUNTER对于由Istio代理处理的每个请求，此计数均递增。
func (i *IstioMetrics) RequestCount(rcd *RequestCountData) (resp *client.ProRequest, err error) {
	return
}

type RequestDurationData struct {
	RequestCountData
	DestinationCanonicalService string `json:"destination_canonical_revision"`
}

//RequestCount (istio_request_duration_milliseconds) 这是DISTRIBUTION衡量请求持续时间的时间。
func (i *IstioMetrics) RequestDuration() (resp *client.ProRequest, err error) {
	return
}

//RequestSize (istio_request_bytes) 这是DISTRIBUTIONHTTP请求主体大小的度量。
func (i *IstioMetrics) RequestSize() (resp *client.ProRequest, err error) {
	return
}


//ResponseSize (istio_response_bytes) 这是DISTRIBUTION衡量HTTP响应主体大小的一个。
func (i *IstioMetrics) ResponseSize() (resp *client.ProRequest, err error) {
	return
}

type TcpByteSentData struct {
}

/*
istio_tcp_sent_bytes_total{connection_security_policy="none",destination_app="unknown",destination_canonical_revision="latest",destination_canonical_service="ingress-nginx",destination_principal="unknown",destination_service="mgmtCluster",destination_service_name="mgmtCluster",destination_service_namespace="ingress-nginx",destination_version="unknown",destination_workload="ingress-nginx-controller",destination_workload_namespace="ingress-nginx",instance="172.40.0.22:15090",job="envoy-stats",namespace="ingress-nginx",pod_name="ingress-nginx-controller-57d6f7ffd8-lkc98",reporter="destination",request_protocol="tcp",response_flags="-",source_app="unknown",source_canonical_revision="latest",source_canonical_service="unknown",source_principal="unknown",source_version="unknown",source_workload="unknown",source_workload_namespace="unknown"}
*/

//TcpByteSent (istio_tcp_sent_bytes_total) 此参数COUNTER用于度量在TCP连接情况下响应期间发送的总字节数。
func (i *IstioMetrics) TcpByteSent() (resp *client.ProRequest, err error) {
	return
}

/*istio_tcp_sent_bytes_total{
connection_security_policy="none",
destination_app="unknown",
destination_canonical_revision="latest",
destination_canonical_service="ingress-nginx",
destination_principal="unknown",
destination_service="mgmtCluster",
destination_service_name="mgmtCluster",
destination_service_namespace="ingress-nginx",
destination_version="unknown",
destination_workload="ingress-nginx-controller",
destination_workload_namespace="ingress-nginx",
instance="172.40.0.22:15090",
job="envoy-stats",
namespace="ingress-nginx",
pod_name="ingress-nginx-controller-57d6f7ffd8-lkc98",
reporter="destination",
request_protocol="tcp",
response_flags="-",
source_app="unknown",
source_canonical_revision="latest",
source_canonical_service="unknown",
source_principal="unknown",
source_version="unknown",
source_workload="unknown",
source_workload_namespace="unknown"}	*/

type TcpByteReceivedData struct {
}

//Tcp Byte Received（istio_tcp_received_bytes_total）：COUNTER用于测量在TCP连接情况下请求期间接收的总字节数。
func (i *IstioMetrics) TcpByteReceived() (resp *client.ProRequest, err error) {
	return
}

//Tcp Connections Opened（istio_tcp_connections_opened_total）：COUNTER每个打开的连接的增量。
func (i *IstioMetrics) TcpConnectionsOpened() (resp *client.ProRequest, err error) {
	return
}

//Tcp连接已关闭（istio_tcp_connections_closed_total）：COUNTER对于每个关闭的连接，此值均递增。
func (i *IstioMetrics) TcpConnectionsClosed() (resp *client.ProRequest, err error) {
	return
}
