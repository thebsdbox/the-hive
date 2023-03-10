ARG VERSION=docker:dind
FROM ${VERSION}

RUN apk add --no-cache py-pip python3-dev libffi-dev openssl-dev git tmux apache2-utils vim build-base gettext-dev curl bash-completion bash util-linux jq openssh openssl tree \
    && ln -s /usr/local/bin/docker /usr/bin/docker

ENV GOPATH /root/go
ENV PATH $PATH:$GOPATH


ENV DOCKER_TLS_CERTDIR=""
ENV DOCKER_CLI_EXPERIMENTAL=enabled
ENV COMPOSE_VERSION=2.6.1


# Add bash completion and set bash as default shell
RUN mkdir -p /usr/lib/docker/cli-plugins \
    && curl -LsS https://github.com/docker/compose/releases/download/v$COMPOSE_VERSION/docker-compose-linux-x86_64 -o /usr/lib/docker/cli-plugins/docker-compose \
    && chmod +x /usr/lib/docker/cli-plugins/docker-compose

# Add bash completion and set bash as default shell
RUN mkdir /etc/bash_completion.d \
    && curl -sS https://raw.githubusercontent.com/docker/cli/master/contrib/completion/bash/docker -o /etc/bash_completion.d/docker \
    && sed -i "s/ash/bash/" /etc/passwd
 
# Replace modprobe with a no-op to get rid of spurious warnings
# (note: we can't just symlink to /bin/true because it might be busybox)
RUN rm /sbin/modprobe && echo '#!/bin/true' >/sbin/modprobe && chmod +x /sbin/modprobe

# Install a nice vimrc file and prompt (by soulshake)
COPY ["docker-prompt", "sudo", "/usr/local/bin/"]
COPY [".vimrc", ".profile", ".inputrc", ".gitconfig", "./root/"]
COPY ["motd", "/etc/motd"]
COPY ["daemon.json", "/etc/docker/"]


# Move to our home
WORKDIR /root

# Setup certs and ssh keys
RUN mkdir -p /var/run/pwd/certs && mkdir -p /var/run/pwd/uploads \
    && ssh-keygen -N "" -t rsa -f  /etc/ssh/ssh_host_rsa_key >/dev/null \
    && mkdir ~/.ssh && ssh-keygen -N "" -t rsa -f ~/.ssh/id_rsa \
    && cat ~/.ssh/id_rsa.pub > ~/.ssh/authorized_keys

# Install kind
RUN curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.17.0/kind-linux-amd64 \
    && chmod +x ./kind \
    && sudo mv ./kind /usr/local/bin/kind

# Install kubectl
RUN curl -Lo ./kubectl https://dl.k8s.io/release/v1.25.3/bin/linux/amd64/kubectl \
    && chmod +x ./kubectl \
    && sudo mv ./kubectl /usr/local/bin/kubectl

# Install Cilium CLI
RUN curl -L --fail --remote-name-all https://github.com/cilium/cilium-cli/releases/download/v0.13.0/cilium-linux-amd64.tar.gz \
    && tar xzv -f cilium-linux-amd64.tar.gz -C /usr/local/bin && rm cilium-linux-amd64.tar.gz

# Install helm-3
RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 \
    && chmod +x get_helm.sh && ./get_helm.sh && rm ./get_helm.sh

## Get Kind configuration
RUN curl -L https://docs.cilium.io/en/v1.13/_downloads/5d189e574ddd393eb1bcf12c6937ce0f/kind-config.yaml > kind.yaml

# Remove IPv6 alias for localhost and start docker in the background ...
CMD cat /etc/hosts >/etc/hosts.bak && \
    sed 's/^::1.*//' /etc/hosts.bak > /etc/hosts && \
    sed -i "s/\PWD_IP_ADDRESS/$PWD_IP_ADDRESS/" /etc/docker/daemon.json && \
    sed -i "s/\DOCKER_TLSENABLE/$DOCKER_TLSENABLE/" /etc/docker/daemon.json && \
    sed -i "s/\DOCKER_TLSCACERT/$DOCKER_TLSCACERT/" /etc/docker/daemon.json && \
    sed -i "s/\DOCKER_TLSCERT/$DOCKER_TLSCERT/" /etc/docker/daemon.json && \
    sed -i "s/\DOCKER_TLSKEY/$DOCKER_TLSKEY/" /etc/docker/daemon.json && \
    mount -t securityfs none /sys/kernel/security && \
    echo "root:root" | chpasswd &> /dev/null && \
    /usr/sbin/sshd -o PermitRootLogin=yes -o PrintMotd=no 2>/dev/null && \
    dockerd &>/docker.log & \
    while true ; do script -q -c "/bin/bash -l" /dev/null ; done
# ... and then put a shell in the foreground, restarting it if it exits
