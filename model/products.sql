-- MySQL dump 10.16  Distrib 10.1.28-MariaDB, for Win32 (AMD64)
--
-- Host: localhost    Database: products
-- ------------------------------------------------------
-- Server version	10.1.28-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product` (
  `category` varchar(20) DEFAULT NULL,
  `products_id` varchar(20) DEFAULT NULL,
  `products_name` varchar(50) DEFAULT NULL,
  `stocks` int(11) DEFAULT NULL,
  `price` varchar(20) DEFAULT NULL,
  `promotion` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES ('Books','111','Do Androids Live?',51,'70000','Discount 10%'),('Books','112','Pride and Prejudice',40,'85000','free shipping'),('Books','113','How To Be Londoner',59,'100000','Discount 10%'),('Books','114','Silence Breaking',30,'150000','Discount 50%'),('Books','115','Black Diaries',74,'65000','free shipping'),('Books','116','The Man Without Qualities',85,'200000','Discount 50%'),('Books','117','To Kill a Mockingbird',99,'350000','Discount 10%'),('Fashion','118','Zara Herringbone Culottes',1,'750000','free shipping'),('Fashion','119','Moschino Baby Red Shoulder Bag',79,'350000','free shipping'),('Fashion','120','Rose Garden Gown',30,'800000','Discount 10%'),('Fashion','121','Queercore brogue over the knee boot',56,'1000000','free shipping'),('Stationary','122','Wreck This Journal Book ',73,'100000','Discount 50%'),('Stationary','123','Triple Color Pen',1,'75000','free shipping'),('Stationary','124','Unforgettable Pen',71,'85000','Discount 50%'),('Beauty Products','125','King Kylie Lipstick ',54,'200000',''),('Beauty Products','126','Huda Beauty Palette',52,'300000','free shipping'),('Beauty Products','127','Catokan Babyliss ipro',67,'1500000','Discount 50%'),('Beauty Products','128','Maybeline Foundation',76,'200000','free shipping'),('Electronics','129','Arduino Uno R3',33,'85000','Discount 10%'),('Home appliances','130','Toaster T17',0,'3500000','free shipping'),('Home appliances','131','Coffee & Tea Maker R35',95,'3750000','Discount 10%');
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-11-11 22:12:24
