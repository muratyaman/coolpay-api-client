# coolpay-api-client

my first project in Go language! a RESTful API acting as a HTTP client to Coolpay API:

[http://docs.coolpayapi.apiary.io/](http://docs.coolpayapi.apiary.io/)

I don't remember the last time I used a redhat-based linux distro!

## requirements

### fedora 24 box using vagrant

running on a vagrant box using fedora 24

`cd ~`

`mkdir vagrant-fedora24`

`cd vagrant-fedora24`

`vagrant init fedora/24-cloud-base`

`vagrant up`

`vagrant ssh`

### go language

`sudo dnf install go`

edit .bash_profile

`vi .bash_profile`

include the following:


> GOPATH=/usr/local/go/bin

> PATH=\$PATH:\$HOME/.local/bin:\$HOME/bin:\$GOPATH

> export PATH

> export GOPATH=\$HOME/gocode



logout and login again

`exit`

`vagrant ssh`

### git

`sudo dnf install git`

### beego framework

`go get github.com/astaxie/beego`

`go install  github.com/astaxie/beego`

`go get github.com/beego/bee`

### clone this repo

`git clone https://github.com/muratyaman/coolpay-api-client.git`

`cd coolpay-api-client`


