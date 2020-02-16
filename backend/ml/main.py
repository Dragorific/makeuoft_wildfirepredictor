import time
from elasticsearch import Elasticsearch

time.sleep(30)

es = Elasticsearch(['http://elasticsearch:9200'])

es.indices.create(index="markers", ignore=400)

while(True):
    print("Hello")

