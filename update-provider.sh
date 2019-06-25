#!/bin/bash

if [ -z "${1}" ]
then
  # Get latest version
  version="$(curl -sL https://api.github.com/repos/rhythmictech/terraform-provider-errorcheck/releases/latest | jq -r '.name')"
else
  # Get specific version
  version="${1}"
fi

# Delete old versions
rm terraform.d/plugins/linux_amd64/*
rm terraform.d/plugins/darwin_amd64/*

# Download Linux provider
curl -sSL -o \
  "terraform.d/plugins/linux_amd64/terraform-provider-errorcheck_v${version}" \
  "https://github.com/rhythmictech/terraform-provider-errorcheck/releases/download/${version}/terraform-provider-errorcheck_v${version}_linux_amd64"
chmod 0744 terraform.d/plugins/linux_amd64/*

# Download Mac provider
curl -sSL -o \
  "terraform.d/plugins/darwin_amd64/terraform-provider-errorcheck_v${version}" \
  "https://github.com/rhythmictech/terraform-provider-errorcheck/releases/download/${version}/terraform-provider-errorcheck_v${version}_darwin_amd64"
chmod 0744 terraform.d/plugins/darwin_amd64/*
