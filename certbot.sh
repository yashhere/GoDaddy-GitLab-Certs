#!/bin/bash

if ! [ -x "$(command -v certbot)" ]; then
    echo 'Error: certbot is not installed.' >&2
    exit 1
fi

# https://stackoverflow.com/a/246128/5042046
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"

if [[ -z "${EMAIL_ID}" ]]; then
    echo "Email ID is required to run certbot. Please set environment variable EMAIL_ID to continue"
    exit 1
fi

rm -rf ${DIR}/generated/{config,work,logs}
mkdir -p ${DIR}/generated/{config,work,logs}

certbot --manual \
--preferred-challenges dns \
--agree-tos \
--email "${EMAIL_ID}" \
--no-eff-email \
--expand \
--renew-by-default \
--manual-public-ip-logging-ok \
--noninteractive \
--redirect \
--config-dir ${DIR}/generated/config \
--work-dir ${DIR}/generated/work \
--logs-dir ${DIR}/generated/logs \
--manual-auth-hook ${DIR}/auth_hook.sh \
-d yashagarwal.in \
certonly

key_dir="${DIR}/generated/config/live/yashagarwal.in"

curl -vvv \
--request PUT \
--header "Private-Token:${GITLAB_TOKEN}" \
--form "certificate=@${key_dir}/fullchain.pem" \
--form "key=@${key_dir}/privkey.pem" \
"https://gitlab.com/api/v4/projects/yashhere%2Fyashhere.gitlab.io/pages/domains/yashagarwal.in" > ${DIR}/generated/logs/gitlab.log 2>&1