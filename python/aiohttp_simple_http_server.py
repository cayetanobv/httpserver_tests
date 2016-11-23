
import asyncio
from aiohttp import web


async def testHandler(request):
    data = {"Name":"HelloWorld","Types":["hello","world"]}
    return web.json_response(data)


async def init(loop):
    app = web.Application(loop=loop)
    app.router.add_get('/', testHandler)
    return app


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    app = loop.run_until_complete(init(loop))
    web.run_app(app, port=8090)
