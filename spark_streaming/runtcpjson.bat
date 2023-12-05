docker cp tcp-json.scala spark-master:/tcp-json.scala

docker exec -it spark-master spark/bin/spark-shell --packages org.apache.spark:spark-sql-kafka-0-10_2.12:3.0.0 -i tcp-json.scala