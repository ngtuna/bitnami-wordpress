# bitnami-wordpress

Tiny webapp helps creating the [bitnami wordpress](https://aws.amazon.com/marketplace/pp/B00NN8Y43U) EC2 instance on AWS.

## Running locally

- You need `go`
- If your working copy is not in your `GOPATH`, you need to set it accordingly.
- Install `revel`

  `$ go get github.com/revel/cmd/revel`

- Fetching source code

  `$ git clone https://github.com/ngtuna/bitnami-wordpress.git $GOPATH/src/github.com/ngtuna/bitnami-wordpress`

- Set environment variables for launching configuration

  ```
  $ cat environment
  export IMAGEID=
  export INSTANCETYPE=
  export AWSREGION=
  export KEYNAME=
  export SECURITYGROUP=
  export SUBNET=

  $ source environment
  ```

- Run it

  `revel run github.com/ngtuna/bitnami-wordpress`

## Running with Docker

- Set environment variables for launching configuration in docker-compose.yml
- Run it

  `$ docker-compose up`
