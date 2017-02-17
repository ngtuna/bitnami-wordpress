# bitnami-wordpress

Tiny webapp helps creating the [bitnami wordpress](https://aws.amazon.com/marketplace/pp/B00NN8Y43U) EC2 instance on AWS.

## Running locally

- You need `go`
- If your working copy is not in your `GOPATH`, you need to set it accordingly.
- Install `revel`
  - `go get github.com/revel/cmd/revel`
  - `git clone https://github.com/ngtuna/bitnami-wordpress.git $GOPATH/src/github.com/ngtuna/bitnami-wordpress`
  - `revel run github.com/ngtuna/bitnami-wordpress`
  
## Running with Docker

`$ docker run -p 9000:9000 tuna/bitnami-wordpress-aws`



