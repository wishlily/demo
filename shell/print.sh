#!/bin/bash

FMT_DATE="+%Y-%m-%d %H:%M:%S %z"
# export DEBUG=true
function DEBUG() {
    if [ "$DEBUG" = "true" ]; then
        echo -e "[`date "$FMT_DATE"`][DEBUG] \\c"
        $@
    fi
}

function PRINT() {
    echo -e "[`date "$FMT_DATE"`][INFO] \\c"
    $@
}

function ERROR() {
    echo -e "[`date "$FMT_DATE"`][ERROR] \\c"
    $@
}