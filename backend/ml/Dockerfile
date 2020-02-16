FROM python:3.4-alpine
ADD . /ml
WORKDIR /ml
RUN pip install elasticsearch
CMD ["python", "main.py"]
