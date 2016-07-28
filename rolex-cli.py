#!/usr/bin/env python
# -*- coding: utf-8 -*-

# maintainer: weitao zhou <wtzhou@dataman-inc.com>

import httplib
import urllib
import json
from optparse import OptionParser
import sys

ROLEX_API_URL = '192.168.59.104'
CONFIG_FILE = '.rolex.json'


def _get_access_header():
    try:
        with open(CONFIG_FILE, 'r') as f:
            return json.loads(f.read())
    except IOError:
        print('Plz login firstly')
        sys.exit(1)

def _set_access_header(headers):
    with open(CONFIG_FILE, 'w') as f:
        f.write(json.dumps(headers))


def login(email, password):
    headers = {'Content-type': 'application/json'}
    user_info = {
        "Email": email,
        "Password": password
    }
    json_user_info = json.dumps(user_info)
    connection = httplib.HTTPConnection(ROLEX_API_URL)
    connection.request('POST','/account/v1/login', json_user_info, headers)
    response = connection.getresponse()
    try:
        resp_body = json.loads(response.read().decode())
    except ValueError:
        print("Failed to login url: " + ROLEX_API_URL)
        sys.exit(1)

    if resp_body.get('code', -1) != 0:
        print(resp_body.get('data', 'failed to login'))
        sys.exit(1)
    else:
        headers.update({"Authorization": resp_body.get('data')})

    cookie = response.getheader('Set-Cookie')
    if cookie:
        headers.update({"Cookie": cookie})
    connection.close()

    _set_access_header(headers)

def about_me():
    headers = _get_access_header()
    connection = httplib.HTTPConnection(ROLEX_API_URL)
    # test for transfer session over url query
    #connection.request('GET','/account/v1/aboutme?Authorization='+headers.get('Authorization')+'&Cookie='+urllib.quote(headers.get("Cookie", "")), '', {'Content-type': 'application/json'})
    connection.request('GET','/account/v1/aboutme', '', headers)
    response = connection.getresponse()
    print(response.read().decode())
    connection.close()

def logout():
    headers = _get_access_header()
    connection = httplib.HTTPConnection(ROLEX_API_URL)
    connection.request('POST','/account/v1/logout', '', headers)
    connection.close()

    _set_access_header({})

def list_stack():
    pass

def start_stack():
    pass

def restart_stack():
    pass

def update_stack():
    pass

def scale_stack():
    pass


if __name__ == "__main__":

    parser = OptionParser()
    parser.add_option("-e", "--email", dest="email", help="User email to login",
                      action="store", type="string")
    parser.add_option("-p", "--password", dest="password", help="User password to login",
                      action="store", type="string")
    parser.add_option("-q", "--quiet",
                      action="store_false", dest="verbose", default=True,
                      help="don't print status messages to stdout")
    (options, args) = parser.parse_args()

    login(options.email, options.password)
    about_me()
    #logout()
