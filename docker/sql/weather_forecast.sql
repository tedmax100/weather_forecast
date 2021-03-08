-- mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 8.0.23, for Linux (x86_64)
--
-- Host: localhost    Database: weather_forecast
-- ------------------------------------------------------
-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `weather_forecast`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `weather_forecast` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `weather_forecast`;

--
-- Table structure for table `element`
--

DROP TABLE IF EXISTS `element`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `element` (
  `id` int unsigned NOT NULL,
  `name` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `element`
--

LOCK TABLES `element` WRITE;
/*!40000 ALTER TABLE `element` DISABLE KEYS */;
INSERT INTO `element` VALUES (1,'Wx','天氣現象'),(2,'PoP','降雨機率'),(3,'CI','舒適度'),(4,'MinT','最低溫度'),(5,'MaxT','最高溫度');
/*!40000 ALTER TABLE `element` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location`
--

DROP TABLE IF EXISTS `location`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `location` (
  `id` int unsigned NOT NULL,
  `name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location`
--

LOCK TABLES `location` WRITE;
/*!40000 ALTER TABLE `location` DISABLE KEYS */;
INSERT INTO `location` VALUES (1,'臺北市'),(2,'新北市'),(3,'桃園市');
/*!40000 ALTER TABLE `location` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `record`
--

DROP TABLE IF EXISTS `record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `record` (
  `location_id` int NOT NULL,
  `element_id` tinyint unsigned NOT NULL,
  `parameter_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `parameter_unit` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `parameter_value` int DEFAULT NULL,
  `start_time` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `end_time` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`start_time`,`location_id`,`element_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `record`
--

LOCK TABLES `record` WRITE;
/*!40000 ALTER TABLE `record` DISABLE KEYS */;
INSERT INTO `record` VALUES (1,1,'陰短暫陣雨',NULL,11,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(1,2,'50','百分比',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(1,3,'稍有寒意',NULL,NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(1,4,'17','C',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(1,5,'18','C',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(2,1,'陰短暫陣雨',NULL,11,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(2,2,'40','百分比',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(2,3,'稍有寒意',NULL,NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(2,4,'17','C',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(2,5,'18','C',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(3,1,'陰短暫陣雨',NULL,11,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(3,2,'60','百分比',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(3,3,'稍有寒意',NULL,NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(3,4,'17','C',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(3,5,'17','C',NULL,'2021-03-06 12:00:00','2021-03-06 18:00:00'),(1,1,'陰陣雨',NULL,14,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(1,2,'70','百分比',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(1,3,'稍有寒意',NULL,NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(1,4,'16','C',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(1,5,'17','C',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(2,1,'陰陣雨',NULL,14,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(2,2,'70','百分比',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(2,3,'稍有寒意',NULL,NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(2,4,'16','C',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(2,5,'17','C',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(3,1,'陰陣雨',NULL,14,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(3,2,'80','百分比',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(3,3,'稍有寒意',NULL,NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(3,4,'16','C',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(3,5,'17','C',NULL,'2021-03-06 18:00:00','2021-03-07 06:00:00'),(1,1,'陰短暫陣雨',NULL,11,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(1,2,'50','百分比',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(1,3,'稍有寒意',NULL,NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(1,4,'16','C',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(1,5,'19','C',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(2,1,'陰短暫陣雨',NULL,11,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(2,2,'50','百分比',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(2,3,'稍有寒意',NULL,NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(2,4,'16','C',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(2,5,'19','C',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(3,1,'陰短暫陣雨',NULL,11,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(3,2,'40','百分比',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(3,3,'稍有寒意',NULL,NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(3,4,'16','C',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(3,5,'18','C',NULL,'2021-03-07 06:00:00','2021-03-07 18:00:00'),(1,1,'陰短暫雨',NULL,11,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(1,2,'40','百分比',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(1,3,'稍有寒意',NULL,NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(1,4,'16','C',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(1,5,'18','C',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(2,1,'陰短暫雨',NULL,11,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(2,2,'40','百分比',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(2,3,'稍有寒意',NULL,NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(2,4,'17','C',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(2,5,'18','C',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(3,1,'陰天',NULL,7,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(3,2,'20','百分比',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(3,3,'寒冷至稍有寒意',NULL,NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(3,4,'15','C',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(3,5,'17','C',NULL,'2021-03-07 18:00:00','2021-03-08 06:00:00'),(1,1,'陰天',NULL,7,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(1,2,'10','百分比',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(1,3,'稍有寒意',NULL,NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(1,4,'16','C',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(1,5,'18','C',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(2,1,'陰天',NULL,7,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(2,2,'10','百分比',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(2,3,'稍有寒意',NULL,NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(2,4,'17','C',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(2,5,'19','C',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(3,1,'陰時多雲',NULL,6,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(3,2,'10','百分比',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(3,3,'寒冷至稍有寒意',NULL,NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(3,4,'15','C',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(3,5,'18','C',NULL,'2021-03-08 06:00:00','2021-03-08 18:00:00'),(1,1,'多雲',NULL,4,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(1,2,'10','百分比',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(1,3,'稍有寒意',NULL,NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(1,4,'16','C',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(1,5,'17','C',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(2,1,'多雲時晴',NULL,3,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(2,2,'10','百分比',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(2,3,'稍有寒意',NULL,NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(2,4,'16','C',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(2,5,'18','C',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(3,1,'多雲時晴',NULL,3,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(3,2,'10','百分比',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(3,3,'寒冷至稍有寒意',NULL,NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(3,4,'15','C',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00'),(3,5,'17','C',NULL,'2021-03-08 18:00:00','2021-03-09 06:00:00');
/*!40000 ALTER TABLE `record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `password` varchar(100) NOT NULL,
  `token` varchar(80) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'testaccount','$2a$10$x8EXY8o2VIr4oATTslypLOVVezP6w.cXvHG74WY8NsTOIphUogVaW\n','abcd1234','2021-03-05 06:35:22','2021-03-06 16:35:07');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-07 15:41:17
