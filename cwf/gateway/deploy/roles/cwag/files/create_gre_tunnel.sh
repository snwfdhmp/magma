#!/usr/bin/env bash
#
# Copyright (c) 2018-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.

ovs-vsctl add-br br0
sysctl net.ipv4.ip_forward=1
iptables -t mangle -A FORWARD -i br0 -p tcp --tcp-flags SYN,RST SYN -j TCPMSS --set-mss 1400
iptables -t mangle -A FORWARD -o br0 -p tcp --tcp-flags SYN,RST SYN -j TCPMSS --set-mss 1400
ovs-vsctl add-port br0 gre0 -- set interface gre0 type=gre ofport_request=32768 options:remote_ip=flow options:key=flow