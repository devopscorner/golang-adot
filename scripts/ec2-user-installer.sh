#!/usr/bin/env sh

# Path xterm-256-color
export TERM="xterm-256color"

export ANSIBLE_VERSION=2.12.2
# export ANSIBLE_TOWER_CLI_VERSION=3.3.9   # Depreciated, change to AWX-Cli (awxkit)
export AWXKIT_VERSION=21.9.0
export PACKER_VERSION=1.8.6
export TERRAFORM_VERSION=1.3.9
export TERRAGRUNT_VERSION=0.43.2
export TERRASCAN_VERSION=1.18.0
export HELMFILE_VERSION=0.144.0
export KUBECTL_VERSION=1.26.3
export K9S_VERSION=0.27.3
export CHECKOV_VERSION=2.3.21
export AWS_CLI_VERSION=2.11.0
export TFSEC_VERSION=1.28
export VERIFY_CHECKSUM=false
export GOLANG_VERSION=1.19.6

export DOCKER_PATH="/usr/bin/docker"
export DOCKER_COMPOSE_PATH="/usr/bin/docker-compose"
export DOCKER_COMPOSE_VERSION="1.29.2"
export DOCKER_BUILDX_VERSION="0.10.3"

sudo yum update -y
sudo yum groupinstall "Development Tools" -y

sudo yum install -y \
      git \
      nano \
      curl \
      zip \
      unzip \
      wget \
      tmux \
      zsh \
      mc

sudo yum install -y \
      bash \
      jq \
      vim-enhanced \
      telnet \
      openssl11 \
      openssl11-devel \
      libffi-devel \
      bzip2-devel \
      python37 \
      python3-pip \
      golang

# ==================== #
#  Install Terragrunt  #
# ==================== #
sudo wget -O /usr/local/bin/terragrunt \
      https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 &&\
      chmod +x /usr/local/bin/terragrunt &&\
      # ================ #
      #  Install Packer  #
      # ================ #
      wget -O packer_${PACKER_VERSION}_linux_amd64.zip \
      https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_amd64.zip &&\
      sudo unzip packer_${PACKER_VERSION}_linux_amd64.zip -d /usr/local/bin &&\
      rm -f packer_${PACKER_VERSION}_linux_amd64.zip &&\
      # =================== #
      #  Install Terrascan  #
      # =================== #
      wget -O terrascan.tar.gz \
      https://github.com/accurics/terrascan/releases/download/v${TERRASCAN_VERSION}/terrascan_${TERRASCAN_VERSION}_Linux_x86_64.tar.gz &&\
      sudo tar -zxf terrascan.tar.gz -C /usr/local/bin &&\
      sudo chmod +x /usr/local/bin/terrascan &&\
      rm terrascan.tar.gz &&\
      # =================== #
      #  Install Infracost  #
      # =================== #
      sudo curl -fsSL https://raw.githubusercontent.com/infracost/infracost/master/scripts/install.sh | bash

# =============== #
#  Install TFSec  #
# =============== #
sudo curl -s https://raw.githubusercontent.com/aquasecurity/tfsec/master/scripts/install_linux.sh | bash

