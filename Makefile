
.PHONY: protos

all:
	bash ./scripts/build_checks.sh
protos: 
	cd smartbftprotos && protoc -I=. \
	-I${GOPATH}/src \
	--gogofaster_out=:. \
	messages.proto logrecord.proto