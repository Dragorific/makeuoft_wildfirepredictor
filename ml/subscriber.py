## @file subscriber.py
#  @author Veerash Palanichamy
#  @date Feb 15, 2019
from sklearn.model_selection import RandomizedSearchCV, GridSearchCV
from sklearn.metrics import confusion_matrix, f1_score
from sklearn.externals import joblib
import paho.mqtt.client as mqtt
import numpy as np

solace_url = "mr2j0vvhki1l0v.messaging.solace.cloud"
solace_port = 20262
solace_user = "solace-cloud-client"
solace_passwd =  "vv59buiu782ds6bt868pp144ns"
solace_clientid = ""
topic_sensor_info = "sensor_data"

def predict(temp, hum):
    input = [5,5,6,12,90,70,650]
    input.append(float(temp))
    input.append(float(hum))
    input.append(50)
    input.append(0)
    # input=np.array(input)
    y_pred = clf.predict(input)
    return y_pred[0]

def on_connect(client, userdata, flags, rc):  # The callback for when the client connects to the broker
    print("Connected with result code {0}".format(str(rc)))  # Print result of connection attempt
    client.subscribe(topic_sensor_info)  # Subscribe to the topic “digitest/test1”, receive any messages published on it

def on_message(client, userdata, msg):  # The callback for when a PUBLISH message is received from the server.
    sensor = (str(msg.payload).split("\\r\\n")[0].split("b'")[1].split(","))    # Print a received msg
    print(predict(sensor[-1], sensor[-2]))

clf = joblib.load('model.joblib')
client = mqtt.Client(solace_clientid) # Create instance of client
client.username_pw_set(solace_user, password=solace_passwd)
client.on_connect = on_connect # Define callback function for successful connection
client.on_message = on_message # Define callback function for receipt of a message
client.connect(solace_url, solace_port, 60) # Connect to Solace Event Broker
# client.loop_start()
client.loop_forever()  # Start networking daemon