#!/usr/bin/env bash

gcloud container clusters get-credentials dev-cluster --zone us-central1-a

kubectl create -f ${1}
