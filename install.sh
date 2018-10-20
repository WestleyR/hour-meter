#!/bin/sh
#
# Created by: Westley K
# email: westley@sylabs.io
# Date: Oct 20, 2018
# https://github.com/WestleyK/hour-meter
# Version-1.0.1
#
# MIT License
#
# Copyright (c) 2018 WestleyK
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
#

CODE_NAME="hour-meter.go"
INFO_NAME="hour_meter_info.go"
SCRIPT_NAME="hour-meter"

DATE=` date "+%B %d, %Y, %I:%M:%S %p" `
WHO=` whoami `
ON=` hostname `
WHERE=` pwd `
KERN=` uname -r `
ARCH=` uname -m `

touch $INFO_NAME
cat /dev/null > $INFO_NAME

cat << END_OF_FILE > $INFO_NAME
package main
import "fmt"
func info() {
    fmt.Print("Compiled date: $DATE\n")
    fmt.Print("Compiled by: $WHO\n")
    fmt.Print("Compiled on: $ON\n")
    fmt.Print("Compiled in: $WHERE\n")
    fmt.Print("Compiled on kernel: $KERN\n")
    fmt.Print("Compiled on architecture: $ARCH\n")
}
END_OF_FILE

if [ $(cat ~/.bashrc | grep hour-meter | wc -l) -eq 0 ]; then
    cat << EOF >> ~/.bashrc
if [ \$(ps aux | grep hour-meter | wc -l ) -le 1 ]; then
    ./hour-meter
fi
EOF
fi

wget https://raw.githubusercontent.com/WestleyK/hour-meter/master/hour-meter.go

echo "Compileing code..."
go build $CODE_NAME $INFO_NAME
rm -f $INFO_NAME
rm -f $CODE_NAME
cp $SCRIPT_NAME ~/
~/./$SCRIPT_NAME &
echo "Done."

#
# End script
#
