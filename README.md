# MakeUofT WildfirePredictor
MakeUofT project, uses temperature, light and humidity information and processes into a machine learning model to predict wildfires before they occur.

This project was developed on a Vue.js-based frontend, and a Golang and Python-based backend. It was deployed using the Kubernetes software "Docker" and incorporated the Solace PubSub+ WebBroker to communicate between a Qualcomm Dragonboard 410c and the WebBroker over WiFi. The Dragonboard acted as an all-in-one meteorological unit that can observe the environment around it using the Grove temperature and humidifier sensor, as well as a photoresistor. The Dragonboard was also used to communicate its GPS position for planned scalability involving autonomous deployed drones. 

# Development Mode
To deploy the containers use; `docker-compose up -d --build`
To remove/prune the containers use; 
`docker-compose stop`
`docker image prune`
Front-end can be accessed from: `web.127.0.0.1.xip.io/`
Kibana can be accessed from: `kibana.127.0.0.1.xip.io/`
