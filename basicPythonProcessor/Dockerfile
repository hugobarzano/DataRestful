FROM ubuntu:latest
RUN apt-get update
RUN apt-get install -y python
RUN apt-get install -y python-pip
RUN pip install --upgrade pip


ENV APP app.py
ADD . /usr/local/bin/

RUN pip install Flask==0.10.1
RUN pip install requests==2.19.1
RUN pip install urllib3==1.23


EXPOSE 5000
ENTRYPOINT python /usr/local/bin/$APP
