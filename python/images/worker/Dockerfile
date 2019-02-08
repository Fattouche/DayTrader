FROM python

#ENV http_proxy 'http://192.168.1.1:3128'
#ENV https_proxy 'https://192.168.1.1:3128'

RUN mkdir /service
COPY . /service
WORKDIR /service

RUN apt-get update && apt-get install -y vim

COPY ./requirements.txt .
RUN pip install --upgrade --trusted-host pypi.org --trusted-host files.pythonhosted.org --default-timeout=100 pip
RUN pip install --trusted-host pypi.org --trusted-host files.pythonhosted.org --default-timeout=100 -r /startup/requirements.txt

COPY . /service

ENTRYPOINT ["bash", "/entrypoint.sh"]
