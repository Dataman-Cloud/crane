#!/usr/bin/env python
# -*- coding: utf-8 -*-

# maintainer: weitao zhou <wtzhou@dataman-inc.com>

import httplib
import urllib
import json
import argparse
import sys

ROLEX_API_URL = '192.168.59.104'
CONFIG_FILE = '.rolex.json'

__author__ = "weitao zhou"

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


def login(args):
    headers = {'Content-type': 'application/json'}
    user_info = {
        "Email": args.email,
        "Password": args.password
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

def about_me(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(ROLEX_API_URL)
    # test for transfer session over url query
    #connection.request('GET','/account/v1/aboutme?Authorization='+headers.get('Authorization')+'&Cookie='+urllib.quote(headers.get("Cookie", "")), '', {'Content-type': 'application/json'})
    connection.request('GET','/account/v1/aboutme', '', headers)
    response = connection.getresponse()
    print(response.read().decode())
    connection.close()

def logout(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(ROLEX_API_URL)
    connection.request('POST','/account/v1/logout', '', headers)
    connection.close()

    _set_access_header({})
    print("Logout successfully")

def list_stack():
    pass

def create_stack():
    pass

def restart_stack():
    pass

def update_stack():
    pass

def scale_stack():
    pass


if __name__ == "__main__":

    class MyParser(argparse.ArgumentParser):
        def error(self, message):
            sys.stderr.write('error: %s\n' % message)
            self.print_help()
            sys.exit(2)

    parser = MyParser(prog='Rolex-Cli')
    subparsers = parser.add_subparsers()

    parser_login = subparsers.add_parser('login', help="User login")
    parser_login.add_argument("-e", "--email", help="User email to login", type=str, required=True)
    parser_login.add_argument("-p", "--password", help="User password to login", type=str, required=True)
    parser_login.set_defaults(func=login)

    parser_logout = subparsers.add_parser('logout', help="User logout")
    parser_logout.set_defaults(func=logout)

    parser_aboutme = subparsers.add_parser('aboutme', help="About Me")
    parser_aboutme.set_defaults(func=about_me)

    args = parser.parse_args()
    args.func(args)
