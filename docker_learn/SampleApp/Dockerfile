FROM ubuntu:16.04
MAINTAINER Utkarsh Mani Tripathi <utkarshmani1997@gmail.com>
ADD . /home        
WORKDIR /home
VOLUME /var/server
COPY Server /home
ENTRYPOINT ["./entrypoint.sh"]
EXPOSE 8080
