#!/bin/sh

wfr -e "WAIT_FOR_READINESS" -q

exec /bin/regioncn "$@"