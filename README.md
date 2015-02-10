# GAE Test Grounds
Google App Engine Test

## Development Environment

### On OSX

**TODO** Remove this since it's only for running App Engine apps not through managed vms
```
$ brew install go-app-engine-64
$ goapp
```


* Install Docker
  * Only version 1.3 works at the moment with gcloud
    * TODO: Test 1.4 version of docker/boot2docker
* Install gcloud utility
  * gcloud components update preview app

## Getting Started

Set up the project and assign it your Google Project ID

```
$ go get github.com/Tonkpils/gae-testgrounds
$ cd $GOPATH/src/github.com/Tonkpils/gae-testgrounds
$ gcloud config set project <PROJECT_ID>
$ gcloud preview app setup-managed-vms
```

### Problems

When running `gcloud preview app setup-managed-vms` a issues with TLS came up.

To fix:

```
$ boot2docker ssh
$ docker pull google/docker-registry
```

Then in local machine

```
$ gcloud auth print-refresh-token
```

and insert into a `registry-params.env`

```
$ cat > registry-params.env
GCP_OAUTH2_REFRESH_TOKEN=<YOUR_REFRESH_TOKEN_HERE>
GCS_BUCKET=containers-prod
```

Then runthe docker-registry container

```
$ docker run -d --env-file=registry-params.env -p 5000:5000 google/docker-registry
```

Then run again

```
$ gcloud preview app setup-managed-vms
```


## Running App Locally

To run the application locally simply run from the application's directory:

```
$ gcloud preview app run .
```

This will build and launch the application at `http://localhost:8080/`

**NOTE** Websocket support locally does not currently work.

## Deploying App

Deployment can be done by running

```
$ gcloud preview app deploy .
```

This will deploy and start the application listening on port 8080

To access the application visit:

```
<project_id>.appspot.com
```

You should now see a welcome message from your browser if everything worked fine.
You should also be able to send echo messages through the websocket.

