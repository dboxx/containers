## Install chart
```bash
> helm install --namespace=kube-system quagga . --logtostderr
```

## Open Quagga console
```bash
> kubectl exec -it quagga-55f7v -c quagga -- vtysh
```
