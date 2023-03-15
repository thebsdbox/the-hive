# The-Hive 🐝

We love docker and Kubernetes so much that we put docker in docker and kubernetes in that... if you want to go deeper, then feel free!

### Requirements

* [Docker `20.10.0+`](https://docs.docker.com/install/)
* [Go](https://golang.org/dl/) (stable release)

### Development

```bash
# Clone this repo locally
git clone https://github.com/thebsdbox/the-hive
cd play-with-docker

# Verify the Docker daemon is running
docker run hello-world

# Load the IPVS kernel module. Because swarms are created in dind,
# the daemon won't load it automatically
sudo modprobe xt_ipvs

# Ensure the Docker daemon is running in swarm mode
docker swarm init

# Get the latest franela/dind image
docker pull franela/dind

# Optional (with go1.14): pre-fetch module requirements into vendor
# so that no network requests are required within the containers.
# The module cache is retained in the pwd and l2 containers so the
# download is a one-off if you omit this step.
go mod vendor

# Start PWD as a container
docker-compose up
```

Now navigate to [http://localhost](http://localhost) and click the green "Start" button
to create a new session, followed by "ADD NEW INSTANCE" to launch a new terminal instance.

Notes:

* There is a hard-coded limit of 5 Docker playgrounds per session. After 4 hours sessions are deleted.
* If you want to override the DIND version or image then set the environmental variable i.e.
  `DIND_IMAGE=franela/docker<version>-rc:dind`. Take into account that you can't use standard `dind` images, only [franela](https://hub.docker.com/r/franela/) ones work.

### Port forwarding

In order for port forwarding to work correctly in development you need to make `*.localhost` to resolve to `127.0.0.1`. That way when you try to access  `pwd10-0-0-1-8080.host1.localhost`, then you're forwarded correctly to your local PWD server.

You can achieve this by setting up a `dnsmasq` server (you can run it in a docker container also) and adding the following configuration:

```
address=/localhost/127.0.0.1
```

Don't forget to change your computer's default DNS to use the dnsmasq server to resolve.

## FAQ

### How can I connect to a published port from the outside world?


If you need to access your services from outside, use the following URL pattern `http://ip<hyphen-ip>-<session_jd>-<port>.direct.labs.play-with-docker.com` (i.e: http://ip2-135-3-b8ir6vbg5vr00095iil0-8080.direct.labs.play-with-docker.com).

### Why is PWD running in ports 80 and 443? Can I change that?

No, it needs to run on those ports for DNS resolve to work. Ideas or suggestions about how to improve this
are welcome


## Troubleshooting

Errors like **Failed to create inotify object: Too many open files**, will require increasing `user.max_inotify_instances` with the command(s) 
```
 sudo sysctl fs.inotify.max_user_instances=8192
 sudo sysctl fs.inotify.max_user_watches=524288
 sudo sysctl -p
```
