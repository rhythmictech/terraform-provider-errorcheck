#!/bin/bash


linux_plugin_dir="terraform.d/plugins/linux_amd64/"
mac_plugin_dir="terraform.d/plugins/darwin_amd64"

if [ -z "${1}" ]
then
  # Default to 2.0.3
  version="2.0.3"
else
  # Get specific version
  version="${1}"
fi

# Delete old versions
[ -d $linux_plugin_dir ] && rm $linux_plugin_dir/*
[ -d $mac_plugin_dir ] && rm $mac_plugin_dir/*

# if it doesn't exist
# then create plugin dir
[ ! -d $linux_plugin_dir ] && mkdir -p $linux_plugin_dir
[ ! -d $mac_plugin_dir ] && mkdir -p $mac_plugin_dir

# Download Linux provider
curl -sSL -o \
  "$linux_plugin_dir/terraform-provider-errorcheck_v${version}" \
  "https://github.com/rhythmictech/terraform-provider-errorcheck/releases/download/${version}/terraform-provider-errorcheck_v${version}_linux_amd64"
chmod 0744 $linux_plugin_dir/*

# Download Mac provider
curl -sSL -o \
  "$mac_plugin_dir/terraform-provider-errorcheck_v${version}" \
  "https://github.com/rhythmictech/terraform-provider-errorcheck/releases/download/${version}/terraform-provider-errorcheck_v${version}_darwin_amd64"
chmod 0744 $mac_plugin_dir/*
