# Promotion import service:

Service monitors some storage, and imports promotions from the storage files. Current implemention allows to monitor only local filesystem, but service can be easily extended to work with another types of storages (ftp, sftp, ftps, s3). 

Also current implementation uses postgres database.

Service has one endpoint to get promotion by id. Request example:

> curl -X GET -H 'Accept: application/json' \
> 'localhost:3000/promotions/d018ef0b-dbd9-48f1-ac1a-eb4d90e57118'

# Build and run

To build and run project, you need to configure config.yaml file and execute commands:
- go mod download
- go build ./main.go
- ./main

## Service configuration

The configuration parameters are described here:

> server: \
>   host: //http server host \
>   port: //http server port \
>
> db: \
>   host: //db server host \
>   port: //db server port \
>   user: //db user \
>   password: //db pass \
>   db-name: //db name \
>   connection-pool: \
>     max-idle-connections: //max idle connections in connection pool \
>     max-open-connections: //max open connections in connection pool \
>     connection-lifetime: //connection lifetime milliseconds \
> local-storage: \
>   monitored-directory: //full path to local directory that should be monitored \
> scheduler: \
>   period: //The period (milliseconds) with which the directory will be monitored \
>   days-delta: //Scheduler is looking for files with modificationDate > time.Now()-days-delta \
>               //this optimizations allows to avoid loading full file-history table into memory \
>   batch-size: //Promotion entities are saved to database in batches. This parameter  \
>               //allows to declare batch size \

## Implementaion details

Service monitors files in storage. When it finds a new file, service reads it line by line, parses each line, collects lines to batches, and saves batches to db. All parsed files are included in file-history table. File-history table is used to avoid rereading of files. 



## Possible improvements:
- It will be better to split service into two different services (promotion service and monitoring service), and put message queue between them. It will allow us to make promotion service scallable. It does not make sense to scale the monitoring service, but we may need to run multiple instances of the promotion service in the future.
- To collect app metrics to prometheus, we can use github.com/prometheus/client_golang/prometheus/promhttp package