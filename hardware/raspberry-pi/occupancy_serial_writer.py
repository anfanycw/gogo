import serial
import requests
import json
import time

ser = serial.Serial('/dev/ttyUSB1', 9600)
while True:
    req_reserve = requests.get(
        "http://ec2-52-27-166-124.us-west-2.compute.amazonaws.com:8080/reservation/downstairs%20shower"
    )
    res = req_reserve.json()
    print res
    if (res["status"] == "RESERVED"):
        name = "%s!" % res["name"]
        ser.write(name.encode("utf-8"))
    time.sleep(5)
