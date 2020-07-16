# istio 集成 prometheus-operator


1. 安装istio

```bash
istioctl install -f operator.yml
```

2. 安装prometheus-operator

```bash
kubectl create ns prometheus-operator

### 创建secret

kubectl apply -f secret.yml 

kubectl label namespace prometheus-operator istio-injection=enabled

```


3. 创建service-monitor

```bash
kubectl create ns servicemonitor.yaml
```

4. 安装prometheus

```bash
kubectl apply -f prome.yml
```

5. 创建Alert Manager

```bash
kubectl create secret generic alertmanager-main --from-file=./alertmanager/alertmanager.yaml -n prometheus-operator

kubectl apply -f alertmanager/alertmanager-stateful.yaml

```