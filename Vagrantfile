# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "bento/ubuntu-16.04"

  config.vm.network "private_network", type: "dhcp"

  config.vm.define "APPSERVER", primary: true do |server|
    server.vm.network "private_network", ip: "192.168.20.2"
    server.vm.network "forwarded_port", guest: 8866, host: 8866
    server.vm.provider "virtualbox" do |vb|
      vb.memory = "1024"
      vb.cpus = "1"
    end
    server.vm.hostname = "APPSERVER"
    server.vm.provision "shell", inline: <<-SHELL
      echo "Hej from server one!" > /var/www/html/index.html
	  echo "====================================="
	  echo "Installing Go... Even thos this might even be needed, since its apart of the golang:jessie image"
	  echo "======================================="
	  wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
      sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
      echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile
      echo "export GOPATH=/go_projects" >> $HOME/.profile
	  echo "=============================="
	  echo "Installing Docker..."
	  echo "=============================="
	  sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
	  sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
	  sudo apt-get update
	  sudo apt-get install -y docker-ce
	  echo "=============================="
	  echo "Installing Docker-Compose..."
	  echo "=============================="
	  sudo -i $(curl -L https://github.com/docker/compose/releases/download/1.15.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose)
	  sudo chmod +x /usr/local/bin/docker-compose
	  echo "============================================"
	  echo "Building Dockerimage localy with Dockerfile"
	  echo "============================================"
	  sudo docker build -t mysite /vagrant
	  echo "============================================"
	  echo "Dockering-up from docker-compose file."
	  echo "============================================"
	  echo "============================================"
	  echo "============================================"
	  echo "===========DONE LOADING SERVER=============="
	  echo "=========192.168.20.2:8866 - Webapp========="
	  echo "============================================"
	  echo "============================================"
	  cp /vagrant/docker-compose.yml /home/vagrant/docker-compose.yml
	  sudo docker-compose up -d
    SHELL
  end

  config.vm.define "DBSERVER" do |client|
    client.vm.network "private_network", ip: "192.168.20.3"
    client.vm.network "forwarded_port", guest: 27017, host: 27017
    client.vm.provider "virtualbox" do |vb|
      vb.memory = "1024"
      vb.cpus = "1"
    end
    client.vm.hostname = "DBSERVER"
    client.vm.provision "shell", inline: <<-SHELL
      echo "Hej from server 2!" > /var/www/html/index.html
	  echo "Installing MongoDB"
      sudo apt-get -y install mongodb-server
	  sudo sed -i '/bind_ip = / s/127.0.0.1/0.0.0.0/' /etc/mongodb.conf
	  sudo service mongodb restart
	  echo "============================================"
	  echo "============================================"
	  echo "===========DONE LOADING SERVER=============="
	  echo "=========192.168.20.3:27017 - DB============"
	  echo "============================================"
	  echo "============================================"
    SHELL
  end

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    sudo apt-get update
    sudo apt-get -y install apache2  
  SHELL
end
