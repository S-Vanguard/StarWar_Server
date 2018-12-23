# FROM mysql:5.7.21

# #设置密码登录
# # ENV MYSQL_ROOT_PASSWORD 123456
# ENV MYSQL_ALLOW_EMPTY_PASSWORD yes


# #将所需文件放到容器中
# RUN mkdir starWar
# COPY setup.sh /starWar/setup.sh
# COPY starWar.sql /starWar/starWar.sql

# #设置容器启动时执行的命令
# CMD ["sh", "/starWar/setup.sh"]

FROM mysql:5.7.21

#设置登录密码
ENV MYSQL_ROOT_PASSWORD 123456
#定义工作目录
ENV WORK_PATH /opt/starWar
#定义会被容器自动执行的目录
ENV AUTO_RUN_DIR /docker-entrypoint-initdb.d

# 创建文件夹
RUN mkdir -p $WORK_PATH

#将所需文件放到容器中
COPY starWar.sql $WORK_PATH/
COPY setup.sh $AUTO_RUN_DIR/

#给执行文件增加可执行权限
RUN chmod a+x $AUTO_RUN_DIR/setup.sh