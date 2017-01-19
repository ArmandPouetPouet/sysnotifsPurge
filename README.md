# sysnotifsPurge

WebService with : 
- /index route to display a "I'm alive" message
- /CheckRatio route to clean AWS sysnotifs ElasticSearch cluster

Cleaning means : 
1°) Request data from ES
2°) Ratio space occupied and available and compare it to a threshold
3°) Delete oldest indices until ratio gets back under threshold

Service must be deployed on a static IP, since AWS ES access is locked with IP.

TODO : 
- Have clean env variables
- Some unit testing
- Compare code to best practices
