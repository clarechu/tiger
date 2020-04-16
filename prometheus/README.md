# istio-promethues


istio 定义资源模版

```yaml
apiVersion: "config.istio.io/v1alpha2"
kind: instance
metadata:
  name: requestsize
  namespace: istio-system
spec:
  compiledTemplate: metric
  params:
    value: request.size | 0
    dimensions:
      source_version: source.labels["version"] | "unknown"
      destination_service: destination.service.host | "unknown"
      destination_version: destination.labels["version"] | "unknown"
      response_code: response.code | 200
    monitored_resource_type: '"UNSPECIFIED"'
```