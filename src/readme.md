# Go project

## Version numbering
Following the conventions on [Semantic Versioning](https://semver.org/).

## Environmental variables

Configuration parameters are stored as environment variables (.env).

| Key | Value |
| --- | --- |
| BUILD_DATE | Date of build. |
| BUILD_NUMBER | Sequential number of build. |
| BUILD_ENV | Target environment: PROD/TEST/DEV |
| open311MongoURI | MongoDB Atlas connection string in mongodb+srv version |
| open311SentryDSN | Sentry telemetry connection |
| open311port | Port of the listener as string, e.g. :8080 (note the colon) |
| open311_SSH_CERT | Path to SSH cert |
| open311_SSH_KEY | Path to SSH public key |
| open311_LOGLOCAL | If 1, write local log file on open311_LOGPATH |
| open311_LOGPATH | Path including filename to local log (e.g. /var/log/open311) |
| sentryDSN | DSN string for Sentry telemetry |

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
