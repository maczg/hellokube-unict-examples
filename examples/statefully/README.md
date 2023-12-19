### kubernetes statefulset vs deployment example 
 
This example show the difference between statefulset and deployment in kubernetes.

Briefly:
- statelfulset is used for stateful applications like databases. It provides stable network identity and persistent storage.
  this means that when a statefulset scale out, a new persistent volume is created and attached to the new pod. Keep in mind that the content
  of the new volume is not the same as the existing one. Application should be in charge of sincronizing the data between the volumes.

- deployment is more oriented to stateless applications. Scaling out/down does not care about persistentVolume. It means that when a deployment
   scale out, a new pod is created and attached to the same persistentVolume. The content of the new pod is the same as the existing one. It can be daungerous share same persistentVolume between different pods without 
   a policy to manage the datas. 

### Makefile

Look at Makefile to see how to create a statefulset and a deployment. It provides also utility command for create/destroy/configure kind cluster locally. 

NB: On Windows/Mac you may need to change routine describe by Makefile