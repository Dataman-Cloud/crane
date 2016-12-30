#!/bin/bash

. ./config_and_precheck.sh
. ./functions.sh

. ./misc_v1_help.sh
. ./misc_v1_config.sh
# FIXME: not working for health
# . ./misc_v1_health.sh

. ./account_v1_login.sh
. ./api_v1_nodes.sh
. ./api_v1_nodes_get.sh
. ./api_v1_networks.sh
. ./api_v1_networks_get.sh

. ./api_v1_stacks_create.sh
. ./api_v1_stack_get.sh
. ./api_v1_stacks_list.sh
. ./api_v1_stack_delete.sh

# FIXME: namespace is not create correctly
# . ./registry_v1_namespace.sh
