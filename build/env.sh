#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
ethdir="$workspace/src/github.com/akroma-project"
if [ ! -L "$ethdir/akroma" ]; then
    mkdir -p "$ethdir"
    cd "$ethdir"
    ln -s ../../../../../. akroma
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

GO111MODULE=auto
export GO111MODULE
# Run the command inside the workspace.
cd "$ethdir/akroma"
PWD="$ethdir/akroma"

# Launch the arguments with the configured environment.
exec "$@"
