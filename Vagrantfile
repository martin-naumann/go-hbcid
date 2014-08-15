# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "debian73"
  config.vm.box_url = "http://puppet-vagrant-boxes.puppetlabs.com/debian-73-x64-virtualbox-nocm.box"

  config.vm.define "server" do |server|
    server.vm.hostname = "go-hbcid-server"
    server.vm.network :private_network, ip: "10.10.10.2"
    server.vm.provision "shell", path: "provisioning/setup_server.sh"
  end
end
