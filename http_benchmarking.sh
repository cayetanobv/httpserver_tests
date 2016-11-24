
####################################
# TESTING HTTP SERVERS - BENCHMARK #
####################################


_BCONNS=400
_BTHREADS=1
_BTIME=5
_BHOST="http://127.0.0.1:8090/"
_BJSON={"Name":"HelloWorld","Types":["hello","world"]}
#
#  Author: Cayetano Benavent, 2016.
#  cayetano.benavent@geographica.gs
#
#  This program is free software; you can redistribute it and/or modify
#  it under the terms of the GNU General Public License as published by
#  the Free Software Foundation; either version 2 of the License, or
#  (at your option) any later version.
#
#  This program is distributed in the hope that it will be useful,
#  but WITHOUT ANY WARRANTY; without even the implied warranty of
#  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#  GNU General Public License for more details.
#
#  You should have received a copy of the GNU General Public License
#  along with this program; if not, write to the Free Software
#  Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
#  MA 02110-1301, USA.
#

_PYTHON_SANIC="python3 python/sanic_simplehttpserver.py"
_GOLANG_STD="./golang/simple_http_server_json"
_GOLANG_VALYALA="./golang/valyala_simple_httpserver"
_NODEJS="node nodejs/simple_httpserver_json.js"
_PYTHON_AIOHTTP="python3 python/aiohttp_simple_http_server.py"
_PYTHON_AIOHTTP_UVLOOP="python3 python/aiohttp_uvloop_simple_http_server.py"
_PYTHON_FLASK="python3 python/flask_simple_http_server.py"

declare -a cmds_arr=(
    "${_PYTHON_SANIC}"
    "${_GOLANG_STD}"
    "${_GOLANG_VALYALA}"
    "${_NODEJS}"
    "${_PYTHON_AIOHTTP}"
    "${_PYTHON_AIOHTTP_UVLOOP}"
    "${_PYTHON_FLASK}"
  )

printf "\n[%s]\n" "$(date)"
printf "\nStarting benchmark...\n"
printf "\nSimple HTTP Servers with a JSON response:\n %s\n" "${_BJSON}"
printf "\nVersions:\n"
printf "\tOperative System - \t [%s] \n" "$(uname -o)"
printf "\tKernel - \t [%s] \n" "$(uname -vr)"
printf "\tHardware-platform - \t [%s] \n" "$(uname -i)"
printf "\tPython - \t [%s] \n" "$(python3 --version)"
printf "\tGolang - \t [%s] \n" "$(go version)"
printf "\tNodeJS - \t [%s] \n" "$(node --version)"


for cmd in "${cmds_arr[@]}"
  do
    printf "\nStart process: %s\n" "$cmd";
    $cmd & sleep 1;
    wrk -t"${_BTHREADS}" -c"${_BCONNS}" -d"${_BTIME}"s "${_BHOST}" ;
    pkill -9 -x -f "$cmd";
    sleep 2;
    printf "\nFinish process: %s\n" "$cmd";
    echo "------";
  done
