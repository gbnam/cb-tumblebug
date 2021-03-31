#!/bin/bash

#function install_agent() {
    FILE=../conf.env
    if [ ! -f "$FILE" ]; then
        echo "$FILE does not exist."
        exit
    fi

	source ../conf.env
	AUTH="Authorization: Basic $(echo -n $ApiUsername:$ApiPassword | base64)"

	echo "####################################################################"
	echo "## Install monitoring agent to MCIS "
	echo "####################################################################"

	CSP=${1}
	REGION=${2:-1}
	POSTFIX=${3:-developer}

	source ../common-functions.sh
	getCloudIndex $CSP

	MCISID=${CONN_CONFIG[$INDEX,$REGION]}-${POSTFIX}
	curl -H "${AUTH}" -sX POST http://$TumblebugServer/tumblebug/ns/$NS_ID/monitoring/install/mcis/$MCISID -H 'Content-Type: application/json' -d \
		'{
			"command": "echo -n [CMD] Works! [Public IP: ; curl https://api.ipify.org ; echo -n ], [Hostname: ; hostname ; echo -n ]"
		}' | json_pp #|| return 1
#}

#install_agent