#!/usr/bin/env bash
etcdctl --endpoints http://10.132.0.2:2379,http://10.132.0.3:2379,http://10.132.0.4:2379 member list