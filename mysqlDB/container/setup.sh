# mysql -uroot -p123456 < /opt/starWar/starWar.sql
mysql -uroot -p$MYSQL_ROOT_PASSWORD <<EOF
source $WORK_PATH/starWar.sql;
service start mysql
