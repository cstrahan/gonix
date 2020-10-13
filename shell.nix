with import <nixpkgs> { };

runCommand "dummy" {
  buildInputs = [
    #go
    stdenv openssl pkgconfig
    cyrus_sasl # for mgo
    ragel
    stdenv
    stdenv.cc
  ];
  shellHook = ''
    unset SSL_CERT_FILE
    #export GOPATH=$(readlink -f ../../../..):${go}/share/go
    PATH=$(readlink -f ../../bin):$PATH
  '';
} ""

/*
Stuff used by vim-go:
---------------------
gocode
gometalinter
goimports
godef
oracle
gorename
golint
errcheck
gotags
asmfmt
motion
*/
