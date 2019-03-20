-- phpMyAdmin SQL Dump
-- version 4.8.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Mar 19, 2019 at 04:30 PM
-- Server version: 5.7.24
-- PHP Version: 7.3.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tzx_tmp`
--

-- --------------------------------------------------------

--
-- Table structure for table `poi`
--

DROP TABLE IF EXISTS `poi`;
CREATE TABLE IF NOT EXISTS `poi` (
  `Id` int(11) NOT NULL,
  `Name` varchar(56) COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `Type` int(11) NOT NULL COMMENT '类型',
  `Video` varchar(254) COLLATE utf8mb4_bin NOT NULL,
  `Imgs` varchar(254) COLLATE utf8mb4_bin NOT NULL,
  `Introduction` varchar(54) COLLATE utf8mb4_bin NOT NULL,
  `Lon` varchar(24) COLLATE utf8mb4_bin NOT NULL,
  `Lat` varchar(24) COLLATE utf8mb4_bin NOT NULL,
  `Label` varchar(24) COLLATE utf8mb4_bin NOT NULL,
  `Price` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `ShopHours` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `PlayTime` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `Region` int(11) NOT NULL,
  `Image` varchar(254) COLLATE utf8mb4_bin NOT NULL,
  `Dubbing` varchar(254) COLLATE utf8mb4_bin NOT NULL,
  `Time` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `poi`
--

INSERT INTO `poi` (`Id`, `Name`, `Type`, `Video`, `Imgs`, `Introduction`, `Lon`, `Lat`, `Label`, `Price`, `ShopHours`, `PlayTime`, `Region`, `Image`, `Dubbing`, `Time`) VALUES
(1, '测试', 1, '测试', '测试', '测试', '测试', '测试', '测试', '测试', '测试', '测试', 1, '测试', '测试', 123123123);

-- --------------------------------------------------------

--
-- Table structure for table `route`
--

DROP TABLE IF EXISTS `route`;
CREATE TABLE IF NOT EXISTS `route` (
  `Id` int(56) NOT NULL,
  `City` varchar(56) COLLATE utf8mb4_bin NOT NULL,
  `Name` char(56) COLLATE utf8mb4_bin NOT NULL,
  `Video` char(254) COLLATE utf8mb4_bin NOT NULL,
  `Imgs` char(254) COLLATE utf8mb4_bin NOT NULL,
  `Price` int(11) NOT NULL,
  `GuideId` int(11) NOT NULL,
  `ChangeTime` int(11) NOT NULL,
  `Time` int(11) NOT NULL,
  `ImgArr` varchar(354) COLLATE utf8mb4_bin NOT NULL,
  `Label` varchar(354) COLLATE utf8mb4_bin NOT NULL,
  `ImgArr2` varchar(354) COLLATE utf8mb4_bin NOT NULL,
  `ImgArr3` varchar(354) COLLATE utf8mb4_bin NOT NULL,
  `ImgArr4` varchar(354) COLLATE utf8mb4_bin NOT NULL,
  `ViewCount` int(11) NOT NULL,
  `SaleCount` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `route`
--

INSERT INTO `route` (`Id`, `City`, `Name`, `Video`, `Imgs`, `Price`, `GuideId`, `ChangeTime`, `Time`, `ImgArr`, `Label`, `ImgArr2`, `ImgArr3`, `ImgArr4`, `ViewCount`, `SaleCount`) VALUES
(1, '上海', '上海', '上海', '上海', 1, 1, 1, 1, '1', '1', '1', '1', '1', 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `routepoi`
--

DROP TABLE IF EXISTS `routepoi`;
CREATE TABLE IF NOT EXISTS `routepoi` (
  `Id` int(11) NOT NULL,
  `RouteId` int(11) NOT NULL COMMENT '路线ID',
  `PoiId` int(11) NOT NULL COMMENT 'POI ID',
  `Sort` int(11) NOT NULL COMMENT '排序',
  `Time` int(11) NOT NULL COMMENT '时间',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `routepoi`
--

INSERT INTO `routepoi` (`Id`, `RouteId`, `PoiId`, `Sort`, `Time`) VALUES
(1, 1, 1, 1, 123123132);

-- --------------------------------------------------------

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
CREATE TABLE IF NOT EXISTS `task` (
  `Id` int(11) NOT NULL,
  `PoiId` int(11) NOT NULL COMMENT '景点表ID',
  `Type` int(1) NOT NULL,
  `Interaction` int(1) NOT NULL,
  `TaskCopy` text COLLATE utf8mb4_bin NOT NULL,
  `ButtonCopy` varchar(10) COLLATE utf8mb4_bin NOT NULL,
  `ImgArr` text COLLATE utf8mb4_bin NOT NULL,
  `Sort` int(11) NOT NULL,
  `Time` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `task`
--

INSERT INTO `task` (`Id`, `PoiId`, `Type`, `Interaction`, `TaskCopy`, `ButtonCopy`, `ImgArr`, `Sort`, `Time`) VALUES
(1, 1, 1, 1, '1', '1', '1', 1, 1);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
