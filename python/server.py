import json
import pickle
from base64 import b64encode, b64decode

import requests

def authorization():
    return {
        "isBase64Encoded": False,
        "statusCode": 401,
        "headers": {},
        "body": "Please provide correct SCF-Tocken",
    }

def main_handler(event: dict,contex: dict):
    try:
        token = event['headers']['scf-token']
    except:
        return authorization()

    if toke != '12345':
        return authorization()

    data = event['body']
    kwargs = json.loads(data)
    kwargs['data'] = b64decode(kwargs['data'])

    r = requests.request(**kwargs, verify=False, allow_redirects=False)

    serialized_response = pickle.dumps(r)

    return {
        "isBase64Encoded": False,
        "statusCode": r.status_code,
        "headers": r.headers,
        "body": b64encode(serialized_response).decode('utf-8'),
    }