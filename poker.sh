#!/usr/bin/expect -f
set host [lindex $argv 0];
set user [lindex $argv 1];
set pass [lindex $argv 2];
set port [lindex $argv 3];
spawn ssh ${user}@${host} -p ${port}
expect "?* "
send "yes\r"
expect "password:*"
send "${pass}\r"
expect "$ "
send "help\r"
expect "$ "
send "exit\r"