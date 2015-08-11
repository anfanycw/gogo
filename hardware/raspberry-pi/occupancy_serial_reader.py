import serial
import requests
import json

ser = serial.Serial('/dev/ttyUSB1', 9600)
while True:
    status = ser.readline().strip()
    print "Current Occupancy: " + status
    payload = {'status': status,'location': 'downstairs shower'}
    req = requests.post(
        "http://ec2-52-27-166-124.us-west-2.compute.amazonaws.com:8080/occupancy",
        data=json.dumps(payload)
    )
