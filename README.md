# GoDaddy-GitLab-Certs

A set of scripts to update the HTTPS certificates of my website [https://yashagarwal.in](https://yashagarwal.in) using GoDaddy and GitLab APIs.

## Usage Instructions

First of all, set the four environmental variables `EMAIL_ID`, `GODADDY_KEY`, `GODADDY_SECRET` and `GITLAB_TOKEN` in your shell.

To generate API keys for GoDaddy, sign into your account and follow the instructions on [this](https://developer.godaddy.com/keys) page.

To generate personal access token for GitLab API, follow the instructions given [here](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#creating-a-personal-access-token).

To run these scripts, you will require [Certbot](https://certbot.eff.org/) and [GoLang](https://github.com/golang/go) installed on your machine.

Once you have finished setting up your system, clone the repository and navigate to it. 

In the file `certbot.sh`, edit the `certbot` command to include your domain names.

In the subsequent `curl` command in the same file, edit the URL part to include your repository address where your webiste's code sits. Follow the documentation of GitLab Pages Domain API [here](https://docs.gitlab.com/ee/api/pages_domains.html).

Now, in the `auth_hook.sh` file, towards the end, edit the `if` statement to compare the `CERTBOT_DOMAIN` variable to the the domain name, which is the last in the list that you defined in the certbot command above. For example -

```bash
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
-d photos.yashagarwal.in \
-d yashagarwal.in \
-d readings.yashagarwal.in \ # THIS IS THE LAST DOMAIN DEFINED
certonly
```

So, in the `auth_hook.sh` file -

```bash
# Comparing with the last domain
if [ "${CERTBOT_DOMAIN}"  =  "readings.yashagarwal.in" ];  then
	for  (( i=0; i<5; i++ ));  do
		sleep 60s
	done
fi
```

Domain registrars require some time to publish the changes in the DNS records. Here I have assumed 5 minutes in case of GoDaddy, and it works fine for me. Change the time limit according to your experience.

Also update the `key_dir` vairable in the `certbot.sh` file to include your last domain name as explained above. So in the above case, key_dir variable will be -

```bash
key_dir="${DIR}/generated/config/live/readings.yashagarwal.in"
```

Now, while you are in the root of the directory, execute the following command - 
```bash
bash certbot.sh
```

If everything works as expected, then you will see following message -
```text
Congratulations! Your certificate and chain have been saved at: ...
```

The script will automatically update your https certificate on the GitLab pages also. Your website will reflect the changes in some time.

## Authors

* **Yash Agarwal** - *Initial work* - [Pallav Agarwal](https://github.com/pallavagarwal07) 

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org) This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

This script by [Pallav Agarwal](https://github.com/pallavagarwal07/NamesiloCert) was the initial inspiration for this work. You can observe a lot of similarities in both the codes. Basically, I have just modified his code to work with GoDaddy API.