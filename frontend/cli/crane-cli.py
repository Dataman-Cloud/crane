#!/usr/bin/env python
# -*- coding: utf-8 -*-

# maintainer: weitao zhou <wtzhou@dataman-inc.com>

import httplib
import urllib
import json
import argparse
import sys

CRANE_API_URL = '192.168.59.105'
CONFIG_FILE = '.crane.json'

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


###############################################
#                accounts                     #
###############################################
def login(args):
    headers = {'Content-type': 'application/json'}
    user_info = {
        "Email": args.email,
        "Password": args.password
    }
    json_user_info = json.dumps(user_info)
    connection = httplib.HTTPConnection(CRANE_API_URL)
    connection.request('POST','/account/v1/login', json_user_info, headers)
    response = connection.getresponse()
    try:
        resp_body = json.loads(response.read().decode())
    except ValueError:
        print("Failed to login url: " + CRANE_API_URL)
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
    connection = httplib.HTTPConnection(CRANE_API_URL)
    # test for transfer session over url query
    #connection.request('GET','/account/v1/aboutme?Authorization='+headers.get('Authorization')+'&Cookie='+urllib.quote(headers.get("Cookie", "")), '', {'Content-type': 'application/json'})
    connection.request('GET','/account/v1/aboutme', '', headers)
    response = connection.getresponse()
    resp_body = json.loads(response.read().decode())
    print(json.dumps(resp_body, indent=4, sort_keys=True))
    connection.close()

def logout(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(CRANE_API_URL)
    connection.request('POST','/account/v1/logout', '', headers)
    connection.close()

    _set_access_header({})
    print("Logout successfully")

def get_mygroups(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(CRANE_API_URL)
    connection.request('GET','/account/v1/accounts/' + args.uid + '/groups', '', headers)
    response = connection.getresponse()
    resp_body = json.loads(response.read().decode())
    print(json.dumps(resp_body, indent=4, sort_keys=True))
    connection.close()


###############################################
#                Stacks                       #
###############################################
def list_stack(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(CRANE_API_URL)
    connection.request('GET','/api/v1/stacks', '', headers)
    response = connection.getresponse()
    try:
        resp_body = json.loads(response.read().decode())
        print(json.dumps(resp_body, indent=4, sort_keys=True))
    except IOError:
        print('Failed to list stacks')
        sys.exit(1)
    connection.close()

def create_stack(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(CRANE_API_URL)
    try:
        request_body = {
            "Stack": json.loads(args.file.read()),
            "Namespace": args.stack_name
        }
        connection.request('POST','/api/v1/stacks?group_id='+args.group, json.dumps(request_body), headers)
    except ValueError:
        print('JSON type error')
        sys.exit(1)

    response = connection.getresponse()
    try:
        resp_body = json.loads(response.read().decode())
        print(json.dumps(resp_body, indent=4, sort_keys=True))
    except IOError:
        print('Failed to create stack')
        sys.exit(1)
    connection.close()

def list_stack_services(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(CRANE_API_URL)
    connection.request('GET','/api/v1/stacks/' + args.stack_name + '/services', '', headers)
    response = connection.getresponse()
    try:
        resp_body = json.loads(response.read().decode())
        print(json.dumps(resp_body, indent=4, sort_keys=True))
    except IOError:
        print('Failed to list stacks services')
        sys.exit(1)
    connection.close()

def scale_service(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(CRANE_API_URL)
    request_body = {
        "Scale": args.amount
    }
    connection.request('PATCH','/api/v1/stacks/' + args.stack_name + '/services/' + args.service_id, json.dumps(request_body), headers)
    response = connection.getresponse()
    try:
        resp_body = json.loads(response.read().decode())
        print(json.dumps(resp_body, indent=4, sort_keys=True))
    except IOError:
        print('Failed to scale service task')
        sys.exit(1)
    connection.close()

def restart_stack():
    pass

def update_stack():
    pass


###############################################
#                networks                     #
###############################################
def list_networks(args):
    headers = _get_access_header()
    connection = httplib.HTTPConnection(CRANE_API_URL)
    connection.request('GET','/api/v1/networks', '', headers)
    response = connection.getresponse()
    resp_body = json.loads(response.read().decode())
    print(json.dumps(resp_body, indent=4, sort_keys=True))
    connection.close()


if __name__ == "__main__":

    class MyParser(argparse.ArgumentParser):
        def error(self, message):
            sys.stderr.write('error: %s\n' % message)
            self.print_help()
            sys.exit(2)

    parser = MyParser(prog='Crane-Cli')
    subparsers = parser.add_subparsers()

    parser_login = subparsers.add_parser('login', help="User login")
    parser_login.add_argument("-e", "--email", help="User email to login", type=str, required=True)
    parser_login.add_argument("-p", "--password", help="User password to login", type=str, required=True)
    parser_login.set_defaults(func=login)

    parser_logout = subparsers.add_parser('logout', help="User logout")
    parser_logout.set_defaults(func=logout)

    parser_aboutme = subparsers.add_parser('aboutme', help="About Me")
    parser_aboutme.set_defaults(func=about_me)

    parser_get_mygroups = subparsers.add_parser('mygroups', help="Return my groups info")
    parser_get_mygroups.add_argument("-u", "--uid", help="My user id (find it by cmd aboutme)", type=str, required=True)
    parser_get_mygroups.set_defaults(func=get_mygroups)

    parser_list_stacks = subparsers.add_parser('list_stack', help="List my stacks")
    parser_list_stacks.set_defaults(func=list_stack)

    parser_create_stack = subparsers.add_parser('create_stack', help="Create stack")
    parser_create_stack.add_argument("-G", "--group", help="The group_id stack will belong to", type=str, required=True)
    parser_create_stack.add_argument("-n", "--stack_name", help="Stack Name", type=str, required=True)
    parser_create_stack.add_argument("-f", "--file", help="bundle json containing stack info", type=argparse.FileType('r'), required=True)
    parser_create_stack.set_defaults(func=create_stack)

    parser_list_stack_s = subparsers.add_parser('list_stack_services', help="List stack services")
    parser_list_stack_s.add_argument("-n", "--stack_name", help="Stack Name", type=str, required=True)
    parser_list_stack_s.set_defaults(func=list_stack_services)

    parser_scale_service = subparsers.add_parser('scale_service', help="Scale service tasks")
    parser_scale_service.add_argument("-n", "--stack_name", help="Stack Name", type=str, required=True)
    parser_scale_service.add_argument("-s", "--service_id", help="Service ID (find it by cmd list_stack_services)", type=str, required=True)
    parser_scale_service.add_argument("-a", "--amount", help="Desired task amount", type=int, required=True)
    parser_scale_service.set_defaults(func=scale_service)

    parser_list_networks = subparsers.add_parser('list_networks', help="List Networks")
    parser_list_networks.set_defaults(func=list_networks)

    args = parser.parse_args()
    args.func(args)
