# Go example app using Gin web framework

## Deploy to TAP

tanzu apps workload apply go-gin --git-repo https://github.com/markusrt/go-gin --git-branch main --type web --build-env "BP_KEEP_FILES=templates/*"