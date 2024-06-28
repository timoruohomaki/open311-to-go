# Go project

## Environmental variables

Configuration parameters are stored as environment variables. In Windows they can be managed with sysdm.cpl.

| Key | Value |
| --- | --- |
| open311MongoURI | MongoDB Atlas connection string in mongodb+srv version |
| open311SentryDSN | Sentry telemetry connection |
| open311port | Port of the listener, e.g. 8080 |
| open311_SSH_CERT | Path to SSH cert |
| open311_SSH_KEY | Path to SSH public key |
| open311_LOGPATH | Path including filename to local log |


>[!TIP]
>While it is year 2024, it still might be required to reboot your Windows 10 after setting environmental variables. 
