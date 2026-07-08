# Github Integration

# Work in progress

## Overview

**Github Integration**

## Configuration

- `token` (required) The github token that will be used to authenticate to github's API.

## Examples

```sh
# Create a github configuration
$ plakar source add my-github location=github://. token=$GITHUB_TOKEN organization=MyOrg
# Backup my github
$ plakar backup "@my-github"
```