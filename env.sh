if type go >/dev/null; then
	:
else
    export PATH=$PATH:/usr/local/go/bin
fi
export GOPATH=$PWD
export PATH=$PATH:$PWD/bin

# https://github.com/coopernurse/gorp
export GORP_TEST_DSN=test/testuser/TestPasswd9
export GO_TEST_DSN=testuser:TestPasswd9@/test
