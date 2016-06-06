import time
import logging
import json
from kafka import KafkaConsumer, KafkaClient
from kafka.common import ConnectionError
from constants import PY_SLACK, BOT_NAME
import commands.doorcode


def init_consumer():
    end = time.time() + 20
    while time.time() < end:
        try:
            client = KafkaClient('kafka:9092')
            client.ensure_topic_exists('messages')
            consumer = KafkaConsumer('messages',
                                     bootstrap_servers=['kafka:9092'])
            return consumer
        except Exception:
            pass
        time.sleep(1)


def main(consumer):
    logging.debug("CONSUMER: {}".format(consumer))
    if consumer:
        for message in consumer:
            m = json.loads((message.value.decode()))
            if m:
                try:
                    logging.debug("Consumer received: {}".format(m))
                    # say hello
                    msg = m.get('text')
                    if 'doorcode' in msg:
                        PY_SLACK.chat_post_message(
                            m.get('channel'),
                            'The doorcode is: `{0}`!'.format(123456),
                            username=BOT_NAME
                        )
                    else:
                        PY_SLACK.chat_post_message(
                            m.get('channel'),
                            'Hi <@{}>!'.format(m.get('user')),
                            username=BOT_NAME
                        )   
                except Exception as e:
                    logging.exception(e)


if __name__ == '__main__':
    consumer = None
    while True:
        try:
            if not consumer:
                consumer = init_consumer()
            main(consumer)
        except ConnectionRefusedError:
            pass
        except ConnectionError:
            pass
        except Exception as e:
            logging.exception(e)
        time.sleep(.5)
