#!/bin/sh

docktool --quiet wait -e="WAIT_"

exec /bin/regioncn "$@"