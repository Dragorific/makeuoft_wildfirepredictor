## @file publisher.py
#  @author Veerash Palanichamy
#  @date Feb 15, 2019

import paho.mqtt.client as mqtt
import paho.mqtt.publish as publish
import time
# read data using Pin GPIO21
# Connection parms for Solace Event Broker
solace_url = "mr2j0vvhki1l0v.messaging.solace.cloud"
solace_port = 20262
solace_user = "solace-cloud-client"
solace_passwd =  "vv59buiu782ds6bt868pp144ns"
solace_clientid = "raspberry_pi"
topic_sensor_info = "devices/temperature/events"
#     payload = "Hello from Raspberry Pi"
#     solace_url = "mr1u6o37qngitn.messaging.solace.cloud"
#     solace_port = 1883
#     solace_user = "solace-cloud-client"
#     solace_passwd = "2g71evn41v1va1jioggvch69je"
#     solace_clientid = "vats_id"
#     solace_pi_topic = "devices/+/events"
#     #solace_pi_topic = "devices/#"

# MQTT Client Connectivity to Solace Event Broker
client = mqtt.Client(solace_clientid)
client.username_pw_set(username=solace_user,password=solace_passwd)
print ("Connecting to solace {}:{} as {}". format(solace_url, solace_port, solace_user))
client.connect(solace_url, port=solace_port)
client.loop_start()

# Publish  Sensor streams to Solace Ebent Broker
num_messages=50
while True:
    #print("Temp: %d C" % result.temperature +' '+"Humid: %d %%" % result.humidity)
    # Read  Temp and humidity sensotr outputs
    temp_payload = "message"
    client.publish(topic_sensor_info, temp_payload, qos=1)
    num_messages=num_messages-1
    time.sleep(3)
client.loop_stop()
client.disconnect()