## Voting app

Voting app is a simple web application that allows users to vote between two options. It is written in Python and uses the Flask web framework. Worker is a background process that fetches votes from Redis and stores them in Postgres.
It is based on the [Docker sample voting app](https://github.com/dockersamples/example-voting-app). 

[src](src) contains the source code for the application.

[manifests](manifests) contains the Kubernetes manifests for the application.

[docs](../../docs) contains general kubernetes concepts