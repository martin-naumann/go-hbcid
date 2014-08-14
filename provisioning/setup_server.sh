sudo apt-get -y update
sudo apt-get install -y git mercurial

wget https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz
tar xvf go1.2.linux-amd64.tar.gz
sudo mv go /opt/

sudo ln -s /opt/go/bin/go /usr/local/bin/go
sudo ln -s /opt/go/bin/gofmt /usr/local/bin/gofmt
sudo ln -s /opt/go/bin/godoc /usr/local/bin/godoc

sudo echo "export GOROOT=/opt/go" >> /etc/profile
sudo echo "export GOPATH=~/go" >> /etc/profile
