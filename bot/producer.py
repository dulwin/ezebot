import time
import logging
import json
from kafka import SimpleProducer, KafkaClient
from slackclient import SlackClient

PRODUCER = None

def init_kafka_producer():
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
            logging.info("RETURNING: {}".format(PRODUCER))
            return PRODUCER
        except Exception as e:
            exception = e
        time.sleep(.8)
    if exception:
        raise(exception)

def main():
    global PRODUCER
    try:
        logging.info("Initializing Producer")
        PRODUCER = init_kafka_producer()
        sc = SlackClient('xoxb-48248869860-FaqrZbiGJOqSHV1V5XNOJ48B')
        if sc.rtm_connect():
            while True:
                try:
                    messages = sc.rtm_read()
                    if messages:
                        for message in messages:
                            if message.get('type') == 'message' \
                                and 'text' in message:
                                keys = ['tagged_users_ids', 'message']
                                m = json.dumps({message.get(k) for k in keys}).encode('utf-8')

                                PRODUCER.send_messages(m, 'messages')
                    time.sleep(0.01)
                except Exception as e:
                    logging.exception(e)
        else:
            print("Connection Failed, invalid token?")

    except Exception as e:
        logging.exception(e)


if __name__ == '__main__':
    main()
