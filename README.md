## Hello Kubernetes 👋

This repository contains examples and exercises for the Distributed Computing course at University of Catania.

[docs](docs) contains theory notes and tips regarding Kubernetes, Docker and related technologies. 
The aim is to extend [docs](docs) overtime with more content and examples and use it as a reference for the course.

⚠️ **This repository is a WIP**. Any contribution is welcome.

It covers the following topics:

- [x] Setting up a Kubernetes development local cluster
- [x] Deploying various applications to the cluster
- [x] Exposing the applications to the outside world
- [x] Scaling the applications
- [x] Deep dive into Kubernetes objects and basic debugging practices
- [x] Monitoring tools


### Contents

- [kind](kind) contains relevant files to set up a local kubernetes cluster using kind.

- [examples](examples) contains examples and PoC to demonstrate containers and kubernetes features.
  - [voting-app](examples/voting-app) contains source and deployment file for two-option voting application.

- [docs](docs) contains theory notes and tips regarding Kubernetes, Docker and related technologies.
