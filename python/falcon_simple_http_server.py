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

# gunicorn falcon_simple_http_server:app -b localhost:8090

import falcon
import json


class TestHandler:
    def on_get(self, req, resp):
        data = {"Name":"HelloWorld","Types":["hello","world"]}
        resp.status = falcon.HTTP_200
        resp.body = json.dumps(data)


app = falcon.API()
tests = TestHandler()
app.add_route('/', tests)
