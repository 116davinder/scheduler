Vagrant.configure("2") do |config|

  config.vm.define "golang" do |golang|
    golang.vm.box = "geerlingguy/centos7"
    golang.vm.hostname = "golang"
    golang.vm.network :private_network, ip: "192.168.56.100"
    golang.vm.provision "shell",
      inline: 'rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO && curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo && yum install golang git -y'
  end
  config.vm.provider "virtualbox" do |vb|
    vb.gui = false
    vb.memory = "256"
    vb.cpus =  1
  end
  config.vm.synced_folder ".", "/home/vagrant/projects", type: "nfs"
  config.ssh.insert_key = false

end
