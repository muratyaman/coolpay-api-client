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

`sudo dnf update`

### mysql community server 5.7

`sudo dnf install mysql-community-server-5.7.15-1.fc24.x86_64`

### go language

`sudo dnf install go`

`mkdir ~/gocode`

`mkdir ~/gocode/src`

`mkdir ~/gocode/bin`

edit .bash_profile

`vi .bash_profile`

include the following:


> GOPATH=$HOME/gocode

> PATH=$PATH:$HOME/.local/bin:$HOME/bin:$GOPATH/bin

> export PATH

> export GOPATH


logout and login again

`exit`

`vagrant ssh`

### git

`sudo dnf install git`

### beego framework, bee tool, uuid library

`go get github.com/astaxie/beego`

`go install github.com/astaxie/beego`

`go get github.com/beego/bee`

`go get github.com/satori/go.uuid`

### clone this repo

`cd ~`

`git clone https://github.com/muratyaman/coolpay-api-client.git`

`cd gocode/src`

`ln -s ~/coolpay-api-client/myapi myapi`


### create database

check and run the SQL files in folder db-migration


### run app

`bee run`
