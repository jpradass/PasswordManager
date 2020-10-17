# PasswordManager
Tool to manage passwords and keep them safe and encrypted in a remote mongo database. 

## How to use it
You could download the latest release as an executable for your OS, or you can clone this repo and compile from the code. 
Then, we only have to rename the binary as "pm" or some another name, and place it, along with configuration folder with conf.yaml inside, inside our %PATH to execute it from everywhere in our pc.

We have to make sure we have a remote MongoDB deployed (you can deploy a local instance or even in Docker as well) with a database and a collection ready to use. Mongo Atlas has a free tier remote MongoDB enough for the purpose of using this. 

Here's a list of commands and subcommands 

* init
* list
* help
* get username|password service
* set service username password
* update username|password service
* delete service


## TO DO List
1. Configure a time range for this service to warn if a password overpass that limit
2. Create a password generator based on a list of common words from internet (better than random pc generated passwords)
3. Store services.json in a service as AWS S3 or whatever that could be accessible everywhere, independent of my pc
4. Flush command to remove all services stored and a delete command to remove only one service
