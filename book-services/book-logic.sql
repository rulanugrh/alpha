-- MySQL dump 10.19  Distrib 10.3.34-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: book
-- ------------------------------------------------------
-- Server version       10.3.34-MariaDB-0+deb10u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `books` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `stock` bigint(20) DEFAULT NULL,
  `author` longtext DEFAULT NULL,
  `examplar` bigint(20) DEFAULT NULL,
  `seller_id` bigint(20) unsigned DEFAULT NULL,
  `category_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_books_deleted_at` (`deleted_at`),
  KEY `fk_sellers_books` (`seller_id`),
  KEY `fk_categories_books` (`category_id`),
  CONSTRAINT `fk_categories_books` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
  CONSTRAINT `fk_sellers_books` FOREIGN KEY (`seller_id`) REFERENCES `sellers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (1,'2023-04-06 15:45:43.256','2023-04-06 15:45:43.256',NULL,'Mystic of Celtic',50000,10,'Pradana',5,4,3),(2,'2023-04-06 16:19:24.033','2023-04-06 16:19:24.033',NULL,'Mystic of Yunani',50000,10,'Sagara',5,5,4),(3,'2023-04-06 16:23:26.869','2023-04-06 16:23:26.869',NULL,'Mystic of Nordik',50000,10,'Anugerah',5,6,5),(4,'2023-04-06 16:29:59.752','2023-04-06 16:29:59.752',NULL,'Mystic of Mesir',50000,10,'Ryan',5,7,6);
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'2023-04-05 16:34:40.830','2023-04-05 16:34:40.830',NULL,'romance',''),(2,'2023-04-06 15:43:57.896','2023-04-06 15:43:57.896',NULL,'Mysteri',''),(3,'2023-04-06 15:45:43.253','2023-04-06 15:45:43.253',NULL,'Mystic of Celtic',''),(4,'2023-04-06 16:19:24.031','2023-04-06 16:19:24.031',NULL,'Mystic of Yunani',''),(5,'2023-04-06 16:23:26.849','2023-04-06 16:23:26.849',NULL,'Mystic of Nordik',''),(6,'2023-04-06 16:29:59.750','2023-04-06 16:29:59.750',NULL,'Mystic of Mesir','');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sellers`
--

DROP TABLE IF EXISTS `sellers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sellers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sellers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sellers`
--

LOCK TABLES `sellers` WRITE;
/*!40000 ALTER TABLE `sellers` DISABLE KEYS */;
INSERT INTO `sellers` VALUES (1,'2023-04-05 16:33:29.315','2023-04-05 16:33:29.315',NULL,'arul'),(3,'2023-04-06 15:43:27.878','2023-04-06 15:43:27.878',NULL,'Arul'),(4,'2023-04-06 15:45:43.245','2023-04-06 15:45:43.245',NULL,'Mystic of Celtic'),(5,'2023-04-06 16:19:24.024','2023-04-06 16:19:24.024',NULL,'Mystic of Yunani'),(6,'2023-04-06 16:23:26.847','2023-04-06 16:23:26.847',NULL,'Mystic of Nordik'),(7,'2023-04-06 16:29:59.749','2023-04-06 16:29:59.749',NULL,'Mystic of Mesir');
/*!40000 ALTER TABLE `sellers` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-06 12:40:43