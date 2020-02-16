import paho.mqtt.client as mqtt
import time
import json
from elasticsearch import Elasticsearch

time.sleep(30)

# This variable will house all the arrays of markers and continuously update the ES database

#es = None

# This defines all the solace data we reference
solace_url = "mr2j0vvhki1l0v.messaging.solace.cloud"
solace_port = 20262
solace_user = "solace-cloud-client"
solace_passwd = "vv59buiu782ds6bt868pp144ns"
solace_clientid = ""                                # Leave clientID empty for default key, errors otherwise
topic_sensor_info = "sensor_data"                    # Markers

qos = 1                                             #qos will always be a constant of 1



def on_connect(client, userdata, flags, rc):  # The callback for when the client connects to the broker
    print("Connected with result code {0}".format(str(rc)))  # Print result of connection attempt
    client.subscribe(topic_sensor_info)  # Subscribe to the topic “digitest/test1”, receive any messages published on it


def on_message(client, userdata, msg):  # The callback for when a PUBLISH message is received from the server.
    sensor = (str(msg.payload).split("\\r\\n")[0].split("b'")[1].split(","))      #This will break up the string ("CAN,lat,long") into [name, latitude, longitude], and add it into marker
    es.index(index=sensor[0], doc_type='_doc', id=sensor[0], body={"doc": {"data": {"sensor":sensor[1:]}, "timestamp": }})
    print("Message received-> " + msg.topic + " " + str(msg.payload) + " " + str(sensor[0]))  # Print a received msg to cehck




#if __name__ == "__main__":
    ## Subscribe to the Solace Broker and receive sensor information
es = Elasticsearch(['http://elasticsearch:9200'])
print("Index has been created, started listening")

# Create client for markers
client = mqtt.Client(solace_clientid)  # Create instance of marker client


print("Main has begun")
client.username_pw_set(solace_user, password=solace_passwd)  # set username and password

# Connect to both clients
client.on_connect = on_connect # Define callback function for successful connection
client.on_message = on_message # Define callback function for receipt of a message
client.connect(solace_url, solace_port, 60) # Connect to Solace Event Broker
# client.loop_start()
client.loop_forever()  # Start networking daemon