#!/bin/bash

. ./config_and_precheck.sh
. ./functions.sh

. ./misc_v1_help.sh
. ./misc_v1_config.sh
. ./misc_v1_health.sh

. ./account_v1_login.sh
. ./api_v1_nodes.sh
. ./api_v1_nodes_get.sh
. ./api_v1_networks.sh
. ./api_v1_networks_get.sh

. ./api_v1_stacks_create.sh
. ./api_v1_stack_get.sh
. ./api_v1_stacks_list.sh
. ./api_v1_stack_delete.sh

. ./registry_v1_namespace.sh
