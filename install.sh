#!/bin/bash

release_url=$(curl -s https://api.github.com/repos/miguelmota/go-webhook-server/releases/latest | grep "browser_download_url.*Linux_x86_64.tar.gz" | cut -d : -f 2,3 | tr -d \")

(
  cd /tmp
  wget $release_url -O gws.tar.gz

  tar -xvzf gws.tar.gz gws
  chmod +x gws
  mv gws /usr/local/bin/gws
  rm -rf gws gws.tar.gz
  echo "Installed in $(which gws)"
)