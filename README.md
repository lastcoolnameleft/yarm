# YARM = YAML + ARM

This is my first "real" Go program.  I can't stop you from laughing at my code, but I sure would appreciate constructive feedback instead.

## Getting started

```shell
# Install YARM
go install github.com/lastcoolnameleft/yarm
cd $GOPATH/src/github.com/lastcoolnameleft/yarm

# Build a template from the sample
yarm generate content/values/simple-linux-vm.yaml -s $SUBSCRIPTION_ID -g $RESORUCE_GROUP -o json > tmp/vm.json

# Deploy the template
az group create -n $RESOURCE_GROUP -l $LOCATION
az group deployment create --template-file tmp/vm.json -g $RESOURCE_GROUP

# Verify the deployment
az vm list -g $RESOURCE_GROUP
```