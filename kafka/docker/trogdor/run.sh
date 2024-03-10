#!/bin/bash

/opt/kafka/bin/trogdor.sh agent -c /etc/trogdor/trogdor.conf -n kafka02 &
/opt/kafka/bin/trogdor.sh coordinator -c /etc/trogdor/trogdor.conf -n kafka02 &

wait
