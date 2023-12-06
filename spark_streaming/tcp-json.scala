import spark.implicits._
import org.apache.spark.sql.functions._
import org.apache.spark.sql.types._
import org.apache.spark.sql.streaming._
val schema = new StructType()
    .add("name", StringType)
    .add("float", FloatType)
    .add("qualityValue", StringType)
    .add("marketplaceName", StringType)
    .add("marketplaceUrl", StringType)
    .add("USDPrice", IntegerType)
    
val toStringM = udf((payload: Array[Byte]) => new String(payload))

val lines = spark
  .readStream
  .format("kafka")
  .option("kafka.bootstrap.servers", "kafka-test:9092")
  .option("subscribe", "skins")
  .load()

val parsedDF = lines
  .select(from_json(toStringM(col("value")), schema))


val skins = parsedDF
  .select("from_json(UDF(value)).name", "from_json(UDF(value)).float", "from_json(UDF(value)).qualityValue", "from_json(UDF(value)).marketplaceName", "from_json(UDF(value)).marketplaceUrl", "from_json(UDF(value)).USDPrice")

val skinsTimestamp = skins.withColumn("timestamp", current_timestamp())


val windowedCounts = skinsTimestamp
  .withWatermark("timestamp", "1 minute")  
  .groupBy(window($"timestamp", "1 minute"),  $"name")
  .agg(avg("USDPrice").as("Average_Price"))
  .select("window.start", "window.end", "name", "Average_Price")


val query = windowedCounts
  .writeStream
  .outputMode("complete")  
  .format("console") 
  .start()


