# Kubernetes

## Environment

For simplicty, I have been using Rancher to create a sandbox Kubernetes cluster. It was fast to spin up and I've enjoyed working with it so far.

## nginx.yaml

In this toy example, I used a `kubectl` dry run to generate a minimal yaml configuration for a pod and played around with some of the tags.

## nginx-deploy.yaml

Scaling up a bit. Instead of just creating one pod, I'm making a whole deployment. Interesting to watch how the containers responded to an exit code of 1. They enter a CrashLoopBackoff and the changes are not applied to every container.

## Mealie

Checking out creating namespaces and using the [Mealie](https://hub.docker.com/r/hkotel/mealie) container.
