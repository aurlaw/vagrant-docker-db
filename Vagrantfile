VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(2) do |config|
  
  config.vm.box = "ubuntu/trusty64"
  config.vm.network "private_network", ip: "192.168.20.100"
  config.vm.hostname = "vagrant-docker-db"

  config.vm.provision :docker
  config.vm.provision :docker_compose, yml: "/vagrant/docker-compose.yml", rebuild: false, run: "always"

end