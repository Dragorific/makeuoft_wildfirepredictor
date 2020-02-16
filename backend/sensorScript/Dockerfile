FROM python:3.4-alpine
ADD . /ml
WORKDIR /ml
RUN pip install elasticsearch
RUN pip install paho-mqtt
RUN pip install sklearn
CMD ["python", "sensorMain.py"]
