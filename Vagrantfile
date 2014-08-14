# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "precise64"
  config.vm.box_url = "http://images.cwc.io/2014/06/debian75-8GB-20140611.box"

  config.vm.define "riak" do |riak|
    riak.vm.hostname = "go-hbcid-riak"
    riak.vm.network :private_network, ip: "176.17.17.11"
    riak.vm.provision "shell", path: "provisioning/setup_riak.sh"
  end

  config.vm.define "server" do |server|
    server.vm.hostname = "go-hbcid-server"
    server.vm.network :private_network, ip: "176.17.17.20"
    server.vm.provision "shell", path: "provisioning/setup_server.sh"
  end
end
