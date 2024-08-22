# HAWKEYE TOOLBOX

application for maintanance and watchout of [toolbox](https://github.com/carnivuth/toolbox) instances

## USE CASE

In my personal workflow i often install my personal [toolbox](https://github.com/carnivuth/toolbox) in a pletora of different systems with different users and i often forgot to remove them after the work is done, the idea with this service is to provide a simple but usefull api for registering toolbox installation and removal to keep track of what i have forgot in some environment

## FEATURES
- registration of new toolbox instances trough http requests through a identifier
- deletion of uninstalled toolbox instances
- api for retrieving current installed toolboxes
- storage on a sqlite db
- configuration loaded as env vars

## DATA FORMAT

the data format used follow this structure
```json
{
    "Timestamp":"timestamp of the installation",
    "User":"user that has installed the toolbox",
    "Hostname":"hostname of the machine",
    "Hash":"hash of User hostname and ps1",
    "Ps1":"ps1 variable"
}'
```
