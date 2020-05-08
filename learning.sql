-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: May 09, 2020 at 04:21 AM
-- Server version: 5.7.30-0ubuntu0.18.04.1
-- PHP Version: 7.2.24-0ubuntu0.18.04.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `learning`
--

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `desc` text,
  `user_id` int(10) UNSIGNED DEFAULT NULL,
  `tag` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `slug`, `desc`, `user_id`, `tag`) VALUES
(1, NULL, '2020-05-09 02:46:30', '2020-05-09 04:20:46', 'pertama', 'judul-pertama', 'pertama', NULL, 'pertama'),
(2, '2020-05-06 09:10:52', '2020-05-09 02:50:08', NULL, 'Judul 2', 'judul-2', 'Judul 2', NULL, 'funy'),
(3, '2020-05-09 01:52:58', '2020-05-09 01:52:58', NULL, 'Judul test', 'judul-test', 'baru body test', 1, NULL),
(4, '2020-05-09 01:56:31', '2020-05-09 03:16:32', NULL, 'Judul 2', 'dari-gugle', 'Judul 2', 3, 'funy'),
(5, '2020-05-09 02:03:04', '2020-05-09 02:03:04', NULL, ' gugle is best', 'gugle-is-best', ' test gugel', 3, 'update'),
(6, '2020-05-09 02:30:48', '2020-05-09 02:30:48', NULL, ' gugle is best', 'gugle-is-best1588966247', ' test gugel2', 3, 'update');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `social_id` varchar(255) DEFAULT NULL,
  `provider` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `role` tinyint(1) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `full_name`, `email`, `social_id`, `provider`, `avatar`, `role`) VALUES
(1, '2020-05-08 12:56:47', '2020-05-08 12:56:47', NULL, '', 'Bayu Muhammad Iqbal', 'bayuiqballl13@gmail.com', '42398906', 'github', 'https://avatars0.githubusercontent.com/u/42398906?v=4', 0),
(2, '2020-05-09 01:15:00', '2020-05-09 01:15:00', NULL, '', '', '', '', 'github', '', 0),
(3, '2020-05-09 01:54:54', '2020-05-09 01:54:54', NULL, '', 'Bayu Iqbal', 'bayuiqballl13@gmail.com', '101397635815402603496', 'google', 'https://lh3.googleusercontent.com/a-/AOh14GiHME8bTu6VUjvtpLme0MjZr743PK7ade3_8NEtUA', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uix_articles_slug` (`slug`),
  ADD KEY `idx_articles_deleted_at` (`deleted_at`),
  ADD KEY `articles_user_id_users_id_foreign` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `articles`
--
ALTER TABLE `articles`
  ADD CONSTRAINT `articles_user_id_users_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
