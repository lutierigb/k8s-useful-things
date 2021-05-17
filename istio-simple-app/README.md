**This example creates the following services:**

 - Service A (which is just an nginx proxy configured with proxy_pass to service B)
 - Service B V1 - just an ngnix service a yellow page 
 - Service B V1 - just an ngnix service a grey page 

**The following Istio Resources are created: **
 - Gateway -  that attaches to the default istio ingress, opens port 80 and answer for requests to myweb.example.com
 - VirtualService - routes all the requests to myweb.example.com to k8s svc Service-A 

**Connecting to the service:**

```
export ING_IP=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
curl -v http://myweb.example.com --resolve myweb.example.com:80:${ING_IP}
```
