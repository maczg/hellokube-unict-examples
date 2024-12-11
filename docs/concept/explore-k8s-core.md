# Exploring Kubernetes Architecture Concepts with Kind

After you have successfully installed **Kind** and created a local Kubernetes cluster, 
you can explore Kubernetes core architecture concepts by interacting with the cluster you’ve created. 

## Overview by Kubectl

### 1. Examine the Cluster Nodes

- **List the nodes in your cluster:**
  ```bash
  kubectl get nodes
    ```

Kind runs a Kubernetes cluster inside Docker containers.
Each node is actually a container acting as a node in the cluster.

- **Check node details:**
  ```bash
    kubectl describe node <node-name>
    ```
This will give you information about the node’s roles, conditions, and allocated resources.

### 2. Control Plane and Worker Nodes

The control plane **(API server, Scheduler, Controller Manager, and etcd)** runs within the "control-plane" node container.
The Kubernetes **API server** is the central management entity that receives all REST requests for the Kubernetes cluster.
It validates and configures data for the API objects that include pods, services, replication controllers, and others.
The API server services REST operations and provides the frontend to the cluster’s shared state through which all other components interact.

Worker nodes run the **kubelet** and **kube-proxy** services. 
Kubelet is responsible for managing the pods and containers on the node. 
It communicates with the **API server** to receive instructions and report the status of the node. 
In particular, it watches for **PodSpec** changes via the **Kube API Server**.

Kube-proxy is a network proxy that runs on each node in your cluster. It maintains network (routing) rules on nodes.
These network rules allow network communication to your Pods from network sessions inside or outside of your cluster.

### 3. Kubernetes Objects

Kubernetes objects are persistent entities in the Kubernetes system.
Kubernetes organizes (most of) resources (like Pods, Services, Deployments, etc.) into named collections called **namespaces**.
A namespace is a way to divide cluster resources between multiple users or to manage resources for different purposes (like production, development, or testing).

- **List all namespaces:**
  ```bash
  kubectl get namespaces
  ```
  
- **List all objects in a namespace:**
  ```bash
    kubectl get all -n <namespace>
    ```

- **List pods, deployments, and services:**
  ```bash
    kubectl get pods --all-namespaces
    kubectl get deployments --all-namespaces
    kubectl get svc --all-namespaces
    ```

### 4. Create and Explore Applications

- **Create a simple Nginx deployment:**
  ```bash
  kubectl create deployment nginx --image=nginx
  ```

## What's happening on the cluster nodes?

Explore the cluster nodes and the Kubernetes objects running on them.

SSH-ing into a node via docker:

```bash
docker exec -it <node-name> bash
```

Look at processes running on the node:

```bash
ps -aux
```


