## @file subscriber.py
#  @author Veerash Palanichamy
#  @date Feb 15, 2019

import paho.mqtt.client as mqtt

solace_url = "mr2j0vvhki1l0v.messaging.solace.cloud"
solace_port = 20262
solace_user = "solace-cloud-client"
solace_passwd =  "vv59buiu782ds6bt868pp144ns"
solace_clientid = ""
topic_sensor_info = "sensor_data"
topic_start_info = "start_data"

def on_connect2(client, userdata, flags, rc):  # The callback for when the client connects to the broker
    print("Connected with result code {0}".format(str(rc)))  # Print result of connection attempt
    client2.subscribe(topic_start_info)

def on_message2(client, userdata, msg):  # The callback for when a PUBLISH message is received from the server.
    print("Message received2-> " + msg.topic + " " + str(msg.payload))

client2 = mqtt.Client(solace_clientid) # Create instance of client
client2.username_pw_set(solace_user, password=solace_passwd)
client2.on_connect = on_connect2 # Define callback function for successful connection
client2.on_message = on_message2 # Define callback function for receipt of a message
client2.connect(solace_url, solace_port, 60) # Connect to Solace Event Broker

# client.loop_start()
client2.loop_forever()