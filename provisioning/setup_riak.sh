# Get the riak apt repository plus key
curl http://apt.basho.com/gpg/basho.apt.key | sudo apt-key add -
sudo bash -c "echo deb http://apt.basho.com $(lsb_release -sc) main > /etc/apt/sources.list.d/basho.list"

# Install riak
sudo apt-get update -y
sudo apt-get install -y riak

# Adjust open file limit to 4096 (recommended for riak)
sudo echo "root        nofiles    4096" >> /etc/security/limits.conf
sudo echo "vagrant     nofiles    4096" >> /etc/security/limits.conf
sudo echo "riak        nofiles    4096" >> /etc/security/limits.conf

# Configure it and fire it up
sudo sed -i.bak "s/127\.0\.0\.1/0.0.0.0/" /etc/riak/app.config
sudo /etc/init.d/riak restart
