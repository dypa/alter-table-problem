MYSQL ALTER TABLE for huge number of rows
--------------------------------------

## mysql >= 5.6
ALTER TABLE ALGORITHM=INPLACE
see https://www.percona.com/blog/2014/11/18/avoiding-mysql-alter-table-downtime/ and https://www.percona.com/blog/wp-content/uploads/2014/11/DDLFlow1.png

## mysql < 5.6
example
```
CREATE DATABASE IF NOT EXISTS test;
USE test;
DROP TABLE IF EXISTS test;
CREATE TABLE test (
    id INT NOT NULL AUTO_INCREMENT,
	amount FLOAT(10,4) NOT NULL,
PRIMARY KEY(id)) ENGINE=InnoDB;
```

`go run main.go > data.sql`

`vagrant up && vagrant ssh`

`vagrant@wheezy:~$ mysql -u root -p dbpass < /home/project/data.sql`

```
vagrant@wheezy:~$ pt-online-schema-change u=root,p=dbpass,D=test,t=test --no-drop-old-table --alter "CHANGE amount amount decimal(15,4) NOT NULL" --execute
No slaves found.  See --recursion-method if host wheezy has slaves.
Not checking slave lag because no slaves were found and --check-slave-lag was not specified.
Operation, tries, wait:
  analyze_table, 10, 1
  copy_rows, 10, 0.25
  create_triggers, 10, 1
  drop_triggers, 10, 1
  swap_tables, 10, 1
  update_foreign_keys, 10, 1
Altering `test`.`test`...
Creating new table...
Created new table test._test_new OK.
Altering new table...
Altered `test`.`_test_new` OK.
2017-01-07T13:52:21 Creating triggers...
2017-01-07T13:52:21 Created triggers OK.
2017-01-07T13:52:21 Copying approximately 10000811 rows...
Copying rows caused a MySQL error 1265: Data truncated for column 'amount' at row 1
No more warnings about this MySQL error will be reported.  If --statistics was specified, mysql_warning_1265 will list the total count of this MySQL error.
Copying `test`.`test`:  27% 01:17 remain
Copying `test`.`test`:  55% 00:48 remain
Copying `test`.`test`:  79% 00:22 remain
2017-01-07T13:54:14 Copied rows OK.
2017-01-07T13:54:14 Swapping tables...
2017-01-07T13:54:14 Swapped original and new tables OK.
Not dropping old table because --no-drop-old-table was specified.
2017-01-07T13:54:14 Dropping triggers...
2017-01-07T13:54:14 Dropped triggers OK.
Successfully altered `test`.`test`.
vagrant@wheezy:~$
```