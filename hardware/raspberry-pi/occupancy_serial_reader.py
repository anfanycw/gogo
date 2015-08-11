import serial
import requests
import json

ser = serial.Serial('/dev/ttyUSB0', 9600)
while True:
    status = ser.readline().strip()
    print "Current Occupancy: " + status
    payload = {'status': status,'location': 'downstairs shower'}
    req = requests.post(
          "http://requestb.in/pwdbwqpw",
          data=json.dumps(payload)
    )
