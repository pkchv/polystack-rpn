import asyncio
import json
import config
from log import log
from calculator import compute
from nats.aio.client import Client

client = Client()

async def message_handler(message):
    log.info('Message received: %s' % repr(message))
    data = json.loads(message.data.decode())
    expression = data['expression']
    log.info('Evaluating expression: %s' % expression)
    result = compute(expression)
    response = json.dumps({ "result": result })
    log.info('Publishing on subject "%s": %s' % (config.pub_subject, response))
    await client.publish(config.pub_subject, response.encode())

async def main(loop):
    await client.connect(config.nats_uri, loop=loop)
    await client.subscribe(config.sub_subject,  cb=message_handler)

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    try:
        loop.run_until_complete(main(loop))
        loop.run_forever()
    except Exception as e:
        log.error(e)
    except KeyboardInterrupt:
        log.info('Program interrupted by user.')
    finally:
        log.info('Shutdown, cleaning up.')
        loop.stop()
        client.close()
