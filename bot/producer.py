import time
import logging
import json
from kafka import SimpleProducer, KafkaClient
from slackclient import SlackClient
from constants import BOT_ID, SLACK_TOKEN, PY_SLACK
import re

PRODUCER = None

def init_producer():
    global PRODUCER
    start = time.time()
    end = start + 60
    exception = None
    while time.time() < end:
        exception = None
        try:
            client = KafkaClient('kafka:9092')
            client.ensure_topic_exists('messages')
            PRODUCER = SimpleProducer(client)
            return PRODUCER
        except Exception as e:
            exception = e
        time.sleep(.8)
    if exception:
        raise(exception)

def main():
    global PRODUCER
    try:
        logging.debug("Initializing Producer")
        p = re.compile(r'(<@.\w+>)')
        PRODUCER = init_producer()
        sc = SlackClient(SLACK_TOKEN)
        if sc.rtm_connect():
            while True:
                try:
                    messages = sc.rtm_read()
                    for message in messages:
                        if message.get('type') == 'message' \
                            and message.get('user') != BOT_ID \
                            and BOT_ID in message.get('text', ''):
                            m = json.dumps(
                                {'user': message.get('user'),
                                 'text': p.sub('', message.get('text')),
                                 'channel': message.get('channel')}).encode('utf-8')

                            logging.debug('Producer recieved: {}'.format(m))
                            try:
                                PRODUCER.send_messages('messages', m)
                            except LeaderNotAvailableError:
                                time.sleep(1)
                                PRODUCER.send_messages('messages', m)
                            #PRODUCER.send_messages('messages', m)
                    time.sleep(0.01)
                except Exception as e:
                    logging.exception(e)
        else:
            print("Connection Failed, invalid token?")

    except Exception as e:
        logging.exception(e)


if __name__ == '__main__':
    main()
