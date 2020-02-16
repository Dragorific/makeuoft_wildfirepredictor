## @file publisher.py
#  @author Veerash Palanichamy
#  @date Feb 15, 2019

import paho.mqtt.client as mqtt
import paho.mqtt.publish as publish

import serial
import time

solace_url = "mr2j0vvhki1l0v.messaging.solace.cloud"
solace_port = 20262
solace_user = "solace-cloud-client"
solace_passwd =  "vv59buiu782ds6bt868pp144ns"
solace_clientid = "raspberry_pi"
topic_sensor_info = "sensor_data"
topic_start_info = "start_data"

ard = serial.Serial('/dev/tty96B0', 9600)

if __name__ == '__main__':
    	print("Welcome to the Humidity & Temperature reader!!!")
    	client = mqtt.Client(solace_clientid)
    	client.username_pw_set(username=solace_user,password=solace_passwd)
    	print ("Connecting to solace {}:{} as {}". format(solace_url, solace_port, solace_user))
    	client.connect(solace_url, port=solace_port)
    	client.loop_start()
	while True:
		ardOut = ard.readline()
		if ardOut.find("Start:") != -1:
			ardData = ardOut.split("Start:")[1]
			ardName = ardData.split(',')[0]
			ardLat = ardData.split(',')[1]
			ardLong = ardData.split(',')[2]
			client.publish(topic_start_info, ardData, qos=1)
			print(ardData)
		if ardOut.find("Data:") != -1:
			ardData = ardOut.split("Data:")[1]
			ardName = ardData.split(',')[0]
			ardHumid = ardData.split(',')[1]
			ardTemp = ardData.split(',')[2]
			ardLight = ardData.split(',')[3]
        		print(ardOut)
        		client.publish(topic_sensor_info, ardData, qos=1)
	client.loop_stop()
	client.disconnect()
