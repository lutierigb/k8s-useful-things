## What it does

Tries to connect to BigTable admin and BigTable API endpoints, creates tables, read rows, etc. Can be used to test connectivity from GKE to BigTable. Uses gRPC for the connection.

## Build and push image to Artifact Registry
```
V=X; docker build . -t asia-southeast1-docker.pkg.dev/lgb123/yellow/bigtable-connection:${V}; docker push asia-southeast1-docker.pkg.dev/lgb123/yellow/bigtable-connection:${V}
```