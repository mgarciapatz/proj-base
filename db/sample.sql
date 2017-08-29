CREATE TABLE `sample_data` (
  `autokey` bigint(20) NOT NULL,
  `sample_name` varchar(50) NOT NULL,
  `sample_password` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `sample_data`
  ADD PRIMARY KEY (`autokey`);

ALTER TABLE `sample_data`
  MODIFY `autokey` bigint(20) NOT NULL AUTO_INCREMENT;

ALTER TABLE `sample_data` CHANGE `sample_name` `sample_first_name` VARCHAR(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL;
ALTER TABLE `sample_data` CHANGE `sample_password` `sample_last_name` VARCHAR(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL;
