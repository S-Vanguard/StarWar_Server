mysql -uroot -p$MYSQL_ROOT_PASSWORD <<EOF
source $WORK_PATH/starWar.sql;
service start mysql
