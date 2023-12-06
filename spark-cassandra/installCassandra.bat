docker cp spark-cassandra-connector spark-master:/spark-cassandra-connector
docker cp installCassandra.sh spark-master:/installCassandra.sh
docker exec -it spark-master bin/bash installCassandra.sh