#!/usr/bin/bash

BIN_PATH=bin/onewallet

echo "+ Building the project"
go build -o $BIN_PATH -v
echo " - Build completed"

echo "+ Creating symlink"
if [ -h ~/.local/$BIN_PATH ]; then
    echo " - Symlink already exists"
else
    ln -s "$(pwd)/$BIN_PATH" ~/.local/$BIN_PATH
    echo " - Symlink created"
fi