# =============== #
#  Install TFenv  #
# =============== #
git clone https://github.com/tfutils/tfenv.git $HOME/.tfenv &&\
      echo 'export PATH="$HOME/.tfenv/bin:$PATH"' >> $HOME/.bashrc &&\
      ln -sf $HOME/.tfenv/bin/* /usr/local/bin &&\
      mkdir -p $HOME/.local/bin/ &&\
      ln -sf $HOME/.tfenv/bin/* $HOME/.local/bin

# =================== #
#  Install Terraform  #
# =================== #
/usr/local/bin/tfenv install ${TERRAFORM_VERSION} &&\
    /usr/local/bin/terraform use ${TERRAFORM_VERSION} &&\
    /usr/local/bin/terraform version

# ============== #
#  Install Helm  #
# ============== #
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 &&\
      chmod +x get_helm.sh &&\
      export VERIFY_CHECKSUM=${VERIFY_CHECKSUM} &&\
      ./get_helm.sh &&\
      helm version

# ===================== #
#  Install Helm Plugins #
# ===================== #
helm repo add stable https://charts.helm.sh/stable

helm plugin install https://github.com/databus23/helm-diff &&\
      helm plugin install https://github.com/futuresimple/helm-secrets &&\
      helm plugin install https://github.com/hypnoglow/helm-s3.git &&\
      helm plugin install https://github.com/aslafy-z/helm-git.git &&\
      helm plugin install https://github.com/rimusz/helm-tiller &&\
      helm repo update

# ================== #
#  Install Helmfile  #
# ================== #
sudo wget -O /usr/local/bin/helmfile \
      https://github.com/roboll/helmfile/releases/download/v${HELMFILE_VERSION}/helmfile_linux_amd64 &&\
      sudo chmod +x /usr/local/bin/helmfile &&\
      helmfile --version

# ================= #
#  Install Kubectl  #
# ================= #
sudo wget -O /usr/local/bin/kubectl \
      https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl &&\
      sudo chmod +x /usr/local/bin/kubectl

# ============= #
#  Install K9S  #
# ============= #
sudo wget -O /usr/local/bin/k9s \
      https://github.com/derailed/k9s/releases/download/v${K9S_VERSION}/k9s_Linux_amd64.tar.gz &&\
      sudo chmod +x /usr/local/bin/k9s

# ====================== #
#  Install Tmux Plugins  #
# ====================== #
git clone https://github.com/tmux-plugins/tpm $HOME/.tmux/plugins/tpm
cp config/.tmux.conf $HOME
cp config/tmuxcolors.tmux $HOME

# ============= #
#  Install GVM  #
# ============= #
sudo yum install go -y

curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash
gvm install go${GOLANG_VERSION}; sync
gvm use go${GOLANG_VERSION} --default; sync

# ================ #
#  Install Docker  #
# ================ #
sudo amazon-linux-extras install docker

# ======================= #
#  Install Docker Buildx  #
# ======================= #
curl --silent -L https://github.com/docker/buildx/releases/download/v${DOCKER_BUILDX_VERSION}/buildx-v${DOCKER_BUILDX_VERSION}.linux-amd64 -o buildx-v${DOCKER_BUILDX_VERSION}.linux-amd64
chmod a+x buildx-v${DOCKER_BUILDX_VERSION}.linux-amd64
mkdir -p $HOME/.docker/cli-plugins
mv buildx-v${DOCKER_BUILDX_VERSION}.linux-amd64 $HOME/.docker/cli-plugins/docker-buildx

# ========================= #
#  Install Docker-Compose   #
# ========================= #
sudo curl -L https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m) -o ${DOCKER_COMPOSE_PATH}
sudo chmod +x /usr/bin/docker-compose

touch $HOME/.zshrc
echo '' >> $HOME/.zshrc
echo '### Docker ###
export DOCKER_CLIENT_TIMEOUT=300
export COMPOSE_HTTP_TIMEOUT=300' >> $HOME/.zshrc

##### CONFIGURE DOCKER #####
sudo usermod -a -G docker ec2-user

# ======== #
#  Others  #
# ======== #
## Set Locale
sudo touch /etc/environment
sudo echo 'LANG=en_US.utf-8' >> /etc/environment
sudo echo 'LC_ALL=en_US.utf-8' >> /etc/environment

## Adding Custom Sysctl
sudo echo 'vm.max_map_count=524288' >> /etc/sysctl.conf
sudo echo 'fs.file-max=131072' >> /etc/sysctl.conf

# =================== #
#  Install AWSCli v2  #
# =================== #
curl -s https://awscli.amazonaws.com/awscli-exe-linux-x86_64-${AWS_CLI_VERSION}.zip -o awscliv2.zip &&\
      unzip awscliv2.zip &&\
      ./aws/install --bin-dir /usr/local/bin/

# =================== #
#  Install Oh-My-ZSH  #
# =================== #
sudo yum install util-linux-user -y
sudo chsh -s $(which zsh) $(whoami)

sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# ========================= #
#  Install Oh-My-ZSH Theme  #
# ========================= #
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k
cp config/.p10k.zsh $HOME
