#!/bin/sh

docktool wait -e="WAIT_"

exec /bin/regioncn "$@"