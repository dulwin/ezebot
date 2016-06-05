import time
import logging
import requests
import json
from kafka import KafkaConsumer, KafkaClient
from kafka.common import ConnectionError


def init_consumer():
    end = time.time() + 20
    while time.time() < end:
        try:
            client = KafkaClient('kafka:9092')
            client.ensure_topic_exists('messages')
            consumer = KafkaConsumer('messages',
                                     group_id='persist_consumer',
                                     bootstrap_servers=['kafka:9092'])
            return consumer
        except Exception:
            pass
        time.sleep(1)


def main(consumer):

    logging.info("CONSUMER: {}".format(consumer))
    if consumer:
        for message in consumer:
            processed = json.loads((message.value.decode()))
            if processed:
                try:
                    l = "Consumer received: {}".format(processed.get('text'))
                    logging.info(l)
                    respond(processed)

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
