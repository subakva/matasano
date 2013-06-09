matasano
========

Solutions for the Matasano Crypto Challenge in Go.

    $ cd $GOPATH
    $ git clone git@github.com:subakva/matasano.git src/subakva/matasano
    $ go install subakva/matasano/mcrypto && bin/mcrypto
    $ pushd src; find subakva/matasano -type d -name "problem*" | xargs go test; popd
