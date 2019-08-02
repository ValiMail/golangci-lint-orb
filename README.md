# valimail/golangci-lint - CircleCI Orb

[![CircleCI](https://circleci.com/gh/ValiMail/golangci-lint-orb.svg?style=svg&circle-token=540455df00def7479a3c81c9ab9ac8ab4f810178)](https://circleci.com/gh/ValiMail/golangci-lint-orb)

CircleCI Orb for managing dependencies like Ruby gems and JS packages.

Published at https://circleci.com/orbs/registry/orb/valimail/golangci-lint

## Development

Work on files in the `src` directory and when you're ready you can "pack",
"validate", and "publish" the orb.

```bash
# Build orb.yml from files in src folder
circleci config pack src > orb.yml

# Validate orb.yml
circleci orb validate orb.yml

# For a mutable "dev" version, increment DEV_ORB_VERSION (below) and publish orb.yml.
DEV_ORB_VERSION=dev:0.2.0
circleci orb publish orb.yml valimail/golangci-lint@$DEV_ORB_VERSION

# For an immutable release, increment ORB_VERSION (below) and publish orb.yml.
# Note: only GitHub admins of ValiMail can do this currently.
ORB_VERSION=0.2.0
circleci orb publish orb.yml valimail/golangci-lint@$ORB_VERSION

# Update version for tests in .circlci/config.yml
```

## Setup

This was created using the following commands:

```bash
# Initial setup of valimail namespace
circleci namespace create valimail github ValiMail

# Create this orb
circleci orb create valimail/golangci-lint
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/ValiMail/golangci-lint-orb.

## Credit

Forked from https://github.com/timakin/golangci-lint-orb
