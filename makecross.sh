#!/bin/env sh

OSS="windows linux darwin"
ENTRYPOINT="./cmd/treport-gui"

echo "Cross compiling"
echo "    ..."

for OS in $OSS; do
    echo
    echo "\$> fyne-cross $OS $ENTRYPOINT"
    echo
    fyne-cross "$OS" "$ENTRYPOINT"
done
