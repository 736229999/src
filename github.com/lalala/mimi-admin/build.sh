#!/bin/bash
if [ $# -eq 0 ]; then
    echo "miss GOOS(linux/darwin)"
    exit 1
fi

GOOS=$1
if [[ "$GOOS" != "linux" && "$GOOS" != "darwin" ]]; then
    echo "invalid GOOS $GOOS"
    exit 1
fi

BIN=`pwd`/bin/$GOOS
if [ ! -d $BIN ]; then
    mkdir -p $BIN
fi

echo "GOOS: ${GOOS}"

APPS=(backend dbadminagent frontend)

goBuild() {
    app=$1
    printf "%s" $app
    pushd . > /dev/null
    cd "$app"
    OUT="${app}.$GOOS"
    GOOS=$GOOS GOARCH=amd64 go build -o "${BIN}/$OUT"
    popd > /dev/null
    printf " --> %s\n" $OUT
}

npmBuild() {
    app=$1
    printf "%s" "npm run build..."
    pushd . > /dev/null
    cd "$app"
    npm run build
    tar zcvf dist.tar.gz dist
    mv dist.tar.gz $BIN
    popd > /dev/null
}

for app in ${APPS[@]}; do
    if [ $app == "frontend" ]; then
	npmBuild $app
    else
	goBuild $app
    fi
done
