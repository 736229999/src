#!/bin/sh
pushd . > /dev/null
cd source
PROTO_FILES=""
for file in *.proto
do
	echo $file
	if test -f $file
	then
		PROTO_FILES="$PROTO_FILES $file"
	fi
done

makeGolang()
{
        protoc -I . $PROTO_FILES --go_out=plugins=grpc:../proto
}

makeJava()
{
        protoc $PROTO_FILES --java_out="../java"
}

makeSwift()
{
        protoc $PROTO_FILES --swift_out="../swift"
}

if [ $# -eq 0 ]; then  
        makeGolang
        makeJava
        makeSwift
else
        if [ $1 = 'go' ]; then
                makeGolang
        elif [ $1 = 'java' ]; then
                makeJava
        elif [ $1 = 'swift' ]; then
                makeSwift
        fi
fi 

popd > /dev/null
ls ./proto/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
