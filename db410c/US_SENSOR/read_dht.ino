#include "DHT.h"

DHT dht(A0, DHT11);
const int photoR = A1;

void setup()
{
	Serial.begin(9600);
	pinMode(photoR, INPUT);
	dht.begin();
	Serial.println("Start:usa,37.09024,-95.712891");
}

void loop()
{
	float h = dht.readHumidity();
	float t = dht.readTemperature();
	int light = analogRead(photoR);

	// check if valid, if NaN (not a number) then something went wrong!
	if (isnan(t) || isnan(h)) {
			Serial.println("Failed to read from DHT");
			return;
	}

	Serial.print("Data:");
	Serial.print("us,");
	Serial.print(h);
	Serial.print(",");
	Serial.print(t);
	Serial.print(",");
	Serial.println(light);
	delay(1000);
}

