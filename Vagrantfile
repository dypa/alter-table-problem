Vagrant.configure("2") do |config|
  config.vm.box = "debian/wheezy64"
  config.vm.box_check_update = false

  config.vm.network "private_network", ip: "192.168.50.56"
  config.vm.synced_folder "./", "/home/project", mount_options: ["dmode=777","fmode=666"]

  config.vm.provider "virtualbox" do |vb|
    vb.cpus = 3
    vb.memory = 2024
  end

  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    wget https://repo.percona.com/apt/percona-release_0.1-4.$(lsb_release -sc)_all.deb
    dpkg -i percona-release_0.1-4.$(lsb_release -sc)_all.deb
    apt-get install -y python-software-properties
    apt-key adv --recv-keys --keyserver keyserver.ubuntu.com 0xcbcb082a1bb943db
    add-apt-repository 'deb [arch=amd64,i386] http://mirror.timeweb.ru/mariadb/repo/5.5/debian wheezy main'
    apt-get update
    debconf-set-selections <<< 'mariadb-server-5.5 mysql-server/root_password password dbpass'
    debconf-set-selections <<< 'mariadb-server-5.5 mysql-server/root_password_again password dbpass'
    apt-get install -y mariadb-server
    apt-get install -y percona-toolkit
  SHELL
end
