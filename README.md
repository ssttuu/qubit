# Qubit

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

