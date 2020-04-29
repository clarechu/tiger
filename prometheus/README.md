# istio-promethues


### istio 定义资源模版

#### 对于HTTP，HTTP / 2和GRPC通信，Istio生成以下指标：

1. 请求计数（istio_requests_total）：COUNTER对于由Istio代理处理的每个请求，此计数均递增。

2. 请求持续时间（istio_request_duration_milliseconds）：这是DISTRIBUTION衡量请求持续时间的时间。

3. 请求大小（istio_request_bytes）：这是DISTRIBUTIONHTTP请求主体大小的度量。

响应大小（istio_response_bytes）：这是DISTRIBUTION衡量HTTP响应主体大小的一个。

#### 对于TCP流量，Istio生成以下指标：

Tcp Byte Sent（istio_tcp_sent_bytes_total）：此参数COUNTER用于度量在TCP连接情况下响应期间发送的总字节数。

Tcp Byte Received（istio_tcp_received_bytes_total）：COUNTER用于测量在TCP连接情况下请求期间接收的总字节数。

Tcp Connections Opened（istio_tcp_connections_opened_total）：COUNTER每个打开的连接的增量。

Tcp连接已关闭（istio_tcp_connections_closed_total）：COUNTER对于每个关闭的连接，此值均递增。