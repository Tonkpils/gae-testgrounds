# GAE Test Grounds
Google App Engine Test

## Development Environment

### On OSX

* Install boot2docker
  * Need docker TLS turned off...
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

**ANOTHER**

```
$ boot2docker ssh
$ sudo vi /var/lib/boot2docker/profile
```

and insert the following line

```
export DOCKER_TLS=no
```

Then on local machine

```
export DOCKER_HOST=tcp://192.168.59.103:2375
unset DOCKER_CERT_PATH
unset DOCKER_TLS_VERIFY
```

Seems like gcloud TLS support is a bit messy...

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

