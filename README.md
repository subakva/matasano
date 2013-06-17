matasano crypto challenge
========

Solutions for the Matasano Crypto Challenge in Go.

From private repo:

    $ cd $GOPATH
    $ git clone git@github.com:subakva/matasano.git src/subakva/matasano
    $ pushd src; find subakva/matasano -type d -name "problem*" | xargs go test; popd

From a tarball:

    $ pushd $GOPATH/src
    $ tar xzf subakva.tar.gz
    $ popd
    $ pushd src; find subakva/matasano -type d -name "problem*" | xargs go test; popd

