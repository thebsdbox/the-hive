# play-with-docker

Play With Docker gives you the experience of having a free Alpine Linux Virtual Machine in the cloud
where you can build and run Docker containers and even create clusters with Docker features like Swarm Mode.

Under the hood DIND or Docker-in-Docker is used to give the effect of multiple VMs/PCs.

A live version is available at: http://play-with-docker.com/

## Requirements

Docker 1.13+ is required. 

The docker daemon needs to run in swarm mode because PWD uses overlay attachable networks. For that
just run  `docker swarm init` in the destination daemon.

It's also necessary to manually load the IPVS kernel module because as swarms are created in `dind`, 
the daemon won't load it automatically. Run the following command for that purpose: `sudo modprobe xt_ipvs`


## Development

Start the Docker daemon on your machine and run `docker pull franela/dind`. 

1) Install go 1.7.1 with `brew` on Mac or through a package manager.

2) `go get -v -d -t ./...`

3) Start PWD as a container with docker-compose up.

5) Point to http://localhost and click "New Instance"

Notes:

* There is a hard-coded limit to 5 Docker playgrounds per session. After 4 hours sessions are deleted.
* If you want to override the DIND version or image then set the environmental variable i.e.
  `DIND_IMAGE=franela/docker<version>-rc:dind`. Take into account that you can't use standard `dind` images, only [franela](https://hub.docker.com/r/franela/) ones work.


## FAQ

### How can I connect to a published port from the outside world?

~~We're planning to setup a reverse proxy that handles redirection automatically, in the meantime you can use [ngrok](https://ngrok.com) within PWD running `docker run --name supergrok -d jpetazzo/supergrok` then `docker logs --follow supergrok` , it will give you a ngrok URL, now you can go to that URL and add the IP+port that you want to connect to… e.g. if your PWD instance is 10.0.42.3, you can go to http://xxxxxx.ngrok.io/10.0.42.3:8000 (where the xxxxxx is given to you in the supergrok logs).~~

If you need to access your services from outside, use the following URL pattern `http://pwd<underscore_ip>-<port>.<host#>.labs.play-with-docker.com` (i.e: http://pwd10_2_135_3-80.host3.labs.play-with-docker.com/).

### Why is PWD running in ports 80 and 443?, Can I change that?.

No, it needs to run on those ports for DNS resolve to work. Ideas or suggestions about how to improve this
are welcome
