import os

from flask import Flask
from flask import request
from flask import jsonify

import json
import urllib2
import requests

from test import *
import sys


from operatorBasic import Calc
import time

app = Flask(__name__)


def RegisterServiceToDataresful(my_url,datarestful_url):
    print "\n\n[operatorBasic SERVICE] Added as a processor to Dataresful"
    data= {
        "title": "basicPythonProcessor",
        "url": '"'+my_url+'"'
    }

    req = urllib2.Request(datarestful_url+'/Services/')
    req.add_header('Content-Type', 'application/json')
    response = urllib2.urlopen(req, json.dumps(data))
    print response

def RemoveServiceFromDataresful(datarestful_url):
    print "\n\n[operatorBasic SERVICE] Deleting as a processor from Dataresful"
    r = requests.delete(datarestful_url+'/RemoveService/basicPythonProcessor')
    print r.status_code



@app.route('/basicOperator', methods=['POST'])
def basicOperator():

    if request.json:

        mybody = request.json
        title = mybody.get("title")
        data = mybody.get("data")
        value = mybody.get("value")
        operator = mybody.get("operator")

        print title
        print data
        print value
        print operator
        data_response=Calc(data,operator,value)

        res = {}
        string_response=[]
        for r in data_response:
            string_response.append(str(r))

        res["title"] = title + "_" + operator + "_" + time.strftime("%Y%m%d_%H%M%S")
        res["data"] = string_response
        return  jsonify(res)

    else:
        return "No Dataset received"





if __name__ == '__main__':

    if TestBasicSum() is False:
        sys.exit('[STOP SERVICE] Continuous Testing fail: SUM')
    if TestBasicSub() is False:
        sys.exit('[STOP SERVICE] Continuous Testing fail: SUB')
    if TestBasicMul() is False:
        sys.exit('[STOP SERVICE] Continuous Testing fail: MUL')
    if TestBasicDiv() is False:
        sys.exit('[STOP SERVICE] Continuous Testing fail: DIV')

    print "\n\n[operatorBasic SERVICE] Ready to Roll"

    SERVICE_URL = os.getenv('SERVICE_URL',"http://basic:5000/basicOperator" )
    DATARESTFUL_URL=os.getenv('DATARESTFUL_URL',"http://localhost:8080")

    RegisterServiceToDataresful(SERVICE_URL,DATARESTFUL_URL)
    app.run(host= '0.0.0.0',port=5000)
    RemoveServiceFromDataresful(DATARESTFUL_URL)
