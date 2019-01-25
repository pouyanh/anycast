CREATE USER IF NOT EXISTS 'race' IDENTIFIED WITH mysql_native_password BY 'phi0lambda';
GRANT ALL PRIVILEGES ON `race_%`.* TO 'race'@'%' WITH GRANT OPTION;
