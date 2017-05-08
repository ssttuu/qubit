# Qubit
[![CircleCI](https://circleci.com/gh/stupschwartz/qubit/tree/master.svg?style=svg&circle-token=91ae7b7dd5787a1c7c4250d32b91da07a4a471b0)](https://circleci.com/gh/stupschwartz/qubit/tree/master)

### Setup

1.  Install jq
2.  Install jinja2-cli
3.  Install gclou
4.  Install kubectl

#### Service Account Secret setup
If this hasn't already been done, you'll need to create the service account kubernetes secret.

Download the service key and create the secret.

```bash
kubectl create secret generic google-service-account --from-file=./credentials/service-account-key-file.json
```


# Compute



## Deployment

