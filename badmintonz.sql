-- MySQL dump 10.13  Distrib 8.0.44, for Linux (x86_64)
--
-- Host: localhost    Database: badmintonzz
-- ------------------------------------------------------
-- Server version	8.0.44

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
-- Table structure for table `bookings`
--

DROP TABLE IF EXISTS `bookings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bookings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `start_time` time NOT NULL,
  `end_time` time NOT NULL,
  `type_payment` varchar(50) NOT NULL,
  `status_booking` enum('pending','confirmed','cancelled') NOT NULL,
  `description` text,
  `user_id` int NOT NULL,
  `payments_id` int NOT NULL,
  `service_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `payments_id` (`payments_id`),
  KEY `service_id` (`service_id`),
  CONSTRAINT `bookings_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `bookings_ibfk_2` FOREIGN KEY (`payments_id`) REFERENCES `payments` (`id`),
  CONSTRAINT `bookings_ibfk_3` FOREIGN KEY (`service_id`) REFERENCES `services` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bookings`
--

LOCK TABLES `bookings` WRITE;
/*!40000 ALTER TABLE `bookings` DISABLE KEYS */;
INSERT INTO `bookings` VALUES (1,'2024-01-15','09:00:00','09:30:00','cash','confirmed','Potong rambut biasa',2,1,1,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(2,'2024-01-16','10:00:00','11:00:00','transfer','confirmed','Potong + blow dry',3,2,2,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(3,'2024-01-17','13:00:00','14:30:00','midtrans','confirmed','Creambath paket lengkap',4,3,3,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(4,'2024-01-18','14:00:00','16:30:00','midtrans','pending','Smoothing rambut panjang',5,4,4,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(5,'2024-01-19','09:30:00','11:30:00','midtrans','cancelled','Coloring - dibatalkan pelanggan',6,5,5,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(6,'2024-01-20','11:00:00','12:30:00','qris','confirmed','Perawatan wajah rutin',7,6,6,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(7,'2024-01-21','15:00:00','15:45:00','cash','confirmed','Manicure hari Sabtu',8,7,7,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(8,'2024-01-22','10:00:00','11:30:00','transfer','confirmed','Paket mani-pedi lengkap',2,8,9,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(9,'2024-01-23','13:00:00','14:30:00','midtrans','cancelled','Hair spa - refund diproses',3,9,10,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(10,'2024-01-24','09:00:00','09:30:00','qris','confirmed','Potong cepat',4,10,8,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(11,'2024-01-25','16:00:00','16:30:00','cash','confirmed','Potong rambut sore',5,11,1,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(12,'2024-01-26','11:00:00','12:00:00','midtrans','pending','Booking via app, belum bayar',6,12,2,'2026-04-12 05:29:47','2026-04-12 05:29:47');
/*!40000 ALTER TABLE `bookings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bussiness_hour`
--

DROP TABLE IF EXISTS `bussiness_hour`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bussiness_hour` (
  `id` int NOT NULL AUTO_INCREMENT,
  `day` varchar(50) NOT NULL,
  `start_time` time DEFAULT NULL,
  `end_time` time DEFAULT NULL,
  `is_open` tinyint(1) NOT NULL DEFAULT '0',
  `Description` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bussiness_hour`
--

LOCK TABLES `bussiness_hour` WRITE;
/*!40000 ALTER TABLE `bussiness_hour` DISABLE KEYS */;
INSERT INTO `bussiness_hour` VALUES (1,'Senin','08:00:00','20:00:00',1,'Jam operasional hari Senin','2026-04-12 05:29:47','2026-04-12 05:29:47'),(2,'Selasa','08:00:00','20:00:00',1,'Jam operasional hari Selasa','2026-04-12 05:29:47','2026-04-12 05:29:47'),(3,'Rabu','08:00:00','20:00:00',1,'Jam operasional hari Rabu','2026-04-12 05:29:47','2026-04-12 05:29:47'),(4,'Kamis','08:00:00','20:00:00',1,'Jam operasional hari Kamis','2026-04-12 05:29:47','2026-04-12 05:29:47'),(5,'Jumat','08:00:00','21:00:00',1,'Jam operasional hari Jumat - lebih lama','2026-04-12 05:29:47','2026-04-12 05:29:47'),(6,'Sabtu','07:00:00','21:00:00',1,'Jam operasional hari Sabtu','2026-04-12 05:29:47','2026-04-12 05:29:47'),(7,'Minggu',NULL,NULL,0,'Tutup pada hari Minggu','2026-04-12 05:29:47','2026-04-12 05:29:47');
/*!40000 ALTER TABLE `bussiness_hour` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payments`
--

DROP TABLE IF EXISTS `payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `amount` decimal(10,2) NOT NULL,
  `payment_method` varchar(255) NOT NULL,
  `payment_status` enum('pending','completed','failed','refunded','expired') NOT NULL,
  `payment_url` text,
  `transaction_id` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
INSERT INTO `payments` VALUES (1,35000.00,'cash','completed',NULL,NULL,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(2,60000.00,'transfer','completed',NULL,'TXN-20240101-001','2026-04-12 05:29:47','2026-04-12 05:29:47'),(3,85000.00,'midtrans','completed','https://app.sandbox.midtrans.com/snap/v1/1','TXN-20240102-002','2026-04-12 05:29:47','2026-04-12 05:29:47'),(4,250000.00,'midtrans','pending','https://app.sandbox.midtrans.com/snap/v1/2','TXN-20240103-003','2026-04-12 05:29:47','2026-04-12 05:29:47'),(5,200000.00,'midtrans','failed','https://app.sandbox.midtrans.com/snap/v1/3','TXN-20240104-004','2026-04-12 05:29:47','2026-04-12 05:29:47'),(6,120000.00,'qris','completed',NULL,'TXN-20240105-005','2026-04-12 05:29:47','2026-04-12 05:29:47'),(7,75000.00,'cash','completed',NULL,NULL,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(8,140000.00,'transfer','completed',NULL,'TXN-20240106-006','2026-04-12 05:29:47','2026-04-12 05:29:47'),(9,150000.00,'midtrans','refunded','https://app.sandbox.midtrans.com/snap/v1/4','TXN-20240107-007','2026-04-12 05:29:47','2026-04-12 05:29:47'),(10,75000.00,'qris','completed',NULL,'TXN-20240108-008','2026-04-12 05:29:47','2026-04-12 05:29:47'),(11,35000.00,'cash','completed',NULL,NULL,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(12,60000.00,'midtrans','expired','https://app.sandbox.midtrans.com/snap/v1/5','TXN-20240109-009','2026-04-12 05:29:47','2026-04-12 05:29:47');
/*!40000 ALTER TABLE `payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `services` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `services`
--

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;
INSERT INTO `services` VALUES (1,'Potong Rambut Pria',35000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(2,'Potong Rambut Wanita',60000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(3,'Creambath',85000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(4,'Smoothing',250000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(5,'Coloring Rambut',200000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(6,'Perawatan Wajah Dasar',120000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(7,'Manicure',75000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(8,'Pedicure',75000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(9,'Manicure & Pedicure',140000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47'),(10,'Hair Spa',150000.00,'2026-04-12 05:29:47','2026-04-12 05:29:47');
/*!40000 ALTER TABLE `services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `no_hp` varchar(20) NOT NULL,
  `role` enum('user','admin') NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `no_hp` (`no_hp`),
  UNIQUE KEY `uk_users_email_phone` (`email`,`no_hp`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin_budi','budi@example.com','$2a$10$hashedpassword1','081234567890','admin','2026-04-12 05:29:47','2026-04-12 05:29:47'),(2,'siti_rahayu','siti@example.com','$2a$10$hashedpassword2','081234567891','user','2026-04-12 05:29:47','2026-04-12 05:29:47'),(3,'andi_wijaya','andi@example.com','$2a$10$hashedpassword3','081234567892','user','2026-04-12 05:29:47','2026-04-12 05:29:47'),(4,'dewi_kartika','dewi@example.com','$2a$10$hashedpassword4','081234567893','user','2026-04-12 05:29:47','2026-04-12 05:29:47'),(5,'reza_pratama','reza@example.com','$2a$10$hashedpassword5','081234567894','user','2026-04-12 05:29:47','2026-04-12 05:29:47'),(6,'linda_susanti','linda@example.com','$2a$10$hashedpassword6','081234567895','user','2026-04-12 05:29:47','2026-04-12 05:29:47'),(7,'fajar_nugroho','fajar@example.com','$2a$10$hashedpassword7','081234567896','user','2026-04-12 05:29:47','2026-04-12 05:29:47'),(8,'maya_putri','maya@example.com','$2a$10$hashedpassword8','081234567897','user','2026-04-12 05:29:47','2026-04-12 05:29:47');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-04-17  8:19:42
