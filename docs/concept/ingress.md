## Ingress 

Ingress is a kubernetes resource that allows you to expose your services under a single IP address and route the traffic to the appropriate containers based on the hostname and path.

Ingress resources are handled by ingress controllers. There are many ingress controllers available. The most popular is the nginx ingress controller.

It is important to note that ingress controllers are not part of the kubernetes project but are rather a third party resource. This means that operators are responsible for installing, configuring and maintaining the ingress controller.

From [kubernetes ingress documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/), ingress may provide load balancing, ssl termination and name-based virtual hosting.
Ingress expose HTTP and HTTPS routes from outside the cluster to *services* within the cluster. 

It worth to mention that ingress exposes HTTP/HTTPs route only and it's not suitable fof TCP/UDP traffic. Exposing services that uses arbitrary ports or protocols is allowed by using [NodePort](https://kubernetes.io/docs/concepts/services-networking/service/#nodeport) or [LoadBalancer](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer) services.

### Kind prerequisites

Kind requires additional configuration on cluster config files using `extraPortMappings` to expose ports on the host machine. 
This is required to expose the ingress controller on port 80 and 443. Also, we need to specify a label for the node to be used by ingress controller. For more details
check [kind documentation](https://kind.sigs.k8s.io/docs/user/ingress/).

In order to allow ingresses to respond to requests properly, a match between the domain name and the IP address of the ingress controller must be made. This can be done by adding the following line to the `/etc/hosts` file:
    
```bash
echo "my-test-domain.local" >> /etc/hosts
```

`/etc/hosts` file is a local file that allows you to map hostnames to IP addresses without using external DNS. (testing purposes only)


### Install the nginx ingress controller

Since we are using kind, we can install the nginx ingress controller using the following command:

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
```

### Deploy ingress resource

After installing the ingress controller, we can deploy ingress resources. 

```bash
kubectl apply -f ./voting-app/ingress/ingress.yaml
```

### Test the ingress

Open the browser and navigate on `http://my-test-domain.local`. You should see the voting app.


