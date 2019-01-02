<!-- # compile go from source

1. go/gccgo
2. supported arch `amd64/386/arm/arm64/ppc64/ppc64le/mips/mipsle/mips64/ips64le/s39x`
3. build the bootstrap(go1.4)/crosscompile
4. fetch the source

    ``` sh
    git clone https://go.googlesource.com/go
    ```
5. checkout branch
6. build

    ``` sh
    cd src
    ./all.bash
    ```
environment

    $GOOS
    $GOARCH
    $GOROOT_BOOTSTRAP -->

# [Install go](https://golang.org/doc/install)

Download link: `https://golang.org/dl/`

``` sh
wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
# creating a Go tree in /usr/local/go
sudo -s
tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz

# Add /usr/local/go/bin to the PATH environment variable.
export PATH=$PATH:/usr/local/go/bin
export GOPATH=~/go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

<!-- ## Installing to a custom location

``` sh
# need to set the following
export GOROOT=$HOME/go1.X
export PATH=$PATH:$GOROOT/bin
``` -->

Uninstalling Go

``` sh
rm -rf /usr/local/go
# Remove  PATH environment variable
```