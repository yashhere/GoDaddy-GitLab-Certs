#!/bin/bash

# https://stackoverflow.com/a/246128/5042046
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"

cd ${DIR}
echo "Updating DNS records for ${CERTBOT_DOMAIN}"
go run cmd/main.go "${CERTBOT_DOMAIN}" ${CERTBOT_VALIDATION} "goDaddy"

echo ${CERTBOT_VALIDATION}
echo
if [ "${CERTBOT_DOMAIN}" = "yashagarwal.in" ]; then
    # echo "Waiting for GoDaddy to publish the DNS records"
    for (( i=0; i<9; i++ )); do
        # echo "Minute" ${i}
        sleep 60s
    done
fi