# Go project

## Version numbering
Following the conventions on [Semantic Versioning](https://semver.org/).

## Environmental variables

Configuration parameters are stored as environment variables. In Windows they can be managed with sysdm.cpl.

| Key | Value |
| --- | --- |
| open311MongoURI | MongoDB Atlas connection string in mongodb+srv version |
| open311SentryDSN | Sentry telemetry connection |
| open311port | Port of the listener, e.g. 8080 |
| open311_SSH_CERT | Path to SSH cert |
| open311_SSH_KEY | Path to SSH public key |
| open311_LOGLOCAL | If 1, write local log file on open311_LOGPATH |
| open311_LOGPATH | Path including filename to local log (e.g. /var/log/open311) |


>[!TIP]
>While it is year 2024, it still might be required to reboot your Windows 10 after setting environmental variables. 

## Driver Packages

```
go get -u go.uber.org/zap     
go get github.com/google/uuid     
go get github.com/cloudflare/cfssl/cmd/cfssl  
go get github.com/cloudflare/cfssl/cmd/cfssl    
go get github.com/cloudflare/cfssl/cmd/cfssljson
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/collections
go get go.mongodb.org/mongo-driver/options
go get github.com/hashicorp/serf/serf@v0.10.1
go get github.com/joho/godotenv
```